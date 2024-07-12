// Package router define tasks - queue
package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

// Common struct define the Common object
type Common struct {
	handlers *handler.Common
}

// NewCommon creates a new Common object
func NewCommon(handler *handler.Common) *Common {
	return &Common{
		handlers: handler,
	}
}

// Register register the queue
func (c *Common) Register(eng *worker.Engine) {
	eng.RegisterQueue(c.handlers.HandleHealthCheck)
}
