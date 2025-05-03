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

## Requirements

- docker
- make
- git

## Getting started

- Install requirements
- Build container

      make build
- Clone core repos into /code/ignore. 

      make git
  Will only work once but will git sqlc, plugin-sdk-go, and sqlc-gen-go repositories. This is probably the minimum to get started. Otherwise your IDE / AI wont be able to see all the connections and how it all works together.
  
  Any repos cloned into code/ignore wont become part of this projects code.
- Launch container for development

      make bash

## Notes

Plugins are compiled to WASM. This means they don't have to be developed in Go. [sqlc-gen-greeter](https://github.com/sqlc-dev/sqlc-gen-greeter) is an example of a plugin written in Rust.

## Core Repositories

- [sqlc](https://github.com/sqlc-dev/sqlc)
  - main executable
- [plugin-sdk-go]()
  - connection between sqlc and plugins.
  - TODO is this now part of sqlc? is the code extracted from sqlc? hasn't been updated in a bit. sqlc-gen-go is still pointing at it.

## Plugins Developed by sqlc Team

- [sqlc-dev/repositories](https://github.com/orgs/sqlc-dev/repositories)
  - check here for updated list.
- [sqlc-gen-go](https://github.com/sqlc-dev/sqlc-gen-go)
  - example plugin, code is *extracted* from sqlc. This is already part of sqlc and is only there to fork other plugins from.
  - A big part of the code comes from sqlc\internal\codegen\golang
  - TODO where does the rest come from?

TODO list files output
optional interface


- [sqlc-gen-python](https://github.com/sqlc-dev/sqlc-gen-python)
- [sqlc-gen-kotlin](https://github.com/sqlc-dev/sqlc-gen-kotlin)
- [sqlc-gen-typescript](https://github.com/sqlc-dev/sqlc-gen-typescript)
- [sqlc-gen-greeter](https://github.com/sqlc-dev/sqlc-gen-greeter)
  - rust example. last updated 2022/06.

## Known 3rd party plugins

- [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server)
  - example plugin forked from sqlc-gen-go
  - generates sqlc-gen-go files + 
    - openapi.yml
    - routes.go
    - service.go
    - main.go
- [sqlc-gen-csharp](https://github.com/DaredevilOSS/sqlc-gen-csharp)
  - [Introducing SQLC C# Plugin: A reverse ORM for .NET Developers](https://www.reddit.com/r/dotnet/comments/1hp6sa5/introducing_sqlc_c_plugin_a_reverse_orm_for_net/)



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
