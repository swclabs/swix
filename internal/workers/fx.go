package workers

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/internal/workers/queue"
	"swclabs/swipecore/internal/workers/router"
	"swclabs/swipecore/pkg/lib/worker"

	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		queue.New,
		worker.New,
	),
	fx.Provide(
		handler.NewCommonConsume,
		router.NewCommon,

		handler.NewAccountManagementConsume,
		router.NewAccountManagement,

		NewWriter,
		NewAdapter,
	),
)
