package main

import (
	"log"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/container/article"

	_ "swclabs/swipex/docs/article"
)

func main() {
	a := app.Builder(article.New)
	log.Fatal(a.Run())
}
