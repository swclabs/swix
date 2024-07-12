// Package workers define worker consume
package workers

import (
	"fmt"
	"swclabs/swipecore/internal/types"
)

// Adapter struct define the Adapter object
type Adapter struct {
	engine *Writer
}

// NewAdapter creates a new Adapter object
func NewAdapter(writer *Writer) types.IAdapter {
	return &Adapter{
		engine: writer,
	}
}

// Run implements types.IAdapter.
func (a *Adapter) Run(_ string) error {
	return fmt.Errorf("service unavailable")
}

// StartWorker implements types.IAdapter.
func (a *Adapter) StartWorker(concurrency int) error {
	return a.engine.Run(concurrency)
}
