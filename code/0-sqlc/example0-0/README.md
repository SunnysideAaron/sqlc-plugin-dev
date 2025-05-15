# example0-0: build [Developing sqlc](https://docs.sqlc.dev/en/latest/guides/development.html)

Make certain that you have cloned the [sqlc](https://github.com/sqlc-dev/sqlc) repository into /code/ignore/

Review the instructions from [Developing sqlc](https://docs.sqlc.dev/en/latest/guides/development.html). The following commands might provide some guidance. They are slightly adjusted.

      make bash
      cd /home/code/ignore/0-sqlc
      go build -o /go/bin/sqlc-dev ./cmd/sqlc
      go build -o /go/bin/sqlc-gen-json ./cmd/sqlc-gen-json
      sqlc-dev version

Testing

      go test ./...
      go test --tags=examples ./...

Note that tests are failing. See [issue: sqlc tests are failing](https://github.com/SunnysideAaron/sqlc-plugin-dev/issues/1)
