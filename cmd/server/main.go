/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */
package main

import (
	"swclabs/swix/boot"
	_ "swclabs/swix/boot/init"
	"swclabs/swix/internal/apis"
	"swclabs/swix/internal/config"
)

func main() {
	var flag = boot.APIs | boot.DebugMode
	if config.StageStatus != "dev" {
		flag = boot.APIs | boot.ProdMode
	}
	app := boot.NewApp(flag, boot.NewServer, apis.NewAdapter)
	app.Run()
}
