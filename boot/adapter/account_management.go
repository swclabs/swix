// account management adapter, base on base.adapt.go

package adapter

import (
	"swclabs/swipecore/internal/http"
)

const TypeAccountManagement = "AccountManagementAdapter"

type _AccountManagementAdapter struct {
	server http.IServer
}

func _NewAccountManagement() IAdapter {
	account := &_AccountManagementAdapter{
		server: http.New(),
	}
	// initialize cron job
	// account._StartAccountManagementJob()
	return account
}

func (account *_AccountManagementAdapter) Run(addr string) error {
	account.server.Bootstrap(http.AccountManagementModule)
	return account.server.Run(addr)
}
