// Package router define tasks - queue
package base

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/server"
	"swclabs/swix/pkg/lib/worker"
)

var _ = app.Router(NewRouter)

// NewRouter creates a new Base object
func NewRouter(handler IHandler) IRouter {
	return &Router{
		handlers: handler,
	}
}

// IRouter interface for Base objects
type IRouter interface {
	server.IRouter
}

// Router struct define the Router object
type Router struct {
	handlers IHandler
}

// Register register the queue
func (b *Router) Register(eng worker.IEngine) {
	eng.HandlerFunc(b.handlers.HandleHealthCheck)
}
