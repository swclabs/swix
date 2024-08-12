// Package server define worker consume
package server

import (
	"swclabs/swix/pkg/lib/worker"
)

// IWorker interface for Worker objects
type IWorker interface {
	RegisterQueue(hfn func() (taskName string, fn worker.HandleFunc))
	Run(concurrency int) error
}

// Worker struct define the Worker object
type Worker struct {
	e worker.IEngine
}

// New creates a new Writer object
func New(engine worker.IEngine) IWorker {
	writer := &Worker{
		e: engine,
	}
	return writer
}

// RegisterQueue registers a new task name and handler function
func (msg *Worker) RegisterQueue(hfn func() (taskName string, fn worker.HandleFunc)) {
	msg.e.RegisterQueue(hfn)
}

// Run runs the worker engine
func (msg *Worker) Run(concurrency int) error {
	return msg.e.Run(concurrency)
}
