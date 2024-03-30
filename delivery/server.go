package delivery

import "swclabs/swipe-api/delivery/adapter"

type IServer interface {
	Connect(adapter adapter.IAdapter) error
}

type _Server struct {
	address string
}

func NewServer(addr string) IServer {
	return &_Server{
		address: addr,
	}
}

func (server *_Server) Connect(adapter adapter.IAdapter) error {
	return adapter.Run(server.address)
}
