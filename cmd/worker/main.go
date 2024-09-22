/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swix/app"
	_ "swclabs/swix/app/init"
	"swclabs/swix/internal/workers"
)

func main() {
	app := app.Builder(workers.NewApp)
	_ = app.Run()
}
