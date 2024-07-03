/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */

package main

import (
	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/http"

	"go.uber.org/fx"

	_ "swclabs/swipecore/boot/init"
)

func main() {
	app := fx.New(
		boot.FxRestModule,
		fx.Provide(
			http.NewAdapter,
			boot.NewServer,
		),
		fx.Invoke(boot.Main),
	)
	app.Run()
}
