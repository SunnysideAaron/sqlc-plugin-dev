# example0-5: Sandbox

Developing on sqlc against PostgreSQL.

Note I found building sqlc very slow. You may want to try sqlc-gen-go for initial development first.

Some commands that might be useful.

	make code-bash
	cd /home/code/ignore/sqlc
	go build -o /go/bin/sqlc-dev ./cmd/sqlc

    cd /home/code/0-sqlc/example0-5/
    sqlc-dev generate
    go get github.com/jackc/pgx/v5
    go build ./...
    go run tutorial.go

