// product management adapter, base on base.adapt.go

package delivery

import "github.com/swclabs/swipe-server/internal/http"

type _ProductManagementAdapter struct {
	server http.IServer
}

func NewProductManagementAdapter() IAdapter {
	product := &_ProductManagementAdapter{
		server: http.New(),
	}
	return product
}

func (product *_ProductManagementAdapter) Run(addr string) error {
	product.server.Bootstrap(http.ProductManagementModule)
	return product.server.Run(addr)
}
