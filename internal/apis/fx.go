// Package apis fx define
package apis

import (
	"swclabs/swix/internal/apis/controller"
	"swclabs/swix/internal/apis/router"
	"swclabs/swix/internal/apis/server"

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
