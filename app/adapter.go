package app

import (
	"swclabs/swiftcart/delivery/http"
)

type Adapter struct {
	server *http.Server
}

func NewGinAdapter() *Adapter {
	return &Adapter{
		server: http.New(),
	}
}

func (adapter *Adapter) ListenOn(port string) error {
	return adapter.server.Run(port)
}
