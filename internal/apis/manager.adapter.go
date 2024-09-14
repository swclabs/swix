// Package apis manager
package apis

import (
	"swclabs/swix/internal/apis/router"
	"swclabs/swix/internal/apis/server"
	"swclabs/swix/internal/types"
)

type _ManagerAdapter struct {
	server server.IServer
}

var _ types.IAdapter = (*_ManagerAdapter)(nil)

// NewManagerAdapter returns a new adapter wrapping around the given server
func NewManagerAdapter(
	server server.IServer,
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
