ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: run

default: run

run:
	@echo "Running the application..."
	@go mod tidy
	@go run ./cmd/main.go
