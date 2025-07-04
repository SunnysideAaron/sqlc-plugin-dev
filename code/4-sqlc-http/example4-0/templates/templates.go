package templates

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"mymodule/internal/server/htmx"
	watchersse "mymodule/internal/server/watcher"
)

// Content-Security-Policy
const csp = "default-src 'self'; img-src 'self' data: ; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline';"

type key int

const (
	templatesContext key = iota
	messageContext
	paginationContext
)

type breadCrumb struct {
	Name string
	Href string
}

var (
	ErrTemplateNotFound = errors.New("template not found")

	//go:embed *
	templatesFS embed.FS
	funcs       = template.FuncMap{}

	provider templatesProvider
)

func init() {
	context := os.Getenv("WEB_CONTEXT")
	if context == "" {
		context = "/"
	}
	funcs["WebContext"] = func() string {
		return context
	}
}

const defaultLimit = 10

type Pagination struct {
	request *http.Request
	Limit   int64
	Offset  int64
}

func (p *Pagination) From() int64 {
	return p.Offset + 1
}

func (p *Pagination) To() int64 {
	return p.Offset + p.validLimit()
}

func (p *Pagination) Next() int64 {
	return p.Offset + p.validLimit()
}

func (p *Pagination) Prev() int64 {
	limit := p.validLimit()
	offset := p.Offset - limit
	if offset < 0 {
		offset = 0
	}
	return offset
}

func (p *Pagination) URL(limit, offset int64) string {
	if p == nil {
		return ""
	}
	if offset < 0 {
		offset = 0
	}
	if limit == 0 {
		limit = defaultLimit
	}
	var url strings.Builder
	url.WriteString(p.request.URL.Path)
	url.WriteString("?")
	for k := range p.request.URL.Query() {
		if k == "limit" || k == "offset" {
			continue
		}
		url.WriteString(fmt.Sprintf("%s=%s&", k, p.request.URL.Query().Get(k)))
	}
	if limit > 0 {
		url.WriteString(fmt.Sprintf("limit=%d&offset=%d", limit, offset))
	} else {
		url.WriteString(fmt.Sprintf("offset=%d", offset))
	}
	return url.String()
}

func (p *Pagination) validLimit() int64 {
	limit := p.Limit
	if limit == 0 {
		limit = 10
	}
	return limit
}

type templatesProvider interface {
	FullTemplate(string) (*template.Template, error)
	DynamicTemplate(string) (*template.Template, error)
	TemplatesFS() fs.FS
	DevMode() bool
}

func RegisterHandlers(mux *http.ServeMux, devMode bool) error {
	base := template.New("base.html").Funcs(funcs)
	content := template.New("content.html").Funcs(funcs)
	componentsTemplates := []string{
		"components/breadcrumbs.html",
		"components/hx-context.html",
		"components/message.html",
	}
	layoutTemplates := []string{
		"layout/base.html",
		"layout/header.html",
		"layout/footer.html",
	}
	if devMode {
		templatesPath := "templates"
		provider = &templateDevRender{
			templatesFS:   os.DirFS(templatesPath),
			base:          base,
			content:       content,
			baseTemplates: append(componentsTemplates, layoutTemplates...),
			components:    componentsTemplates,
		}
		watcher, err := watchersse.New(templatesPath, "web")
		if err != nil {
			return err
		}
		watcher.Start(context.Background())
		mux.Handle("GET /reload", watcher)
	} else {
		tr := templateRender{
			templatesFS: templatesFS,
			content: template.Must(content.ParseFS(templatesFS,
				append(componentsTemplates, "layout/content.html")...)),
			base: template.Must(base.ParseFS(templatesFS,
				append(componentsTemplates, layoutTemplates...)...)),
		}
		if err := tr.Compile(); err != nil {
			return err
		}
		provider = &tr
	}
	mux.HandleFunc("GET /", templatesHandler)
	return nil
}

func templatesHandler(w http.ResponseWriter, r *http.Request) {
	if err := RenderHTML[any](w, r, nil); err != nil {
		slog.Error("render html", "error", err, "path", r.URL.Path)
	}
}

