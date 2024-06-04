// account management adapter

package adapter

import (
	"swclabs/swipecore/internal/http"
	"swclabs/swipecore/internal/http/router"
)

type _AccountManagement struct {
	server http.IServer
}

func NewAccountManagements(
	server http.IServer,
	router *router.AccountManagement,
) IAdapter {
	product := &_AccountManagement{
		server: server,
	}
	product.server.Connect(router)
	return product
}

func (accountmanagement *_AccountManagement) Run(addr string) error {
	return accountmanagement.server.Run(addr)
}
