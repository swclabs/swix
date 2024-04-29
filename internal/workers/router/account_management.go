package router

import (
	"swclabs/swipecore/internal/workers/handler"
	"swclabs/swipecore/pkg/tools/worker"
)

var _AccountManagement = handler.NewAccountManagement()

func AccountManagement(eng *worker.Engine) {
	eng.Queue(_AccountManagement.HandleOAuth2SaveUser)
	eng.Queue(_AccountManagement.HandleSignUp)
	eng.Queue(_AccountManagement.HandleUpdateUserInfo)
}
