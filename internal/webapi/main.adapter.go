package webapi

import (
	"swclabs/swix/internal/types"
	"swclabs/swix/internal/webapi/router"
	"swclabs/swix/internal/webapi/server"
)

type _Adapter struct {
	server server.IServer
}

var _ types.IAdapter = (*_Adapter)(nil)

// NewAdapter returns a new adapter wrapping around the given server
func NewAdapter(
	server server.IServer,
	products router.IProducts,
	manager router.IManager,
	article router.IArticle,
	purchase router.IPurchase,
	classify router.IClassify,
	paydeli router.IPaydeliver,
) types.IAdapter {
	adapter := &_Adapter{
		server: server,
	}

	adapter.server.Connect(products)
	adapter.server.Connect(manager)
	adapter.server.Connect(article)
	adapter.server.Connect(purchase)
	adapter.server.Connect(classify)
	adapter.server.Connect(paydeli)
	return adapter
}

// Run : run all services one server
func (adapter *_Adapter) Run(addr string) error {
	return adapter.server.Run(addr)
}
