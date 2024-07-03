package workers

import (
	"fmt"
	"swclabs/swipecore/internal/types"
)

type Adapter struct {
	engine *Writer
}

func NewAdapter(writer *Writer) types.IAdapter {
	return &Adapter{
		engine: writer,
	}
}

// Run implements types.IAdapter.
func (a *Adapter) Run(addr string) error {
	return fmt.Errorf("service unavailable")
}

// StartWorker implements types.IAdapter.
func (a *Adapter) StartWorker(concurrency int) error {
	return a.engine.Run(concurrency)
}
