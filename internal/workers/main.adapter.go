// Package workers define worker consume
package workers

import (
	"fmt"
	"strconv"
	"swclabs/swix/internal/types"
	"swclabs/swix/internal/workers/router"
	"swclabs/swix/internal/workers/server"
)

// Adapter struct define the Adapter object
type Adapter struct {
	engine server.IWorker
}

// NewAdapter creates a new Adapter object
func NewAdapter(
	writer server.IWorker,
	base router.IBase,
	manager router.IManager,
	purchase router.IPurchase,
) types.IAdapter {

	base.Register(writer)
	manager.Register(writer)
	purchase.Register(writer)

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
