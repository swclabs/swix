package router

import (
	"swclabs/swipe-api/internal/workers/handler"
	"swclabs/swipe-api/pkg/tools/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(common.HandleHealthCheck)
}
