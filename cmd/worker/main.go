/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swix/boot"
	_ "swclabs/swix/boot/init"
	"swclabs/swix/internal/workers"
)

func main() {
	app := boot.App(workers.NewWorkerNode)
	_ = app.Run()
}
