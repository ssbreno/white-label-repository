.PHONY: help dev dev-frontend dev-backend build test test-backend test-frontend \
       lint format docker-up docker-down docker-build docker-logs \
       db-shell db-reset install clean graph

# Default target
help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ──────────────────────────────────────────────────
# Development
# ──────────────────────────────────────────────────

install: ## Install all dependencies (npm + Go modules)
	npm install
	cd apps/backend && go mod download

dev: ## Start all services in development mode
	npm run dev

dev-frontend: ## Start only the frontend
	npm run dev:frontend

dev-backend: ## Start only the backend (requires PostgreSQL)
	npm run dev:backend

# ──────────────────────────────────────────────────
# Build
# ──────────────────────────────────────────────────

build: ## Build all applications
	npm run build

# ──────────────────────────────────────────────────
# Test & Quality
# ──────────────────────────────────────────────────

test: ## Run all tests
	npm run test

test-backend: ## Run Go backend tests
	cd apps/backend && go test -race ./...

test-frontend: ## Run frontend tests
	cd apps/frontend && npm test

lint: ## Lint all projects
	npm run lint

format: ## Format code with Prettier
	npm run format

# ──────────────────────────────────────────────────
# Docker
# ──────────────────────────────────────────────────

docker-up: ## Start all Docker services
	docker-compose up -d

docker-down: ## Stop all Docker services
	docker-compose down

docker-build: ## Build and start Docker services
	docker-compose up -d --build

docker-logs: ## Follow Docker service logs
	docker-compose logs -f

# ──────────────────────────────────────────────────
# Database
# ──────────────────────────────────────────────────

db-shell: ## Open PostgreSQL shell
	docker-compose exec postgres psql -U postgres -d whitelabel_db

db-reset: ## Reset database (drop and recreate)
	docker-compose down -v
	docker-compose up postgres -d
	@echo "Waiting for PostgreSQL..."
	@sleep 3
	@echo "Database reset. Run 'make dev-backend' to apply migrations."

# ──────────────────────────────────────────────────
# Utility
# ──────────────────────────────────────────────────

clean: ## Clean build artifacts and caches
	rm -rf node_modules/.cache .next dist apps/backend/dist .nx/cache coverage

graph: ## Open Nx project dependency graph
	npm run graph
