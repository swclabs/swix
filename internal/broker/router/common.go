package router

import (
	"github.com/swclabs/swipe-api/internal/broker/handler"
	"github.com/swclabs/swipe-api/internal/broker/queue"
	"github.com/swclabs/swipe-api/pkg/tools/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(queue.WorkerHealthCheck, common.HandleHealthCheck)
}
