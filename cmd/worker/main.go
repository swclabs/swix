// Author: - Ho Duc Hung : @kieranhoo
// 		   - Nguyen Van Khoa: @anthony2704
// This is Graduation project in computer science
// 2023 - Ho Chi Minh City University of Technology, VNUHCM

package main

import (
	"go.uber.org/fx"
	"swclabs/swipecore/boot"
)

func main() {
	app := fx.New(
		boot.FxWorkerModule,
		fx.Provide(
			boot.NewWorker,
		),
		fx.Invoke(boot.StartWorker),
	)
	app.Run()
}
