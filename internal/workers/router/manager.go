// Package router define tasks - queue
package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

// Manager struct define the Manager object
type Manager struct {
	handlers *handler.Manager
}

// NewManager creates a new Manager object
func NewManager(handlers *handler.Manager) *Manager {
	return &Manager{
		handlers: handlers,
	}
}

// Register register the queue
func (router *Manager) Register(eng *worker.Engine) {
	eng.RegisterQueue(router.handlers.HandleOAuth2SaveUser)
	eng.RegisterQueue(router.handlers.HandleSignUp)
	eng.RegisterQueue(router.handlers.HandleUpdateUserInfo)
}
