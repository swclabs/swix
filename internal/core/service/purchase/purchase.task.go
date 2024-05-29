package purchase

import "swclabs/swipecore/internal/core/domain"

type IPurchase interface {
	DelayAddToCart(cart domain.CartInfo) error
	DelayInsertOrders(order domain.Orders) error
}

var _ IPurchase = (*Task)(nil)

type Task struct{}

// DelayAddToCart implements IPurchase.
func (p *Task) DelayAddToCart(cart domain.CartInfo) error {
	panic("unimplemented")
}

// DelayInsertOrders implements IPurchase.
func (p *Task) DelayInsertOrders(order domain.Orders) error {
	panic("unimplemented")
}
