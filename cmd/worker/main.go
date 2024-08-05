/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swix/boot"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/workers"

	_ "swclabs/swix/boot/init"
)

func main() {
	if config.StageStatus == "prod" {
		boot.PrepareFor(boot.Worker | boot.ProdMode)
	} else {
		boot.PrepareFor(boot.Worker | boot.DebugMode)
	}
	app := boot.NewApp(boot.NewServer, workers.NewAdapter)
	app.Run()
}