type templateOpts struct {
	Title   string
	Content any
	DevMode bool
}

func defaultTemplateOpts(content any) templateOpts {
	return templateOpts{
		Title:   "Mymodule",
		Content: content,
		DevMode: provider.DevMode(),
	}
}

func ContextWithPagination(ctx context.Context, pagination *Pagination) context.Context {
	return context.WithValue(ctx, paginationContext, pagination)
}

func ContextWithMessage(ctx context.Context, msg htmx.Message) context.Context {
	return context.WithValue(ctx, messageContext, msg)
}

func ContextWithTemplates(ctx context.Context, templates ...string) context.Context {
	return context.WithValue(ctx, templatesContext, templates)
}

type Content[T any] struct {
	Data    T
	Request *http.Request
}

func (c Content[T]) HxRequest() bool {
	return htmx.HXRequest(c.Request)
}

func (c Content[T]) Pagination() *Pagination {
	pagination, _ := c.Request.Context().Value(paginationContext).(*Pagination)
	if pagination != nil {
		pagination.request = c.Request
	}
	return pagination
}

func (c Content[T]) BaseHref() string {
	schema := "http"
	if forwardedProto := c.Request.Header.Get("X-Forwarded-Proto"); forwardedProto != "" {
		schema = forwardedProto
	} else if c.Request.TLS != nil {
		schema = "https"
	}
	context := os.Getenv("WEB_CONTEXT")
	return fmt.Sprintf("%s://%s%s/", schema, c.Request.Host, strings.TrimSuffix(context, "/"))
}

func (c Content[T]) BreadCrumbsFromRequest() []breadCrumb {
	switch c.Request.Pattern {
	case "POST /authors":
		return breadCrumbsFromStrings("Author", "/", "Create Author")
	case "DELETE /authors/{id}":
		return breadCrumbsFromStrings("Author", "/", "List Authors", "/authors", "Delete Author")
	case "GET /authors/{id}":
		return breadCrumbsFromStrings("Author", "/", "List Authors", "/authors", "Get Author")
	case "GET /authors":
		return breadCrumbsFromStrings("Author", "/", "List Authors")
	case "PATCH /authors/{id}/bio":
		return breadCrumbsFromStrings("Author", "/", "List Authors", "/authors", "Get Author", "/authors/"+c.Request.PathValue("id"), "Update Author Bio")
	default:
		switch c.Request.URL.Path {
		case "/app/author/create_author":
			return breadCrumbsFromStrings("Author", "/", "Create Author")
		case "/app/author/delete_author":
			return breadCrumbsFromStrings("Author", "/", "List Authors", "/authors", "Delete Author")
		case "/app/author/get_author":
			return breadCrumbsFromStrings("Author", "/", "List Authors", "/authors", "Get Author")
		case "/app/author/list_authors":
			return breadCrumbsFromStrings("Author", "/", "List Authors")
		case "/app/author/update_author_bio":
			return breadCrumbsFromStrings("Author", "/", "List Authors", "/authors", "Update Author Bio")
		}
	}

	return nil
}

func (c Content[T]) HasQuery(key string) bool {
	return c.Request.URL.Query().Has(key)
}

func (c Content[T]) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c Content[T]) MessageContext() *htmx.Message {
	if msg, ok := c.Request.Context().Value(messageContext).(htmx.Message); ok {
		return &msg
	}
	return nil
}

