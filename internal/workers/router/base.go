// Package router define tasks - queue
package router

import (
	"swclabs/swix/boot"
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/pkg/lib/worker"
)

var _ = boot.Router(NewBase)

// NewBase creates a new Base object
func NewBase(handler handler.IBaseHandler) IBase {
	return &Base{
		handlers: handler,
	}
}

// IBase interface for Base objects
type IBase interface {
	IRouter
}

// Base struct define the Base object
type Base struct {
	handlers handler.IBaseHandler
}

// Register register the queue
func (c *Base) Register(eng worker.IEngine) {
	eng.RegisterQueue(c.handlers.HandleHealthCheck)
}
