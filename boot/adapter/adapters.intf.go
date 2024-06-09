package adapter

import "swclabs/swipecore/internal/http"

// IAdapter interface, used to connect to server instance
type IAdapter interface {
	Run(addr string) error
}

func NewBaseAdapter(server http.IServer) IAdapter {
	adapter := &_Adapter{
		server: server,
	}
	return adapter
}
