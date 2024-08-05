/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swix/boot"
	_ "swclabs/swix/boot/init"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/webapi"
)

func main() {
	if config.StageStatus == "prod" {
		boot.PrepareFor(boot.WebAPI | boot.ProdMode)
	} else {
		boot.PrepareFor(boot.WebAPI | boot.DebugMode)
	}
	app := boot.NewApp(boot.NewServer, webapi.NewAdapter)
	app.Run()
}
