// product management adapter, base on base.adapt.go

package http

import (
	"swclabs/swipecore/internal/http/router"
)

type _ProductsAdapter struct {
	server IServer
}

var _ IAdapter = (*_ProductsAdapter)(nil)

func NewProductsAdapter(server IServer, router router.IProducts) IAdapter {
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
