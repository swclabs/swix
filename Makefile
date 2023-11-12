build:
	go build -ldflags="-s -w" -o ./bin/exe ./cmd

s:
	go run cmd/main.go s

w:
	go run cmd/main.go w

j:
	go run cmd/main.go j

m:
	go run cmd/main.go m
	
doc:
	swag init