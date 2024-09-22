package main

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/container/article"
)

func main() {
	a := app.Builder(article.New)
	_ = a.Run()
}
