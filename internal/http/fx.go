package http

import (
	"go.uber.org/fx"
	"swclabs/swipecore/internal/http/controller"
	"swclabs/swipecore/internal/http/router"
)

var FxModule = fx.Options(
	controller.FxModule,
	router.FxModule,
	fx.Provide(
		NewServer,
	),
)
