ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: deps docs run compose

default: run

deps:
	@echo "Installing dependencies..."
	@npm install

docs:
	@echo "Generating API docs..."
	@npx tsp compile ./docs

run:
	@echo "Running the application..."
	@go mod tidy
	@go run ./cmd/main.go

compose:
	@echo "Running Docker Compose..."
	@docker compose down
	@DB_PORT=${DB_PORT} DB_NAME=${DB_NAME} DB_USER=${DB_USER} DB_PASS=${DB_PASS} docker compose up -d
