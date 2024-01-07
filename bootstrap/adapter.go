package bootstrap

import (
	"swclabs/swiftcart/app"
)

type Adapter struct {
	server *app.Server
}

func NewGinAdapter() *Adapter {
	return &Adapter{
		server: app.New(),
	}
}

func (adapter *Adapter) ListenOn(port string) error {
	return adapter.server.Run(port)
}
