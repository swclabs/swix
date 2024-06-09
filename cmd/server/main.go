/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kieranhoo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */

package main

import (
	"swclabs/swipecore/boot"
	"swclabs/swipecore/boot/adapter"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		boot.FxRestModule,
		fx.Provide(
			adapter.NewAdapter,
			boot.NewServer,
		),
		fx.Invoke(boot.StartServer),
	)
	app.Run()
}
