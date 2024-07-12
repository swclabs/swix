// account management adapter

package wapi

import (
	"fmt"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/wapi/router"
)

type _AccountManagementAdapter struct {
	server IServer
}

var _ types.IAdapter = (*_AccountManagementAdapter)(nil)

func NewAccountManagementsAdapter(
	server IServer,
	router router.IAccountManagement,
) types.IAdapter {
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

// StartWorker implements types.IAdapter.
func (accountManagement *_AccountManagementAdapter) StartWorker(concurrency int) error {
	return fmt.Errorf("services unavailable")
}
