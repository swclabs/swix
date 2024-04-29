package handler

import (
	"swclabs/swipecore/internal/workers/tasks"
	"swclabs/swipecore/pkg/tools/worker"
)

type IPurchasing interface {
	HandleAddToCart() (string, worker.HandleFunc)
	HandleInsertOrders() (string, worker.HandleFunc)
}

var _ IPurchasing = (*Purchasing)(nil)

type Purchasing struct {
	tasks.Purchasing
}

// HandleAddToCart implements IPurchasing.
func (p *Purchasing) HandleAddToCart() (string, worker.HandleFunc) {
	panic("unimplemented")
}

// HandleInsertOrders implements IPurchasing.
func (p *Purchasing) HandleInsertOrders() (string, worker.HandleFunc) {
	panic("unimplemented")
}
