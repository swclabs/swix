package router

import (
	"github.com/swclabs/swipe-server/internal/broker/handler"
	"github.com/swclabs/swipe-server/pkg/tools/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(common.HandleHealthCheck)
}
