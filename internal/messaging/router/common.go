package router

import (
	"swclabs/swiftcart/internal/messaging/handler"
	"swclabs/swiftcart/internal/messaging/queue"
	"swclabs/swiftcart/pkg/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(queue.WorkerHealthCheck, common.HandleHealthCheck)
}
