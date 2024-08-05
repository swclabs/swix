// Package workers define worker consume
package workers

import (
	"swclabs/swix/internal/workers/router"
	"swclabs/swix/pkg/lib/worker"
)

// Writer struct define the Writer object
type Writer struct {
	engine *worker.Engine
}

// NewWriter creates a new Writer object
func NewWriter(
	engine *worker.Engine,
	base *router.Base,
	manager *router.Manager,
) *Writer {

	writer := &Writer{
		engine: engine,
	}

	base.Register(writer.engine)
	manager.Register(writer.engine)

	return writer
}

// Run runs the worker engine
func (msg *Writer) Run(concurrency int) error {
	return msg.engine.Run(concurrency)
}
