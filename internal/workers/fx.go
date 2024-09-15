// Package workers define worker consume
package workers

import (
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/internal/workers/queue"
	"swclabs/swix/internal/workers/router"
	"swclabs/swix/internal/workers/server"
	"swclabs/swix/pkg/lib/worker"

	"go.uber.org/fx"
)

// FxModule module of package workers
var FxModule = fx.Options(
	fx.Provide(
		queue.New,
		worker.New,
	),
	fx.Provide(
		handler.NewBase,
		handler.NewManager,
		handler.NewPurchase,

		router.NewBase,
		router.NewManager,
		router.NewPurchase,
	),
	fx.Provide(
		server.New,
	),
)
