// Package webapi fx define
package webapi

import (
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/internal/webapi/router"

	"go.uber.org/fx"
)

// FxModule represents constructer of controller, router and server
var FxModule = fx.Options(
	controller.FxModule,
	router.FxModule,
	fx.Provide(
		NewServer,
	),
)
