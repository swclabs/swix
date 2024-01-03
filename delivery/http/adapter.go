package http

import "swclabs/swiftcart/delivery/app"

type Adapter struct {
	server *app.App
}

func NewGinAdapter() *Adapter {
	return &Adapter{
		server: app.New(),
	}
}

func (adapter *Adapter) ListenOn(port string) error {
	return adapter.server.Run(port)
}
