// Package wapi fx define
package wapi

import (
	"swclabs/swipecore/internal/wapi/controller"
	"swclabs/swipecore/internal/wapi/router"

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
