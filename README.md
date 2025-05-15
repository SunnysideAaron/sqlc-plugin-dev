# sqlc-plugin-dev

This project contains notes and examples on how to develop sqlc plugins.

I couldn't find any documentation on developing plugins. So I created the help I needed. Hopefully this can be useful to others.
TODO
repos listing
examples

TODO (is this correct?)
if you want to change what is in the files output by sqlc. Then fork sqlc-gen-go. or json? what other outputs baked in?
if you want different files all together then fork sqlc-gen-from-template.
if you want a new plugin written in a different language then???? what? How to do that?

sqlc links:

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
  ```console
  make code-build
  ```    
- Clone core repos into /code/ignore. 
  ```console
  make code-git
  make code-git2
  ```

  These will only work once. "make git" does the core repositories of sqlc, plugin-sdk-go, and sqlc-gen-go. This is probably the minimum to get started. Otherwise your IDE / AI wont be able to see all the connections and how it all works together. "make git2" grabs all other known repos of sqlc plugins. Grab all the examples you can I say! You could of course clone repos one by one depending on what you want to play with.
  
  Any repos cloned into code/ignore wont become part of this projects code.
- Launch container for development
  ```console
  make bash
  ```

## Databases

There are two docker containers of databases for use. MySQL and PostgreSQL.
```console
make my-up
make post-up
make my-down
make post-down
```

Any sql files in the /initdb gets run when the container starts if /persistence is empty.

They save data in /persistence. Delete this folder to reset the database.

These are similar but not the same as the ones in sqlc/docker-compose.yml. I didn't want to mess with docker running inside docker.

## Notes


## Core Repositories

