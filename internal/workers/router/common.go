// Package router define tasks - queue
package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

type Common struct {
	handlers *handler.Common
}

func NewCommon(handler *handler.Common) *Common {
	return &Common{
		handlers: handler,
	}
}

func (c *Common) Register(eng *worker.Engine) {
	eng.RegisterQueue(c.handlers.HandleHealthCheck)
}
