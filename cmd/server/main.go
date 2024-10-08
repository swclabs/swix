/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"log"
	"swclabs/swix/app"
	_ "swclabs/swix/app/init"
	"swclabs/swix/internal/apis"
)

func main() {
	app := app.Builder(apis.NewApp)
	log.Fatal(app.Run())
}
