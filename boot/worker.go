/*
Package boot implement worker consume for swipe application

Example:

	package main

	import (
		"log"

		"swclabs/swipecore/boot"
	)

	func main() {
		w := boot.NewWorker()
		if err := w.Run(10); err != nil {
			log.Fatal(err)
		}
	}
*/

package boot

import (
	"context"
	"fmt"
	"log"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/types"

	"go.uber.org/fx"
)

type IWorker interface {
	IBase
}

type _Worker struct {
	concurrency int
}

// NewWorker create new worker consume
func NewWorker(env config.Env) IWorker {
	return &_Worker{
		concurrency: 10,
	}
}

// Run worker with concurrency is number of worker
func (w *_Worker) Connect(adapter types.IAdapter) error {
	return adapter.StartWorker(w.concurrency)
}

// StartWorker used to start a worker consume server,
// through to fx.Invoke() method
//
// app := fx.New(
//
//	boot.FxWorkerModule,
//	fx.Provide(
//		boot.NewWorker,
//	),
//	fx.Invoke(boot.StartWorker),
//
// )
// app.Run()
func StartWorker(lc fx.Lifecycle, worker IWorker, adapter types.IAdapter) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Fatal(worker.Connect(adapter))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("[Swipe]   OnStop                server worker stopping")
			return nil
		},
	})
}
