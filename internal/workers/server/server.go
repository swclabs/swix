// Package server define worker consume
package server

import (
	"fmt"
	"strconv"
	"swclabs/swix/internal/workers/queue"
	"swclabs/swix/pkg/lib/worker"
)

// IWorker interface for Worker objects
type IWorker interface {
	Run(concurrency string) error
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
func (msg *Worker) Run(concurrency string) error {
	msg.mux.Serve(msg.e)
	_concurrency, err := strconv.Atoi(concurrency)
	if err != nil {
		return fmt.Errorf("invalid concurrency value: %v", err)
	}
	return msg.e.Run(_concurrency)
}
