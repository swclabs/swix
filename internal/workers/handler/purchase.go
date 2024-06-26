package handler

import (
	"swclabs/swipecore/pkg/lib/worker"
)

type IPurchase interface {
	HandleAddToCart() (string, worker.HandleFunc)
	HandleInsertOrders() (string, worker.HandleFunc)
}

var _ IPurchase = (*Purchase)(nil)

type Purchase struct {
}

// HandleAddToCart implements IPurchase.
func (p *Purchase) HandleAddToCart() (string, worker.HandleFunc) {
	panic("unimplemented")
}

// HandleInsertOrders implements IPurchase.
func (p *Purchase) HandleInsertOrders() (string, worker.HandleFunc) {
	panic("unimplemented")
}
