package main

import (
	"example/swiftcart/api"
	"example/swiftcart/internal/config"
	"fmt"
	"log"
)

func main() {
	server := api.NewServer()
	if err := server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port)); err != nil {
		log.Fatal(err)
	}
}
