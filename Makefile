build:
	go build -ldflags="-s -w" -o ./bin/exe ./cmd

s:
	go run cmd/main.go s

w:
	go run cmd/main.go w

j:
	go run cmd/main.go j

m:
	go run cmd/main.go m up
	
swag:
	swag init

t:
	go test -v ./test

env-up:
	docker compose -f docker-compose.env.yml up -d
env-down:
	docker compose down