package router

import (
	"swclabs/swipe-api/internal/workers/handler"
	"swclabs/swipe-api/pkg/tools/worker"
)

var _AccountManagement = handler.NewAccountManagement()

func AccountManagement(eng *worker.Engine) {
	eng.Queue(_AccountManagement.HandleOAuth2SaveUser)
	eng.Queue(_AccountManagement.HandleSignUp)
	eng.Queue(_AccountManagement.HandleUpdateUserInfo)
}
