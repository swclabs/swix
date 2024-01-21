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
	adapter := &_Adapter{
		server: http.New(),
	}
	// initialize cron job
	adapter._StartCommonJob()
	return adapter
}

func (adapter *_Adapter) Run(addr string) error {
	adapter.server.InitAccountManagement()
	return adapter.server.Run(addr)
}

type _AccountManagementAdapter struct {
	server http.IServer
}

func NewAccountManagementAdapter() IAdapter {
	account := &_AccountManagementAdapter{
		server: http.New(),
	}
	// initialize cron job
	account._StartAccountManagementJob()
	return account
}

func (account *_AccountManagementAdapter) Run(addr string) error {
	account.server.InitAccountManagement()
	return account.server.Run(addr)
}
