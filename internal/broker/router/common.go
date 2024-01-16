package router

import (
	"swclabs/swiftcart/internal/broker/handler"
	"swclabs/swiftcart/internal/broker/queue"
	"swclabs/swiftcart/pkg/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(queue.WorkerHealthCheck, common.HandleHealthCheck)
}
