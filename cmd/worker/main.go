/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */

package main

import (
	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/workers"

	_ "swclabs/swipecore/boot/init"
)

func main() {
	boot.PrepareFor(boot.WorkerConsume)
	app := boot.NewApp(boot.NewServer, workers.NewAdapter)
	app.Run()
}
