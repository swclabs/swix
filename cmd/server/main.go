/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swipecore/boot"
	_ "swclabs/swipecore/boot/init"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/webapi"
)

func main() {
	if config.StageStatus == "prod" {
		boot.PrepareFor(boot.RestAPI | boot.ProdMode)
	} else {
		boot.PrepareFor(boot.RestAPI | boot.DebugMode)
	}
	app := boot.NewApp(boot.NewServer, webapi.NewAdapter)
	app.Run()
}
