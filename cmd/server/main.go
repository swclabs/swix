/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swix/app"
	_ "swclabs/swix/app/init"
	"swclabs/swix/internal/apis"
)

func main() {
	app := app.App(apis.NewAPIServer)
	_ = app.Run()
}
