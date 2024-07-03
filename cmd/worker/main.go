/**
 * A: Ho Duc Hung <hunghd.dev@gmail.com> @kyeranyo
 * This is Graduation project in computer science
 * 2023 - Ho Chi Minh City University of Technology, VNUHCM
 */

package main

import (
	"swclabs/swipecore/boot"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		boot.FxWorkerModule,
		fx.Provide(
			boot.NewWorker,
		),
		fx.Invoke(boot.Main),
	)
	app.Run()
}
