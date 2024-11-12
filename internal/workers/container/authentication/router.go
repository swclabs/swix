// Package manager define tasks - queue
package authentication

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/workers/server"
	"github.com/swclabs/swipex/pkg/lib/worker"
)

var _ = app.Router(NewRouter)

// NewRouter creates a new Manager object
func NewRouter(handlers *Handler) IRouter {
	return &Router{
		handlers: handlers,
	}
}

// IRouter interface for Manager objects
type IRouter interface {
	server.IRouter
}

// Router struct define the Manager object
type Router struct {
	handlers *Handler
}

// Register register the queue
func (router *Router) Register(eng worker.IEngine) {
	eng.HandlerFunc("auth.SignUp", router.handlers.SignUp)
	eng.HandlerFunc("auth.OAuth2SaveUser", router.handlers.OAuth2SaveUser)
	eng.HandlerFunc("auth.UpdateUserInfo", router.handlers.UpdateUserInfo)
}
