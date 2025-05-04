code-build:
# @docker compose build code --no-cache
	@docker compose build code

code-bash:
	@docker compose run --remove-orphans --service-ports code bash

# This only works once but will help get things rolling
code-git:
	@git clone https://github.com/sqlc-dev/sqlc ./code/ignore/sqlc
	@git clone https://github.com/sqlc-dev/plugin-sdk-go ./code/ignore/plugin-sdk-go
	@git clone https://github.com/sqlc-dev/sqlc-gen-go ./code/ignore/sqlc-gen-go

code-git2:
	@git clone https://github.com/sqlc-dev/sqlc-gen-python ./code/ignore/sqlc-gen-python
	@git clone https://github.com/sqlc-dev/sqlc-gen-kotlin ./code/ignore/sqlc-gen-kotlin
	@git clone https://github.com/sqlc-dev/sqlc-gen-typescript ./code/ignore/sqlc-gen-typescript

	@git clone https://github.com/DaredevilOSS/sqlc-gen-csharp ./code/ignore/sqlc-gen-csharp
	@git clone https://github.com/kaashyapan/sqlc-gen-fsharp ./code/ignore/sqlc-gen-fsharp
	@git clone https://github.com/tandemdude/sqlc-gen-java ./code/ignore/sqlc-gen-java
	@git clone https://github.com/lcarilla/sqlc-plugin-php-dbal ./code/ignore/sqlc-plugin-php-dbal
	@git clone https://github.com/DaredevilOSS/sqlc-gen-ruby ./code/ignore/sqlc-gen-ruby
	@git clone https://github.com/tinyzimmer/sqlc-gen-zig ./code/ignore/sqlc-gen-zig

	@git clone https://github.com/sqlc-dev/sqlc-gen-greeter ./code/ignore/sqlc-gen-greeter
	@git clone https://github.com/fdietze/sqlc-gen-from-template ./code/ignore/sqlc-gen-from-template
	@git clone https://github.com/walterwanderley/sqlc-gen-go-server ./code/ignore/sqlc-gen-go-server

###########################################################

my-up:
	@docker compose up --remove-orphans mysql -d

my-down:
	@docker compose down mysql

###########################################################

post-up:
	@docker compose up --remove-orphans postgresql -d

post-down:
	@docker compose down postgresql

###########################################################

.PHONY: code-build code-bash code-git
