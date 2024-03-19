package router

import (
	"github.com/swclabs/swipe-api/internal/writer/handler"
	"github.com/swclabs/swipe-api/pkg/tools/worker"
)

var common = handler.NewCommonHandler()

func Common(eng *worker.Engine) {
	eng.Queue(common.HandleHealthCheck)
}
