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

 * See the example below.

Example:

package main

import (
	"log"
	"swclabs/swipecore/boot"
	"swclabs/swipecore/internal/webapipipipi"
	"swclabs/swipecore/internal/types"

	"go.uber.org/fx"
)

func StartServer(server boot.IServer, adapter types.IAdapter) {
	go func() {
		log.Fatal(server.Connect(adapter))
	}()
}

func main() {
	app := fx.New(
		boot.FxModule(),
		fx.Provide(
			webapi.NewAdapter,
			boot.NewServer,
		),
		fx.Invoke(boot.Main),
	)
	app.Run()
}
*/

package boot

import (
	"context"
	"log"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/lib/logger"

	"go.uber.org/fx"

	_ "swclabs/swipecore/boot/init" // init package deps, like docs, migration
)

// IServer connect and run via adapter (webapi, worker, rpc)
type IServer interface {
	Connect(adapter types.IAdapter) error
}

// NewApp used to create Fx Application
func NewApp(serverConstructor func() IServer, adapterConstructors ...interface{}) *fx.App {
	return fx.New(
		fxModule(),
		fx.Provide(adapterConstructors...),
		fx.Provide(serverConstructor),
		fx.Invoke(Main),
	)
}

// Main used to start a server, through to fx.Invoke() method
//
//	boot.PrepareFor(boot.RestAPI | boot.ProdMode)
//	app := fx.New(
//		boot.FxModule(),
//		fx.Provide(
//			webapi.NewAdapter,
//			boot.NewServer,
//		),
//		fx.Invoke(boot.Main), // <-- run here
//	)
//	app.Run()
func Main(lc fx.Lifecycle, server IServer, adapter types.IAdapter) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			logger.Info("Server starting")
			if err := db.MigrateUp(); err != nil {
				return err
			}
			go func() {
				log.Fatal(server.Connect(adapter))
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info("Server stopping")
			return nil
		},
	})
}
