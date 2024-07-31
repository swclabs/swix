// Package router define tasks - queue
package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

// Base struct define the Base object
type Base struct {
	handlers *handler.Base
}

// NewBase creates a new Base object
func NewBase(handler *handler.Base) *Base {
	return &Base{
		handlers: handler,
	}
}

// Register register the queue
func (c *Base) Register(eng *worker.Engine) {
	eng.RegisterQueue(c.handlers.HandleHealthCheck)
}
