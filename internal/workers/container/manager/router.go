// Package manager define tasks - queue
package manager

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/server"
	"swclabs/swix/pkg/lib/worker"
)

var _ = app.Router(NewRouter)

// NewManager creates a new Manager object
func NewRouter(handlers *Handler) IRouter {
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
	handlers *Handler
}

// Register register the queue
func (router *Router) Register(eng worker.IEngine) {
	eng.HandlerFunc("manager.SignUp", router.handlers.SignUp)
	eng.HandlerFunc("manager.OAuth2SaveUser", router.handlers.OAuth2SaveUser)
	eng.HandlerFunc("manager.UpdateUserInfo", router.handlers.UpdateUserInfo)
}
