// product management adapter, base on base.adapt.go

package apis

import (
	"swclabs/swix/internal/apis/router"
	"swclabs/swix/internal/apis/server"
	"swclabs/swix/internal/types"
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
