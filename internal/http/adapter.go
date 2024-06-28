package http

import (
	"swclabs/swipecore/internal/http/router"
)

type _Adapter struct {
	server IServer
}

var _ IAdapter = (*_Adapter)(nil)

// NewAdapter Example
/*
package main

import (

	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/http"

	"go.uber.org/fx"

)

	func main() {
		app := fx.New(
			boot.FxRestModule,
			fx.Provide(
				adapter.NewAdapter,
				boot.NewServerEnv,
			),
			fx.Invoke(boot.StartServer),
		)
		app.Run()
	}
*/
func NewAdapter(server IServer,
	products router.IProducts,
	accountManagement router.IAccountManagement,
	posts router.IPosts,
	purchase router.IPurchase,
) IAdapter {
	adapter := &_Adapter{
		server: server,
	}

	adapter.server.Connect(products)
	adapter.server.Connect(accountManagement)
	adapter.server.Connect(posts)
	adapter.server.Connect(purchase)

	return adapter
}

// Run : run all services one server
func (adapter *_Adapter) Run(addr string) error {
	return adapter.server.Run(addr)
}

func (adapter *_Adapter) Routers() []string {
	return adapter.server.Routes()
}
