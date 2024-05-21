package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(common.HandleHealthCheck)
}
