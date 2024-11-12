// Package router define tasks - queue
package healthcheck

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/workers/server"
	"github.com/swclabs/swipex/pkg/lib/worker"
)

var _ = app.Router(NewRouter)

// NewRouter creates a new Base object
func NewRouter(handler *Handler) IRouter {
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
	handlers *Handler
}

// Register register the queue
func (b *Router) Register(eng worker.IEngine) {
	eng.HandlerFunc("healthcheck.WorkerCheckResult", b.handlers.WorkerCheckResult)
}
