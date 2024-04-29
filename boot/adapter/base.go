package adapter

import (
	"swclabs/swipecore/internal/http"
)

const TypeBase = "BaseAdapter"

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
		http.ProductsModule,
		http.PostsModule,
	)
	return adapter.server.Run(addr)
}
