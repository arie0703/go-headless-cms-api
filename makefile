# Makefile
.PHONY: up down build migrate create-migration

up:
	docker compose up

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker-compose build --progress=plain

down:
	docker compose down

build:
	go build -o app .

migrate:
	goose -dir db/migrations mysql "$$DB_DSN" up

create-migration:
	@read -p "Migration name: " name; \
	goose -dir db/migrations create $$name sql
