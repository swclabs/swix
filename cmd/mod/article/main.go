package main

import (
	"log"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/container/article"
)

func main() {
	a := app.Builder(article.New)
	log.Fatal(a.Run())
}
