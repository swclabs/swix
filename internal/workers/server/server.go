// Package server define worker consume
package server

import (
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/workers/queue"
	"swclabs/swix/pkg/lib/worker"
)

// IWorker interface for Worker objects
type IWorker interface {
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

// New creates a new Writer object
func New(mux IMux) IWorker {
	writer := &Worker{
		e:   worker.New(queue.New()),
		mux: mux,
	}
	return writer
}

// Run runs the worker engine
func (msg *Worker) Run() error {
	msg.mux.Serve(msg.e)
	return msg.e.Run(config.NumberOfWorker)
}
