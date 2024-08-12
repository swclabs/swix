// Package webapi fx define
package webapi

import (
	"swclabs/swix/internal/webapi/controller"
	"swclabs/swix/internal/webapi/router"
	"swclabs/swix/internal/webapi/server"

	"go.uber.org/fx"
)

// FxModule represents constructer of controller, router and server
var FxModule = fx.Options(
	controller.FxModule,
	router.FxModule,
	fx.Provide(
		server.New,
	),
)
