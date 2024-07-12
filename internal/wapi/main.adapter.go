package wapi

import (
	"fmt"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/wapi/router"
)

type _Adapter struct {
	server IServer
}

var _ types.IAdapter = (*_Adapter)(nil)

// NewAdapter returns a new adapter wrapping around the given server
func NewAdapter(
	server IServer,
	products router.IProducts,
	accountManagement router.IAccountManagement,
	posts router.IPosts,
	purchase router.IPurchase,
) types.IAdapter {
	adapter := &_Adapter{
		server: server,
	}

	adapter.server.Connect(products)
	adapter.server.Connect(accountManagement)
	adapter.server.Connect(posts)
	adapter.server.Connect(purchase)

	return adapter
}

// Run : run all services one server
func (adapter *_Adapter) Run(addr string) error {
	return adapter.server.Run(addr)
}

func (adapter *_Adapter) Routers() []string {
	return adapter.server.Routes()
}

// StartWorker implements types.IAdapter.
func (adapter *_Adapter) StartWorker(_ int) error {
	return fmt.Errorf("services unavailable")
}
