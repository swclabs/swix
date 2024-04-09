package adapter

import (
	"swclabs/swipe-api/internal/http"
)

const TypeBase = "BaseAdapter"

// IAdapter interface, used to connect to server instance
type IAdapter interface {
	Run(addr string) error
}

func New(types string) IAdapter {
	switch types {
	case TypeBase:
		return _NewAdapter()
	case TypeAccountManagement:
		return _NewAccountManagement()
	case TypeProductManagement:
		return _NewProductManagement()
	case TypeProducts:
		return _NewProducts()
	}
	return _NewAdapter()
}

type _Adapter struct {
	server http.IServer
}

func _NewAdapter() IAdapter {
	adapter := &_Adapter{
		server: http.New(),
	}
	// initialize cron job
	// adapter._StartCommonJob()
	return adapter
}

// Run : run all services one server
func (adapter *_Adapter) Run(addr string) error {
	adapter.server.Bootstrap(
		http.AccountManagementModule,
		http.ProductManagementModule,
		http.ProductsModule,
	)
	return adapter.server.Run(addr)
}
