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
	var flag = boot.WebAPI | boot.DebugMode
	if config.StageStatus != "dev" {
		flag = boot.WebAPI | boot.ProdMode
	}
	app := boot.NewApp(flag, boot.NewServer, webapi.NewAdapter)
	app.Run()
}