func RenderHTML[T any](w http.ResponseWriter, r *http.Request, content T) (err error) {
	templates := contextTemplates(r)
	if len(templates) == 0 {
		if msg, ok := r.Context().Value(messageContext).(htmx.Message); ok {
			return msg.Render(w, r)
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil
	}
	var (
		tmpl *template.Template
		obj  any
	)
	if htmx.HXRequest(r) {
		tmpl, err = provider.DynamicTemplate(templates[0])
		obj = Content[T]{
			Data:    content,
			Request: r,
		}
	} else {
		tmpl, err = provider.FullTemplate(templates[0])
		obj = defaultTemplateOpts(Content[T]{
			Data:    content,
			Request: r,
		})
	}
	if err != nil {
		switch {
		case errors.Is(err, ErrTemplateNotFound):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return nil
		default:
			slog.ErrorContext(r.Context(), "render html", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if len(templates) > 1 {
		tmpl, err = tmpl.Clone()
		if err != nil {
			return err
		}
		tmpl, err = tmpl.ParseFS(provider.TemplatesFS(), templates[1:]...)
		if err != nil {
			return err
		}
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Content-Security-Policy", csp)
	w.WriteHeader(http.StatusOK)
	err = tmpl.Execute(w, obj)
	return err
}

type templateDevRender struct {
	templatesFS   fs.FS
	base          *template.Template
	content       *template.Template
	baseTemplates []string
	components    []string
}

func (t *templateDevRender) DevMode() bool {
	return true
}

func (t *templateDevRender) FullTemplate(path string) (*template.Template, error) {
	f, err := t.templatesFS.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrTemplateNotFound, path)
	}
	f.Close()
	return template.Must(t.base.Clone()).ParseFS(t.templatesFS, append(t.baseTemplates, path)...)
}

func (t *templateDevRender) DynamicTemplate(path string) (*template.Template, error) {
	f, err := t.templatesFS.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrTemplateNotFound, path)
	}
	f.Close()
	return template.Must(t.content.Clone()).ParseFS(t.templatesFS, append(t.components, "layout/content.html", path)...)
}

func (t *templateDevRender) TemplatesFS() fs.FS {
	return t.templatesFS
}

type templateRender struct {
	templatesFS      fs.FS
	fullTemplates    map[string]*template.Template
	dynamicTemplates map[string]*template.Template
	base             *template.Template
	content          *template.Template
}

func (t *templateRender) DevMode() bool {
	return false
}

func (t *templateRender) Compile() error {
	t.dynamicTemplates = make(map[string]*template.Template)
	t.fullTemplates = make(map[string]*template.Template)
	err := fs.WalkDir(t.templatesFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || strings.HasPrefix(path, "components/") || strings.HasPrefix(path, "layout/") {
			return nil
		}

		if strings.HasSuffix(path, ".html") {
			dynamicClone, err := t.content.Clone()
			if err != nil {
				return err
			}
			tmpl, err := dynamicClone.ParseFS(t.templatesFS, path)
			if err != nil {
				return err
			}
			t.dynamicTemplates[path] = tmpl
			fullClone, err := t.base.Clone()
			if err != nil {
				return err
			}
			tmplFull, err := fullClone.ParseFS(t.templatesFS, path)
			if err != nil {
				return err
			}
			t.fullTemplates[path] = tmplFull
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("compiling templates: %w", err)
	}
	return nil
}

func (t *templateRender) FullTemplate(path string) (*template.Template, error) {
	tmpl, ok := t.fullTemplates[path]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrTemplateNotFound, path)
	}
	return tmpl, nil
}

func (t *templateRender) DynamicTemplate(path string) (*template.Template, error) {
	tmpl, ok := t.dynamicTemplates[path]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrTemplateNotFound, path)
	}
	return tmpl, nil
}

func (t *templateRender) TemplatesFS() fs.FS {
	return t.templatesFS
}

func contextTemplates(r *http.Request) []string {
	if v, ok := r.Context().Value(templatesContext).([]string); ok {
		return v
	}

	path := strings.TrimPrefix(strings.TrimPrefix(r.Pattern, r.Method+" "), "/")
	if path == "" {
		path = strings.TrimPrefix(r.URL.Path, "/")
	}
	if path == "" {
		path = "index"
	}
	path = path + ".html"
	return []string{path}
}

func breadCrumbsFromStrings(items ...string) []breadCrumb {
	breadcrumbs := make([]breadCrumb, 0)
	for i := 0; i < len(items); i = i + 2 {
		var bc breadCrumb
		bc.Name = items[i]
		j := i + 1
		if j < len(items) {
			bc.Href = items[j]
		}
		breadcrumbs = append(breadcrumbs, bc)
	}
	return breadcrumbs
}
