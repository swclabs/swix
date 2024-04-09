// product management adapter, base on base.adapt.go

package adapter

import "swclabs/swipe-api/internal/http"

const TypeProducts = "Products"

type _Products struct {
	server http.IServer
}

func _NewProducts() IAdapter {
	product := &_ProductManagementAdapter{
		server: http.New(),
	}
	return product
}

func (product *_Products) Run(addr string) error {
	product.server.Bootstrap(http.ProductsModule)
	return product.server.Run(addr)
}
