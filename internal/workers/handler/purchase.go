package handler

import (
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/pkg/lib/worker"
)

type IPurchase interface {
	HandleAddToCart() (string, worker.HandleFunc)
	HandleInsertOrders() (string, worker.HandleFunc)
}

var _ IPurchase = (*Purchase)(nil)

type Purchase struct {
	service.PurchasingTask
}

// HandleAddToCart implements IPurchase.
func (p *Purchase) HandleAddToCart() (string, worker.HandleFunc) {
	panic("unimplemented")
}

// HandleInsertOrders implements IPurchase.
func (p *Purchase) HandleInsertOrders() (string, worker.HandleFunc) {
	panic("unimplemented")
}
