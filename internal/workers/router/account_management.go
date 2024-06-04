package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/lib/worker"
)

type AccountManagements struct {
	handlers *handler.AccountManagement
}

func NewAccountManagement(handlers *handler.AccountManagement) *AccountManagements {
	return &AccountManagements{
		handlers: handlers,
	}
}

func (router *AccountManagements) Register(eng *worker.Engine) {
	eng.RegisterQueue(router.handlers.HandleOAuth2SaveUser)
	eng.RegisterQueue(router.handlers.HandleSignUp)
	eng.RegisterQueue(router.handlers.HandleUpdateUserInfo)
}
