package purchase

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/carts"
	"swclabs/swipecore/internal/core/repository/orders"
)

type Purchase struct {
	order orders.IOrdersRepository
	cart  carts.ICartRepository
}

func New(
	order *orders.Orders,
	cart *carts.Carts,
) IPurchaseService {
	return &Purchase{
		cart:  cart,
		order: order,
	}
}

// AddToCart implements domain.IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart domain.CartSchema) {
	panic("unimplemented")
}

// GetCart implements domain.IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, limit int) ([]domain.Carts, error) {
	panic("unimplemented")
}

// GetOrders implements domain.IPurchaseService.
func (p *Purchase) GetOrders(ctx context.Context, limit int) ([]domain.Orders, error) {
	panic("unimplemented")
}

// InsertOrders implements domain.IPurchaseService.
func (p *Purchase) InsertOrders(ctx context.Context, order domain.Orders) error {
	panic("unimplemented")
}
