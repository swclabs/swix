package adapter

import (
	"swclabs/swipecore/internal/http"
	"swclabs/swipecore/internal/http/router"
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
	adapter.server.Connect(router.New(router.TypeAccountManagement))
	adapter.server.Connect(router.New(router.TypeProducts))
	adapter.server.Connect(router.New(router.TypePosts))
	return adapter.server.Run(addr)
}
