.PHONY: build run test test-coverage lint lint-fix docker-build clean help

BINARY_NAME=server
BUILD_DIR=dist
MAIN_PATH=./cmd/server

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Compile the server binary
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

run: ## Run the server in development mode
	go run $(MAIN_PATH)

test: ## Run all unit tests
	go test -race -short ./...

test-coverage: ## Run tests with coverage report
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint: ## Run golangci-lint
	golangci-lint run ./...

lint-fix: ## Run golangci-lint with auto-fix
	golangci-lint run --fix ./...

docker-build: ## Build the Docker image
	docker build -t white-label-backend .

clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR) coverage.out coverage.html
