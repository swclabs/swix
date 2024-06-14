/**
 * boot folder representing the delevery layer in clean architecture
 * you can use this folder to define any configuration settings or
 * operation, start-up applications

 * Package boot implement api server for swipe application

 * You can use _Server to connect to specific service adapters.
 * use fx Framework (uber-go/fx) to create your own adapters
 * with dependency injection pattern.

 * For each FxModule from the layers in the project, you can
 * add them to the Fx.New method to provide the necessary
 * constructors for a smooth application startup.

 * You can find more information about the fx.go in each directory
 * representing the layers of the project

 * Then you can use _Server to connect adapter through Connect methods

 * See the example below.

Example:

package main

import (
	"log"
	"swclabs/swipecore/boot"
	"swclabs/swipecore/boot/adapter"

	"go.uber.org/fx"
)

func StartServer(server boot.IServer, adapter adapter.IAdapter) {
	go func() {
		log.Fatal(server.Connect(adapter))
	}()
}

func main() {
	app := fx.New(
		boot.FxRestModule,
		fx.Provide(
			adapter.NewAdapter,
			boot.NewServer,
		),
		fx.Invoke(StartServer),
	)
	app.Run()
}


*/

package boot

import (
	"context"
	"fmt"
	"log"
	"swclabs/swipecore/boot/adapter"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/db"

	"go.uber.org/fx"
)

type IServer interface {
	// Connect to adapter of other module
	Connect(adapter adapter.IAdapter) error
}

// struct server in project
//
// host:port - 127.0.0.1:8000
type _Server struct {
	address string //
}

// NewServer creates a new server instance
// Use for fx Framework and more
func NewServer(env config.Env) IServer {
	return &_Server{
		address: fmt.Sprintf("%s:%s", env.Host, env.Port),
	}
}

// Connect to module via adapter
//
//	func main() {
//		var (
//			env = config.LoadEnv()
//			commonService = common.New(worker.NewClient(env))
//			commonController = controller.NewCommon(commonService)
//			commonRouter = router.NewCommon(commonController)
//			httpServer = http.NewServer(commonRouter, router.NewDocs())
//			adapt = adapter.NewBaseAdapter(httpServer)
//			server = boot.NewServer(env)
//		)
//
//		log.Fatal(server.Connect(adapt))
//	}
func (server *_Server) Connect(adapter adapter.IAdapter) error {
	return adapter.Run(server.address)
}

// StartServer used to start a server, through to fx.Invoke() method
//
//	app := fx.New(
//		boot.FxRestModule,
//		fx.Provide(
//			adapter.NewAdapter,
//			boot.NewServer,
//		),
//		fx.Invoke(boot.StartServer), // <-- run here
//	)
//	app.Run()
func StartServer(lc fx.Lifecycle, server IServer, adapter adapter.IAdapter) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err := db.MigrateUp(); err != nil {
				return err
			}
			go func() {
				log.Fatal(server.Connect(adapter))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Server stopping")
			return nil
		},
	})
}
