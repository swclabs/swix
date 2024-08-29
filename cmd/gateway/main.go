package main

import (
	"context"
	"log"
	"swclabs/swix/boot"
	"swclabs/swix/internal/cluster"
	"swclabs/swix/internal/cluster/handlers"
	"swclabs/swix/internal/cluster/nodes"
	"swclabs/swix/internal/types"

	"go.uber.org/fx"
)

func start(lc fx.Lifecycle, srv boot.IServer, adapter types.IAdapter) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				log.Fatal(srv.Connect(adapter))
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			handlers.NewGreeter,
			nodes.NewGreeter,
			cluster.New,
			cluster.NewGateway,
			cluster.NewAdapter,
			boot.NewServer,
		),
		fx.Invoke(start),
	)
	app.Run()
}
