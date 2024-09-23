// Package purchase define tasks - queue
package purchase

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/server"
	"swclabs/swix/pkg/lib/worker"
)

var _ = app.Router(NewRouter)

// NewRouter creates a new Purchase object
func NewRouter(handler IHandler) IRouter {
	return &Router{
		handler: handler,
	}
}

// IRouter interface for Purchase objects
type IRouter interface {
	server.IRouter
}

// Router struct define the Router object
type Router struct {
	handler IHandler
}

// Register implements IPurchase.
func (r *Router) Register(eng worker.IEngine) {
	eng.HandlerFunc(r.handler.HandleAddToCart)
}
