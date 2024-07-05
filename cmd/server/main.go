/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */

package main

import (
	"swclabs/swipecore/boot"
	_ "swclabs/swipecore/boot/init"
	"swclabs/swipecore/internal/http"
)

func main() {
	app := boot.NewApp(boot.NewServer, http.NewAdapter)
	app.Run()
}
