package boot

import (
	"fmt"
	"swclabs/swix/internal/config"
)

// struct server in project
//
// host:port - 127.0.0.1:8000
type _Server struct {
	address string //
}

// NewServer creates a new server instance
// Use for fx Framework and more
func NewServer() IServer {
	return &_Server{
		address: fmt.Sprintf("%s:%s", config.Host, config.Port),
	}
}

// NewServerWithAddress creates a new server instance with address
func NewServerWithAddress(address string) IServer {
	return &_Server{
		address: address,
	}
}

// Bootstrap to module via adapter
//
//	func main() {
//		var (
//			baseService    = base.New()
//			baseController = controller.New(baseService)
//			baseRouter     = router.New(baseController)
//			mux    = server.NewServeMux()
//			server = server.New(mux)
//		)
//		mux.Handle(baseRouter)
//		log.Fatal(server.Run("localhost:8000"))
//	}
func (server *_Server) Bootstrap(core ICore) error {
	return core.Run(server.address)
}
