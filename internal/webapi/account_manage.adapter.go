// Package webapi account management
package webapi

import (
	"fmt"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/webapi/router"
)

type _AccountManagementAdapter struct {
	server IServer
}

var _ types.IAdapter = (*_AccountManagementAdapter)(nil)

// NewAccountManagementsAdapter returns a new adapter wrapping around the given server
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
func (accountManagement *_AccountManagementAdapter) StartWorker(_ int) error {
	return fmt.Errorf("services unavailable")
}
