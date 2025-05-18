# example3-0: [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server)

Just the example from the ReadMe of [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server).

Some commands that might be useful.

    make code-bash
    cd /home/code/3-sqlc-gen-go-server/example3-0/
    sqlc generate

Note there is something wrong with the paths. Try running sqlc-gen-go-server with sqlc 1.25.0 not 1.29.0.
https://github.com/walterwanderley/sqlc-gen-go-server/pull/1/files

Note sqlc.yaml is pointing at the local build of of wasm file.

if you want to try building sqlc-gen-go-server

    cd /home/code/ignore/sqlc-gen-go-server
    make all

