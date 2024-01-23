build: 
	go build -ldflags="-s -w" -o ./bin/swipe ./cmd
s: 
	go run cmd/main.go s
w:
	go run cmd/main.go w
m:
	go run cmd/main.go m up
d: 
	swag init
up-b: 
	docker compose -f docker-compose.dev.yml up --build -d
up: 
	docker compose -f docker-compose.dev.yml up -d
env-up:
	docker compose -f docker-compose.env.yml up -d
down: 
	docker compose down