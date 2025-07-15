.PHONY: deps docs

default: docs

deps:
	@echo "Installing dependencies..."
	@npm install

docs:
	@echo "Generating API docs..."
	@npx tsp compile ./docs
