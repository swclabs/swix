package main

import (
	"example/komposervice/api"
	_ "example/komposervice/docs"
	"example/komposervice/internal/config"
	"fmt"
	"log"
)

// @title Microservice API Documentation
// @version 1.0.0
// @description This is a documentation for the Microservice API
// @host
// @basePath /
func main() {
	server := api.NewServer()
	if err := server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port)); err != nil {
		log.Fatal(err)
	}
}
