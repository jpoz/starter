#!make
include .env
export $(shell sed 's/=.*//' .env)

CGO_ENABLED?=0
GOCMD=CGO_ENABLED=$(CGO_ENABLED) go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet

MIGRATION_NAME ?=you_need_to_fill_me_in

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all
all: help

.PHONY: install

## Initial
install: ## install dependencies
	brew install golang-migrate sqlc
	go install github.com/cespare/reflex@latest
	go install gorm.io/gen/tools/gentool@latest
	go mod download

.PHONY: dev
## Dev
dev: ## Run dependencies and server (with reflex)
	docker-compose up -d
	reflex -r '\.(go|html)$$' -s make server

.PHONY: server
server: ## Run a local verion of the server
	${GOCMD} run cmd/server/main.go

.PHONY: cleanup
cleanup: ## Cleanup the docker-compose
	docker-compose down


## Demo
demo_temp: ## Run a demo of the ai
	go run demo/ai/main.go


## Gen
gen: gen_db gen_graphql ## Generate all the things

gen_graphql:
	${GOCMD} run github.com/99designs/gqlgen generate --config gqlgen.yml

gen_db: ## Generate models & queries
	${GOCMD} run dev/gen_query/main.go

## Database
dump_schema: ## Dump the database schema
	pg_dump --schema-only ${DATABASE_URL} > pkg/db/sql/schema.sql

migrate_up: ## Migrate the database up
	migrate -source file://pkg/db/migrations -database ${DATABASE_URL} up

migrate_down: ## Migrate the database down one step
	migrate -source file://pkg/db/migrations -database ${DATABASE_URL} down 1

create_migration: ## Create a new migration
	migrate create -ext sql -dir pkg/db/migrations -seq ${MIGRATION_NAME}

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
