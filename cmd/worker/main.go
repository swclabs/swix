/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"log"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/workers"
)

func main() {
	app := app.Builder(workers.NewApp)
	log.Fatal(app.Run())
}
