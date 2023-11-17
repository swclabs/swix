package main

import (
	"example/swiftcart/api"
	"log"
)

func main() {
	w := api.NewWorker(10)
	if err := w.Run(); err != nil {
		log.Fatal(err)
	}
}
