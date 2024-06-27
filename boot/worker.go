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
	"swclabs/swipecore/internal/workers"

	"go.uber.org/fx"
)

// IWorker interface of type Worker
type IWorker interface {
	Run(concurrency int) error
}

type _Worker struct {
	// engine of worker, worker consume is writer of database
	engine *workers.Writer
}

// NewWorker create new worker consume
func NewWorker(writer *workers.Writer) IWorker {
	return &_Worker{
		engine: writer,
	}
}

// Run worker with concurrency is number of worker
func (w *_Worker) Run(concurrency int) error {
	return w.engine.Run(concurrency)
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
func StartWorker(lc fx.Lifecycle, worker IWorker) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Fatal(worker.Run(10))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("[Swipe]   OnStop                server worker stopping")
			return nil
		},
	})
}
