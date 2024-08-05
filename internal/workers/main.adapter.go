// Package workers define worker consume
package workers

import (
	"fmt"
	"strconv"
	"swclabs/swix/internal/types"
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
func (a *Adapter) Run(concurrency string) error {
	_concurrency, err := strconv.Atoi(concurrency)
	if err != nil {
		return fmt.Errorf("invalid concurrency value: %v", err)
	}
	return a.engine.Run(_concurrency)
}
