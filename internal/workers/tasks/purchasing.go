package tasks

import "swclabs/swipe-api/internal/core/domain"

type IPurchasing interface {
	DelayAddToCart(cart domain.CartInfo) error
	DelayInsertOrders(order domain.Orders) error
}

var _ IPurchasing = (*Purchasing)(nil)

type Purchasing struct{}

// DelayAddToCart implements IPurchasing.
func (p *Purchasing) DelayAddToCart(cart domain.CartInfo) error {
	panic("unimplemented")
}

// DelayInsertOrders implements IPurchasing.
func (p *Purchasing) DelayInsertOrders(order domain.Orders) error {
	panic("unimplemented")
}
