// Package router define tasks - queue
package router

import (
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/internal/workers/server"
)

// IBase interface for Base objects
type IBase interface {
	IRouter
}

// NewBase creates a new Base object
func NewBase(handler handler.IBaseHandler) IBase {
	return &Base{
		handlers: handler,
	}
}

// Base struct define the Base object
type Base struct {
	handlers handler.IBaseHandler
}

// Register register the queue
func (c *Base) Register(eng server.IWorker) {
	eng.RegisterQueue(c.handlers.HandleHealthCheck)
}
