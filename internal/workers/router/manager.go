// Package router define tasks - queue
package router

import (
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/pkg/lib/worker"
)

// IManager interface for Manager objects
type IManager interface {
	IRouter
}

// NewManager creates a new Manager object
func NewManager(handlers handler.IManager) IManager {
	return &Manager{
		handlers: handlers,
	}
}

// Manager struct define the Manager object
type Manager struct {
	handlers handler.IManager
}

// Register register the queue
func (router *Manager) Register(eng worker.IEngine) {
	eng.RegisterQueue(router.handlers.HandleOAuth2SaveUser)
	eng.RegisterQueue(router.handlers.HandleSignUp)
	eng.RegisterQueue(router.handlers.HandleUpdateUserInfo)
}
