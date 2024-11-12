package main

import (
	"log"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/apis/container/article"

	_ "github.com/swclabs/swipex/docs/article"
)

func main() {
	a := app.Builder(article.New)
	log.Fatal(a.Run())
}
