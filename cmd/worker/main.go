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
	var flag = boot.APIs | boot.DebugMode
	if config.StageStatus != "dev" {
		flag = boot.APIs | boot.ProdMode
	}
	app := boot.NewApp(flag, boot.NewWorker, workers.NewAdapter)
	app.Run()
}
