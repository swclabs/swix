// Package server define worker consume
package server

import (
	"swclabs/swipex/internal/config"
	"swclabs/swipex/internal/workers/queue"
	"swclabs/swipex/pkg/lib/worker"
)

// IWorker interface for Worker objects
type IWorker interface {
	Register(taskName string, fn worker.HandleFunc)
	Run() error
}

type IRouter interface {
	Register(eng worker.IEngine)
}

// Worker struct define the Worker object
type Worker struct {
	e   worker.IEngine
	mux IMux
}

// New creates a new worker object
func New(mux IMux) IWorker {
	worker := &Worker{
		e:   worker.New(queue.New()),
		mux: mux,
	}
	return worker
}

// Register register the task
func (worker *Worker) Register(taskName string, fn worker.HandleFunc) {
	worker.e.Register(taskName, fn)
}

// Run runs the worker engine
func (worker *Worker) Run() error {
	worker.mux.Serve(worker.e)
	return worker.e.Run(config.NumberOfWorker)
}
