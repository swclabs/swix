package main

import (
	"log"
	"swclabs/swix/pkg/infra/db"
)

func main() {
	if err := db.MigrateUp(); err != nil {
		log.Fatal(err)
	}
}