- [sqlc](https://github.com/sqlc-dev/sqlc)
  - main executable
- [plugin-sdk-go](https://github.com/sqlc-dev/plugin-sdk-go)
  - connection between sqlc and plugins.

## Plugins Developed in Go

- Check [sqlc-dev/repositories](https://github.com/orgs/sqlc-dev/repositories) and [Database and language support](https://docs.sqlc.dev/en/latest/reference/language-support.html#database-and-language-support) for updated lists.

- [sqlc-gen-from-template](https://github.com/fdietze/sqlc-gen-from-template)
  - Writen in language: Go
  - Produces code for language: Any (Template)
  - Writen by: fdietze
  - This is interesting since it's language agnostic. **TODO look at this.**
- [sqlc-gen-fsharp](https://github.com/kaashyapan/sqlc-gen-fsharp)
  - Writen in language: Go
  - Produces code for language: F#
  - Written by: kaashyapan
- [sqlc-gen-go](https://github.com/sqlc-dev/sqlc-gen-go) **todo flag this**
  - Writen in language: Go
  - Produces code for language: Go
  - Writen by: sqlc team
  - example plugin, code is *extracted* from sqlc. It is there to fork other plugins from.
  - It appears code is extracted from:
    - sqlc\internal\codegen\golang -> sqlc-gen-go\internal
    - sqlc\internal\inflection -> sqlc-gen-go\internal\inflection
  - generates files:
    - db.go
    - models.go
    - query.sql.go
    - querier.go (optional see flag: emit_interface)  
- [sqlc-gen-go-server](https://github.com/walterwanderley/sqlc-gen-go-server)
  - Writen in language: Go
  - Produces code for language: Go    
  - Written by: walterwanderley
  - forked from sqlc-gen-go
  - generates sqlc-gen-go files + 
    - openapi.yml
    - routes.go
    - service.go
    - main.go
  - **TODO** list walters other plugins
  - pulls templates from 
    - [sqlc-http](github.com/walterwanderley/sqlc-http/)
    - [sqlc-connect](github.com/walterwanderley/sqlc-connect)
    - [sqlc-grpc](github.com/walterwanderley/sqlc-grpc)
- [sqlc-gen-java](https://github.com/tandemdude/sqlc-gen-java)
  - Writen in language: Go
  - Produces code for language: Java
  - Written by: tandemdude 
- [sqlc-gen-kotlin](https://github.com/sqlc-dev/sqlc-gen-kotlin)
  - Writen in language: Go
  - Produces code for language: Kotlin
  - Writen by: sqlc team
- [sqlc-gen-python](https://github.com/sqlc-dev/sqlc-gen-python)
  - Writen in language: Go
  - Produces code for language: Python
  - Writen by: sqlc team
- [sqlc-gen-zig](https://github.com/tinyzimmer/sqlc-gen-zig)
  - Writen in language: Go
  - Produces code for language: Zig
  - Written by: tinyzimmer
- [sqlc-plugin-php-dbal](https://github.com/lcarilla/sqlc-plugin-php-dbal)
  - Writen in language: Go
  - Produces code for language: php
  - Written by: lcarilla

## Plugins Developed in Other Languages

Plugins are compiled to WASM. This means they don't have to be developed in Go. [sqlc-gen-greeter](https://github.com/sqlc-dev/sqlc-gen-greeter) is an example of a plugin written in Rust.

- [sqlc-gen-csharp](https://github.com/DaredevilOSS/sqlc-gen-csharp)
  - Writen in language: C#
  - Produces code for language: C#
  - Written by: DaredevilOSS
  - [Introducing SQLC C# Plugin: A reverse ORM for .NET Developers](https://www.reddit.com/r/dotnet/comments/1hp6sa5/introducing_sqlc_c_plugin_a_reverse_orm_for_net/)
- [sqlc-gen-greeter](https://github.com/sqlc-dev/sqlc-gen-greeter)
  - Writen in language: Rust
  - Produces code for language: Rust
  - Writen by: sqlc team
  - Rust example plugin. last updated 2022/06.
- [sqlc-gen-ruby](https://github.com/DaredevilOSS/sqlc-gen-ruby)
  - Writen in language: Ruby
  - Produces code for language: Ruby
  - Written by: DaredevilOSS
- [sqlc-gen-typescript](https://github.com/sqlc-dev/sqlc-gen-typescript)
  - Writen in language: TypeScript
  - Produces code for language: TypeScript
  - Writen by: sqlc team

## Examples Provided by sqlc-plugin-dev

- code/0-sqlc
  - examples 0-0 through 0-3 are all from the sqlc docs. They are included here to make copying for your own testing easier. They are already adjusted to work with the databases provided.
  - /example0-0: [Developing sqlc](https://docs.sqlc.dev/en/latest/guides/development.html)
  - /example0-1: [Getting started with MySQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html)
  - /example0-2: [Getting started with PostgreSQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)
  - /example0-3: [Getting started with SQLite](https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html)
  - /example0-4: Configuration
- code/1-sqlc-gen-from-template
  - /example1-0: [sqlc-gen-from-template](https://github.com/fdietze/sqlc-gen-from-template)

### Other Examples

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

example of overrides https://docs.sqlc.dev/en/latest/reference/config.html
example of rules https://docs.sqlc.dev/en/latest/reference/config.html
example of kotlin and python configs

each plugin example

## TODO

update after release
https://github.com/sqlc-dev/sqlc/discussions/3077

https://github.com/sqlc-dev/sqlc/issues/3945


and update my own post about 3rd party plugins


## AI answers
question:
I want a better understanding of the sqlc library. Anaylyse the code at path E:\Working\sqlc-plugin-dev\code\ignore\sqlc and give me a high over view of how it works in 3 paragraphs or less
---
answer:
sqlc is a SQL compiler that generates type-safe code from SQL queries. The workflow is straightforward: you write SQL queries and schema definitions, sqlc parses and analyzes them, then generates code with type-safe interfaces to those queries in your chosen language. The core functionality involves parsing SQL files, analyzing the schema and queries to understand types and relationships, and then using this information to generate language-specific code that handles database interactions.

The architecture consists of several components: a parser that understands different SQL dialects (PostgreSQL, MySQL, SQLite), an analyzer that determines types and validates queries, and a code generation system that's been extended through plugins. These plugins (like sqlc-gen-go, sqlc-gen-python, sqlc-gen-typescript) allow sqlc to support multiple programming languages beyond its original Go implementation.

Recent versions have added features like managed databases for improved query analysis, linting capabilities through sqlc vet, and schema verification with sqlc verify. The configuration is handled through a YAML file that specifies database engine, input files, output location, and language-specific options. This plugin-based architecture allows the community to extend sqlc's capabilities to additional languages and frameworks.