// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: queries.sql

package author

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio, created_at
) VALUES (
  $1, $2, $3
)
RETURNING id, name, bio, created_at
`

type CreateAuthorParams struct {
	Name      string
	Bio       sql.NullString
	CreatedAt sql.NullTime
}

// http: POST /authors
func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Name, arg.Bio, arg.CreatedAt)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1
`

// http: DELETE /authors/{id}
func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio, created_at FROM authors
WHERE id = $1 LIMIT 1
`

// http: GET /authors/{id}
func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio, created_at FROM authors
ORDER BY name
`

// http: GET /authors
func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthorBio = `-- name: UpdateAuthorBio :exec
UPDATE authors
SET bio = $1
WHERE id = $2
`

type UpdateAuthorBioParams struct {
	Bio sql.NullString
	ID  int64
}

// http: PATCH /authors/{id}/bio
func (q *Queries) UpdateAuthorBio(ctx context.Context, arg UpdateAuthorBioParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthorBio, arg.Bio, arg.ID)
	return err
}
