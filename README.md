# sqlc-plugin-dev

This project contains notes and examples on how to develop sqlc plugins.

I couldn't find any documentation on developing plugins. So I created the help I needed. Hopefully this can be useful to others.

- [sqlc.dev](https://sqlc.dev/)
- [sqlc documentation](https://docs.sqlc.dev/en/latest/index.html)
- [github: sqlc](https://github.com/sqlc-dev/sqlc)
- [github: other sqlc repos](https://github.com/orgs/sqlc-dev/repositories?type=all)

Where to get help:

- [github: sqlc discussions](https://github.com/sqlc-dev/sqlc/discussions)
- [sqlc Discord server](https://discord.com/invite/EcXzGe5SEs)

## Getting started



requirements
 docker
 make
git

install requirements

copy repos locally into same workspace so IDE / AI can see all the connections

build container.

    docker compose build

Launch container for development

    docker compose run --remove-orphans --service-ports code bash

## Repos

- sqlc
  - main executable
- plugin-sdk-go
  - connection between sqlc and plugins
- sqlc-gen-go
  - example plugin, extracted from sqlc.
  - A big part of the code comes from sqlc\internal\codegen\golang
- sqlc-gen-go-server
  - example plugin forked from sqlc-gen-go
  - generates sqlc-gen-go files + 
    - openapi.yml
    - routes.go
    - service.go
    - main.go

## Known 3rd party plugins

https://www.reddit.com/r/dotnet/comments/1hp6sa5/introducing_sqlc_c_plugin_a_reverse_orm_for_net/

https://github.com/DaredevilOSS/sqlc-gen-csharp




## Examples

example sqlc dev tutorial
mysql
postres
sqlite

each plugin example

## TODO

update after release
https://github.com/sqlc-dev/sqlc/discussions/3077


and update my own post about 3rd party plugins
