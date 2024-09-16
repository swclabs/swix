// Package apis fx define
package server

import (
	"swclabs/swix/internal/apis/controller"
	"swclabs/swix/internal/apis/router"

	"go.uber.org/fx"
)

// FxModule represents constructer of controller, router and server
var FxModule = fx.Options(
	controller.FxModule,
	router.FxModule,
)
