ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: deps docs run

default: run

deps:
	@echo "Installing dependencies..."
	@npm install

docs:
	@echo "Generating API docs..."
	@npx tsp compile ./docs

run:
	@echo "Running the application..."
	@go work sync
	@cd ./app; go mod tidy
	@go run ./app/cmd/main.go
