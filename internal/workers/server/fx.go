// Package workers define worker consume
package server

import (
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/internal/workers/router"

	"go.uber.org/fx"
)

// FxModule module of package workers
var FxModule = fx.Options(
	fx.Provide(
		handler.NewBase,
		handler.NewManager,
		handler.NewPurchase,

		router.NewBase,
		router.NewManager,
		router.NewPurchase,
	),
)
