package router

import (
	"github.com/swclabs/swipe-server/internal/broker/handler"
	"github.com/swclabs/swipe-server/pkg/tools/worker"
)

var _AccountManagement = handler.NewAccountManagement()

func AccountManagement(eng *worker.Engine) {
	eng.Queue(_AccountManagement.HandleOAuth2SaveUser)
	eng.Queue(_AccountManagement.HandleSignUp)
	eng.Queue(_AccountManagement.HandleUpdateUserInfo)
}
