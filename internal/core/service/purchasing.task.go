package service

import "swclabs/swipecore/internal/core/domain"

type IPurchasing interface {
	DelayAddToCart(cart domain.CartInfo) error
	DelayInsertOrders(order domain.Orders) error
}

var _ IPurchasing = (*PurchasingTask)(nil)

type PurchasingTask struct{}

// DelayAddToCart implements IPurchasing.
func (p *PurchasingTask) DelayAddToCart(cart domain.CartInfo) error {
	panic("unimplemented")
}

// DelayInsertOrders implements IPurchasing.
func (p *PurchasingTask) DelayInsertOrders(order domain.Orders) error {
	panic("unimplemented")
}
