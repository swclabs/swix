// Package handler implements handler of worker
package handler

import (
	"swclabs/swipecore/pkg/lib/worker"
)

// IPurchase is an interface for Purchase.
type IPurchase interface {
	HandleAddToCart() (string, worker.HandleFunc)
	HandleInsertOrders() (string, worker.HandleFunc)
}

var _ IPurchase = (*Purchase)(nil)

// Purchase is a struct for Purchase.
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
