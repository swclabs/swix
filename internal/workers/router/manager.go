// Package router define tasks - queue
package router

import (
	"swclabs/swix/boot"
	"swclabs/swix/internal/workers/handler"
	"swclabs/swix/pkg/lib/worker"
)

var _ = boot.Router(NewManager)

// NewManager creates a new Manager object
func NewManager(handlers handler.IManager) IManager {
	return &Manager{
		handlers: handlers,
	}
}

// IManager interface for Manager objects
type IManager interface {
	IRouter
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
