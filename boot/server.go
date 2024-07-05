package boot

import (
	"fmt"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/types"
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

// Connect to module via adapter
//
//	func main() {
//		var (
//			commonService    = common.New()
//			commonController = controller.NewCommon(commonService)
//			commonRouter     = router.NewCommon(commonController)
//			httpServer = http.NewServer([]router.IRouter{
//				commonRouter,
//				router.NewDocs(),
//			})
//			adapt = http.NewBaseAdapter(httpServer)
//			server = boot.NewServer()
//		)
//
//		log.Fatal(server.Connect(adapt))
//	}
func (server *_Server) Connect(adapter types.IAdapter) error {
	return adapter.Run(server.address)
}
