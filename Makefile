# Makefile for swipe-api
# a: Duc Hung Ho
# m: hunghd.dev@gmail.com

.PHONY: \
	build lint fmt\
	s w d \
	templ web \
	dev-build dev dev-down \
	compose-build compose compose-down \
	dev-db dev-db-down


build: # Build swipe binaries
	@echo "build swipe ...."
	@go mod download
	@go build -ldflags="-s -w" -o ./bin/swipe ./cmd
	@echo "done"

lint: $(GOLANGCI) ## Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

fmt:
	@find . -type f -name '*.go' -exec goimports -w {} \;

GENERATED_DIR=internal/core/proto
PROTO_DIR=internal/core/proto
proto:
	@protoc --proto_path=./ \
			--go_out=$(GENERATED_DIR) \
			--go-grpc_out=$(GENERATED_DIR) \
			--grpc-gateway_out $(GENERATED_DIR) \
			$(PROTO_DIR)/*.proto


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
