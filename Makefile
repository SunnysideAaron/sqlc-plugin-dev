build:
# @docker compose build code --no-cache
	@docker compose build code

bash:
	@docker compose run --remove-orphans --service-ports code bash

# This only works once but will help get things rolling
git:
	@git clone https://github.com/sqlc-dev/sqlc ./code/ignore/sqlc
	@git clone https://github.com/sqlc-dev/plugin-sdk-go ./code/ignore/plugin-sdk-go
	@git clone https://github.com/sqlc-dev/sqlc-gen-go ./code/ignore/sqlc-gen-go

.PHONY: build bash
