/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"log"

	"github.com/swclabs/swipex/app"
	_ "github.com/swclabs/swipex/docs"

	"github.com/swclabs/swipex/internal/apis"
)

func main() {
	app := app.Builder(apis.NewApp)
	log.Fatal(app.Run())
}
