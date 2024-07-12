package wapi

import (
	"swclabs/swipecore/internal/wapi/controller"
	"swclabs/swipecore/internal/wapi/router"

	"go.uber.org/fx"
)

var FxModule = fx.Options(
	controller.FxModule,
	router.FxModule,
	fx.Provide(
		NewServer,
	),
)
