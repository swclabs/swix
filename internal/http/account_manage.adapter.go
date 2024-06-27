// account management adapter

package http

import (
	"swclabs/swipecore/internal/http/router"
)

type _AccountManagementAdapter struct {
	server IServer
}

var _ IAdapter = (*_AccountManagementAdapter)(nil)

func NewAccountManagementsAdapter(
	server IServer,
	router router.IAccountManagement,
) IAdapter {
	product := &_AccountManagementAdapter{
		server: server,
	}
	product.server.Connect(router)
	return product
}

func (accountManagement *_AccountManagementAdapter) Run(addr string) error {
	return accountManagement.server.Run(addr)
}

func (accountManagement *_AccountManagementAdapter) Routers() []string {
	return accountManagement.server.Routes()
}
