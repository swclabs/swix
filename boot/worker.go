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
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/workers"
	"swclabs/swipecore/pkg/lib/worker"
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
// ENV:
//
// REDIS_HOST=localhost
// REDIS_PORT=6379
// REDIS_PASSWORD=password
func NewWorker() IWorker {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
	return &_Worker{
		engine: workers.New(),
	}
}

// Run worker with concurrency is number of worker
func (w *_Worker) Run(concurrency int) error {
	return w.engine.Run(concurrency)
}
