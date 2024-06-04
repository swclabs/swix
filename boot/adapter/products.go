// product management adapter, base on base.adapt.go

package adapter

import (
	"swclabs/swipecore/internal/http"
	"swclabs/swipecore/internal/http/router"
)

type _Products struct {
	server http.IServer
}

func NewProducts(server http.IServer, router *router.Products) IAdapter {
	product := &_Products{
		server: server,
	}
	product.server.Connect(router)
	return product
}

func (product *_Products) Run(addr string) error {
	return product.server.Run(addr)
}
