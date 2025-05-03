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
