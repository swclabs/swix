// Package manager define tasks - queue
package manager

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/server"
	"swclabs/swix/pkg/lib/worker"
)

var _ = app.Router(NewRouter)

// NewManager creates a new Manager object
func NewRouter(handlers IHandler) IRouter {
	return &Router{
		handlers: handlers,
	}
}

// IManager interface for Manager objects
type IRouter interface {
	server.IRouter
}

// Manager struct define the Manager object
type Router struct {
	handlers IHandler
}

// Register register the queue
func (router *Router) Register(eng worker.IEngine) {
	eng.RegisterQueue(router.handlers.HandleOAuth2SaveUser)
	eng.RegisterQueue(router.handlers.HandleSignUp)
	eng.RegisterQueue(router.handlers.HandleUpdateUserInfo)
}
