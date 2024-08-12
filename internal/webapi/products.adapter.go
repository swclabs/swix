// product management adapter, base on base.adapt.go

package webapi

import (
	"swclabs/swix/internal/types"
	"swclabs/swix/internal/webapi/router"
	"swclabs/swix/internal/webapi/server"
)

type _ProductsAdapter struct {
	server server.IServer
}

var _ types.IAdapter = (*_ProductsAdapter)(nil)

// NewProductsAdapter creates a new adapter wrapping the given server
func NewProductsAdapter(
	server server.IServer,
	router router.IProducts,
) types.IAdapter {
	product := &_ProductsAdapter{
		server: server,
	}
	product.server.Connect(router)
	return product
}

func (product *_ProductsAdapter) Run(addr string) error {
	return product.server.Run(addr)
}
