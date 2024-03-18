# Makefile for swipe-api
# a: Duc Hung Ho
# m: hunghd.dev@gmail.com

build:
	go mod download
	go build -ldflags="-s -w" -o ./bin/swipe ./cmd
s: 
	go run cmd/main.go s
w:
	go run cmd/main.go w
m:
	go run cmd/main.go m up
d: 
	swag init
dev-b: 
	docker compose -f docker-compose.dev.yml up --build -d
dev: 
	docker compose -f docker-compose.dev.yml up -d
dev-down:
	docker compose -f docker-compose.dev.yml down
all-b:
	docker compose up --build -d
all:
	docker compose up -d
all-down:
	docker compose down
db:
	docker compose -f docker-compose.db.yml up -d
db-down:
	docker compose -f docker-compose.db.yml down