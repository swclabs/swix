// Package workers define worker consume
package workers

import (
	"swclabs/swipecore/internal/workers/router"
	"swclabs/swipecore/pkg/lib/worker"
)

// Writer struct define the Writer object
type Writer struct {
	engine *worker.Engine
}

// NewWriter creates a new Writer object
func NewWriter(
	engine *worker.Engine,
	common *router.Common,
	accountManagement *router.AccountManagements,
) *Writer {

	writer := &Writer{
		engine: engine,
	}

	common.Register(writer.engine)
	accountManagement.Register(writer.engine)

	return writer
}

// Run runs the worker engine
func (msg *Writer) Run(concurrency int) error {
	return msg.engine.Run(concurrency)
}
