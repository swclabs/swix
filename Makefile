# Makefile for swipe-api
# a: Duc Hung Ho
# m: hunghd.dev@gmail.com

CHECK_MARK=\033[32m\u2714\033[0m
CROSS_MARK=\033[31m\u2716\033[0m
define build_ok
	printf "$(CHECK_MARK) [\033[34mbuild swipex\033[0m] %-15s   [ \033[32mOK\033[0m ]\n" $1
endef
define build_fail
	printf "$(CROSS_MARK) [\033[34mbuild swipex\033[0m] %-15s   [ \033[31mFAIL\033[0m ]\n" $1
endef
define swag_ok
	printf "$(CHECK_MARK) [\033[34mdocs swipex\033[0m] %-15s   [ \033[32mOK\033[0m ]\n" $1
endef
define swag_fail
	printf "$(CROSS_MARK) [\033[34mdocs swipex\033[0m] %-15s   [ \033[31mFAIL\033[0m ]\n" $1
endef

.PHONY: \
	build lint fmt\
	s w d \
	templ web \
	dev.build dev dev.down \
	compose.build compose compose.down \
	db db.down db.clean

build: cmd/main.go
	@go build -ldflags="-s -w" -o ./bin/swipe ./cmd \
	&& $(call build_ok,swipex) || $(call build_fail,swipex)

GENERATED_DIR=internal/cluster/proto
PROTO_DIR=internal/cluster/proto
proto: $(GENERATED_DIR) $(PROTO_DIR)
	@protoc --proto_path=./\
			--go_out=$(GENERATED_DIR) \
			--go-grpc_out=$(GENERATED_DIR) \
			--grpc-gateway_out=$(GENERATED_DIR) \
			$(PROTO_DIR)/*.proto

s: cmd/main.go
	@go run cmd/main.go s
w: cmd/main.go
	@go run cmd/main.go --start=worker
d: # Init swagger docs
	swag init
fmt:
	@find . -type f -name '*.go' -exec goimports -w {} \;

pre.build:
	@go mod download

build.article: cmd/main.go
	@go build -ldflags="-s -w" -o ./bin/article/swipe ./cmd/mod/article \
	&& $(call build_ok,article) \
	|| $(call build_fail,article)

build.authentication: cmd/main.go
	@go build -ldflags="-s -w" -o ./bin/authentication/swipe ./cmd/mod/authentication \
	&& $(call build_ok,authentication) \
	|| $(call build_fail,authentication)

build.products: cmd/main.go
	@go build -ldflags="-s -w" -o ./bin/products/swipe ./cmd/mod/products \
	&& $(call build_ok,products) \
	|| $(call build_fail,products)

build.purchase: cmd/main.go
	@go build -ldflags="-s -w" -o ./bin/purchase/swipe ./cmd/mod/purchase \
	&& $(call build_ok,purchase) \
	|| $(call build_fail,purchase)

lint: $(GOLANGCI) # Runs golangci-lint with predefined configuration
	@echo "Applying linter"
	@golangci-lint version
	golangci-lint run -c .golangci.yaml ./...

CONTAINER=./internal/apis/container
d.article:
	@swag init \
	--exclude $(CONTAINER)/authentication,$(CONTAINER)/classify,$(CONTAINER)/healthcheck,$(CONTAINER)/products,$(CONTAINER)/purchase \
	-o ./docs/article -q \
	&& $(call swag_ok,article) \
	|| $(call swag_fail,article)

d.authentication:
	@swag init \
	--exclude $(CONTAINER)/article,$(CONTAINER)/classify,$(CONTAINER)/healthcheck,$(CONTAINER)/products,$(CONTAINER)/purchase \
	-o ./docs/authentication -q \
	&& $(call swag_ok,authentication) \
	|| $(call swag_fail,authentication)

d.products:
	@swag init \
	--exclude $(CONTAINER)/article,$(CONTAINER)/classify,$(CONTAINER)/healthcheck,$(CONTAINER)/authentication,$(CONTAINER)/purchase \
	-o ./docs/products -q \
	&& $(call swag_ok,products) \
	|| $(call swag_fail,products)

d.purchase:
	@swag init \
	--exclude $(CONTAINER)/article,$(CONTAINER)/classify,$(CONTAINER)/healthcheck,$(CONTAINER)/authentication,$(CONTAINER)/products \
	-o ./docs/purchase -q \
	&& $(call swag_ok,purchase) \
	|| $(call swag_fail,purchase)

mod: d.article d.authentication d.products d.purchase pre.build build.article build.authentication build.products build.purchase

templ: # Generate templates
	@templ generate
web: cmd/www/main.go
	@go run cmd/www/main.go

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
	@rm -rf ./.swipe
