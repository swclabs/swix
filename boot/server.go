/*
Package boot implement api server for swipe application

Example:

	package main

	import (
		"fmt"
		"log"

		"swclabs/swipecore/boot"
		"swclabs/swipecore/boot/adapter"
		"swclabs/swipecore/internal/config"
	)

	func main() {
		addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
		server := boot.NewServer(addr)
		adapt := adapter.New(adapter.TypeBase)

		if err := server.Connect(adapt); err != nil {
			log.Fatal(err)
		}
	}
*/

package boot

import (
	"context"
	"fmt"
	"swclabs/swipecore/boot/adapter"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/lib/worker"

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

func NewServer(env config.Env) IServer {
	return &_Server{
		address: fmt.Sprintf("%s:%s", env.Host, env.Port),
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

func StartServer(lc fx.Lifecycle, server IServer, env config.Env,
	adapter adapter.IAdapter,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			worker.SetBroker(env.RedisHost, env.RedisPort, env.RedisPassword)
			go server.Connect(adapter)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Server stopping")
			return nil
		},
	})
}
