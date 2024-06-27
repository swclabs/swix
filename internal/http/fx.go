package http

import (
	"swclabs/swipecore/internal/http/controller"
	"swclabs/swipecore/internal/http/router"

	"go.uber.org/fx"
)

var FxModule = fx.Options(
	controller.FxModule,
	router.FxModule,
	fx.Provide(
		NewServer,
	),
)
