// Package router define tasks - queue
package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

// AccountManagements struct define the AccountManagements object
type AccountManagements struct {
	handlers *handler.AccountManagement
}

// NewAccountManagement creates a new AccountManagements object
func NewAccountManagement(handlers *handler.AccountManagement) *AccountManagements {
	return &AccountManagements{
		handlers: handlers,
	}
}

// Register register the queue
func (router *AccountManagements) Register(eng *worker.Engine) {
	eng.RegisterQueue(router.handlers.HandleOAuth2SaveUser)
	eng.RegisterQueue(router.handlers.HandleSignUp)
	eng.RegisterQueue(router.handlers.HandleUpdateUserInfo)
}
