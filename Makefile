# Makefile for swipe-api
# a: Duc Hung Ho
# m: hunghd.dev@gmail.com

.PHONY: help build lint s w d templ web dev-build dev dev-down compose-build compose compose-down dev-db dev-db-down


build: # Build swipe binaries
	@echo "build swipe ...."
	@go mod download
	@go build -ldflags="-s -w" -o ./bin/swipe ./cmd
	@echo "done"

lint: $(GOLANGCI) ## Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

s: # Start server
	@go run cmd/main.go s
w: # Start workers
	@go run cmd/main.go w
d: # Init swagger docs
	@swag init
templ: # Generate templates
	@templ generate
web: # Run web
	@go run pkg/web/web.go

dev-build: # Build and start containers with docker-compose for development
	@docker compose -f \
	docker-compose.dev.yml \
	up --build -d
dev: # Start containers with docker-compose for development
	@docker compose -f \
	docker-compose.dev.yml \
	up -d
dev-down: # Stop and remove containers with docker-compose for development
	@docker compose -f \
	docker-compose.dev.yml \
	down

compose-build: # Build and start containers with docker-compose
	@docker compose up \
	--build -d
compose: # Start containers with docker-compose
	@docker compose up -d
compose-down: # Stop and remove containers with docker-compose
	@docker compose down

dev-db: # Start database containers with docker-compose
	@docker compose -f \
	docker-compose.db.yml \
	up -d
dev-db-down: # Stop and remove database containers with docker-compose
	@docker compose -f \
	docker-compose.db.yml \
	down

# Help target
help: ## Show this help message
	@echo "Usage: make [TARGET]"
	@echo ""
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Examples:"
	@echo "  make build        # Build swipe binaries"
	@echo "  make lint         # Runs golangci-lint with predefined configuration"
	@echo "  make s            # Start server"
	@echo "  make w            # Start workers"
	@echo "  make d            # Init swagger docs"
	@echo "  make templ        # Generate templates"
	@echo "  make web          # Run web"
	@echo "  make dev-build    # Build and start containers with docker-compose for development"
	@echo "  make dev          # Start containers with docker-compose for development"
	@echo "  make dev-down     # Stop and remove containers with docker-compose for development"
	@echo "  make compose-build # Build and start containers with docker-compose"
	@echo "  make compose      # Start containers with docker-compose"
	@echo "  make compose-down # Stop and remove containers with docker-compose"
	@echo "  make dev-db       # Start database containers with docker-compose"
	@echo "  make dev-db-down  # Stop and remove database containers with docker-compose"