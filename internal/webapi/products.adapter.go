// product management adapter, base on base.adapt.go

package webapi

import (
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/webapi/router"
)

type _ProductsAdapter struct {
	server IServer
}

var _ types.IAdapter = (*_ProductsAdapter)(nil)

// NewProductsAdapter creates a new adapter wrapping the given server
func NewProductsAdapter(
	server IServer,
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

func (product *_ProductsAdapter) Routers() []string {
	return product.server.Routes()
}
