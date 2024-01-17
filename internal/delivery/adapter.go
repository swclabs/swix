package delivery

import (
	"github.com/swclabs/swipe-api/internal/http"
)

type IAdapter interface {
	Run(addr string) error
}

type _Adapter struct {
	server http.IServer
}

func NewAdapter() IAdapter {
	// initialize cron job
	_StartCommonJob()
	return &_Adapter{
		server: http.New(),
	}
}

func (adapter *_Adapter) Run(addr string) error {
	adapter.server.InitAccountManagement()
	return adapter.server.Run(addr)
}

type _AccountManagementAdapter struct {
	server http.IServer
}

func NewAccountManagementAdapter() IAdapter {
	// initialize cron job
	_StartCommonJob()
	return &_AccountManagementAdapter{
		server: http.New(),
	}
}

func (account *_AccountManagementAdapter) Run(addr string) error {
	account.server.InitAccountManagement()
	return account.server.Run(addr)
}
