# Makefile for swipe-api
# a: Duc Hung Ho
# m: hunghd.dev@gmail.com

.PHONY: \
	build lint fmt\
	s w d \
	templ web \
	dev.build dev dev.down \
	compose.build compose compose.down \
	db db.down db.clean


build: cmd/main.go
	@go mod download
	@go build -ldflags="-s -w" -o ./bin/swipe ./cmd
	@echo "build swipe ... done"

lint: $(GOLANGCI) # Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	@golangci-lint version
	golangci-lint run -c .golangci.yaml ./...
	

fmt:
	@find . -type f -name '*.go' -exec goimports -w {} \;

GENERATED_DIR=internal/core/proto
PROTO_DIR=internal/core/proto
proto: $(GENERATED_DIR) $(PROTO_DIR)
	@protoc --proto_path=./ \
			--go_out=$(GENERATED_DIR) \
			--go-grpc_out=$(GENERATED_DIR) \
			--grpc-gateway_out $(GENERATED_DIR) \
			$(PROTO_DIR)/*.proto


s: cmd/main.go
	@go run cmd/main.go s
w: cmd/main.go
	@go run cmd/main.go w
d: # Init swagger docs
	@swag init
templ: # Generate templates
	@templ generate
web: pkg/web/web.go
	@go run pkg/web/web.go

dev.build: docker-compose.yml Dockerfile
	@docker compose up --build -d
dev: docker-compose.yml Dockerfile
	@docker compose up -d
dev.down: docker-compose.yml Dockerfile
	@docker compose down

db: docker-compose.db.yml
	@docker compose -f \
	docker-compose.db.yml \
	up -d
db.down: docker-compose.db.yml
	@docker compose -f \
	docker-compose.db.yml \
	down
db.clean: docker-compose.db.yml
	@rm -rf ./boot/.swipe
