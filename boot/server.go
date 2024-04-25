package boot

import "swclabs/swipe-api/boot/adapter"

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

// Connect to module via adapter
//
// Example:
//
//	server := boot.NewServer("localhost:8000")
//	adapter := adapter.NewAdapter()
//	server.Connect(adapter)
func (server *_Server) Connect(adapter adapter.IAdapter) error {
	return adapter.Run(server.address)
}
