// product management adapter, base on base.adapt.go

package wapi

import (
	"fmt"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/internal/wapi/router"
)

type _ProductsAdapter struct {
	server IServer
}

var _ types.IAdapter = (*_ProductsAdapter)(nil)

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

// StartWorker implements types.IAdapter.
func (product *_ProductsAdapter) StartWorker(concurrency int) error {
	return fmt.Errorf("services unavailable")
}
