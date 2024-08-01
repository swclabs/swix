// Package webapi manager
package webapi

import (
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/webapi/router"
)

type _ManagerAdapter struct {
	server IServer
}

var _ types.IAdapter = (*_ManagerAdapter)(nil)

// NewManagerAdapter returns a new adapter wrapping around the given server
func NewManagerAdapter(
	server IServer,
	router router.IManager,
) types.IAdapter {
	product := &_ManagerAdapter{
		server: server,
	}
	product.server.Connect(router)
	return product
}

func (manager *_ManagerAdapter) Run(addr string) error {
	return manager.server.Run(addr)
}
