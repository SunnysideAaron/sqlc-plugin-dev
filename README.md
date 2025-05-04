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
      make git2      
  These will only work once. "make git" does the core repositories of sqlc, plugin-sdk-go, and sqlc-gen-go. This is probably the minimum to get started. Otherwise your IDE / AI wont be able to see all the connections and how it all works together. "make git2" grabs all other known repos of sqlc plugins. Grab all the examples you can I say! You could of course clone repos one by one depending on what you want to play with.
  
  Any repos cloned into code/ignore wont become part of this projects code.
- Launch container for development

      make bash

## databases

There are two docker containers of databases for use. MySQL and PostgreSQL.

    make my-up
    make post-up
    make my-down
    make post-down

Any sql files in the /initdb gets run when the container starts if /persistence is empty.

They save data in /persistence. Delete this folder to reset the database.

These are similar but not the same as the ones in sqlc/docker-compose.yml. I didn't want to mess with docker running inside docker.

## Notes

Plugins are compiled to WASM. This means they don't have to be developed in Go. [sqlc-gen-greeter](https://github.com/sqlc-dev/sqlc-gen-greeter) is an example of a plugin written in Rust.

## Core Repositories

- [sqlc](https://github.com/sqlc-dev/sqlc)
  - main executable
- [plugin-sdk-go](https://github.com/sqlc-dev/plugin-sdk-go)
  - connection between sqlc and plugins.

## Plugins Developed by sqlc Team

- Check [sqlc-dev/repositories](https://github.com/orgs/sqlc-dev/repositories) and [Database and language support](https://docs.sqlc.dev/en/latest/reference/language-support.html#database-and-language-support) for updated lists.
- [sqlc-gen-go](https://github.com/sqlc-dev/sqlc-gen-go)
  - example plugin, code is *extracted* from sqlc. It is there to fork other plugins from.
  - It appears code is extracted from:
    - sqlc\internal\codegen\golang -> sqlc-gen-go\internal
    - sqlc\internal\inflection -> sqlc-gen-go\internal\inflection
  - generates files:
    - db.go
    - models.go
    - query.sql.go
    - querier.go (optional see flag: emit_interface)
- [sqlc-gen-python](https://github.com/sqlc-dev/sqlc-gen-python)
- [sqlc-gen-kotlin](https://github.com/sqlc-dev/sqlc-gen-kotlin)
- [sqlc-gen-typescript](https://github.com/sqlc-dev/sqlc-gen-typescript)
- [sqlc-gen-greeter](https://github.com/sqlc-dev/sqlc-gen-greeter)
  - rust example. last updated 2022/06.

## Known 3rd party plugins

- Check [Community language support](https://docs.sqlc.dev/en/latest/reference/language-support.html#community-language-support) for updated lists.
- [sqlc-gen-from-template](https://github.com/fdietze/sqlc-gen-from-template)
  - This is interesting since it's language agnostic. **TODO look at this.**
- [sqlc-gen-csharp](https://github.com/DaredevilOSS/sqlc-gen-csharp)
  - [Introducing SQLC C# Plugin: A reverse ORM for .NET Developers](https://www.reddit.com/r/dotnet/comments/1hp6sa5/introducing_sqlc_c_plugin_a_reverse_orm_for_net/)
- [sqlc-gen-fsharp](https://github.com/kaashyapan/sqlc-gen-fsharp)
- [sqlc-gen-java](https://github.com/tandemdude/sqlc-gen-java)
- [sqlc-plugin-php-dbal](https://github.com/lcarilla/sqlc-plugin-php-dbal)
- [sqlc-gen-ruby](https://github.com/DaredevilOSS/sqlc-gen-ruby)
- [sqlc-gen-zig](https://github.com/tinyzimmer/sqlc-gen-zig)

## Other plugins

- [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server)
  - forked from sqlc-gen-go
  - generates sqlc-gen-go files + 
    - openapi.yml
    - routes.go
    - service.go
    - main.go

### Examples

- code/sqlc
  - examples 0-3 are all from the sqlc docs. They are included here to make copying for your own testing easier. They are already adjusted to work with the databases provided.
  - /example-0: [Developing sqlc](https://docs.sqlc.dev/en/latest/guides/development.html)
  - /example-1: [Getting started with MySQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html)
  - /example-2: [Getting started with PostgreSQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)
  - /example-3: [Getting started with SQLite](https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html)
- /code/ignore/sqlc/examples
  - These are included with the sqlc code. There isn't any documentation so I'm not clear on what they are examples of. Yet.
  - TODO find out more.
  - /authors
    - basic starter example. Used in docs.
  - /batch
    - I'm guessing something with batch inserts or transactions
  - /booktest
  - /jets
    - has some alter tables in schema
  - /ondeck
    - has schema spread across files

### Examples TODO

example with all settings to make testing easier

each plugin example


## TODO

update after release
https://github.com/sqlc-dev/sqlc/discussions/3077

https://github.com/sqlc-dev/sqlc/issues/3945


and update my own post about 3rd party plugins
