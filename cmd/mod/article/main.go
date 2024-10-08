package main

import (
	"log"
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/container/article"
)

func main() {
	a := app.Builder(article.New)
	log.Fatal(a.Run())
}
