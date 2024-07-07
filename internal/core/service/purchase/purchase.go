package purchase

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/carts"
	"swclabs/swipecore/internal/core/repository/orders"
)

type Purchase struct {
	Order orders.IOrdersRepository
	Cart  carts.ICartRepository
}

func New(
	order orders.IOrdersRepository,
	cart carts.ICartRepository,
) IPurchaseService {
	return &Purchase{
		Cart:  cart,
		Order: order,
	}
}

// DeleteItemFromCart implements domain.IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, userId int64, inventoryId int64) error {
	return p.Cart.RemoveItem(ctx, userId, inventoryId)
}

// AddToCart implements domain.IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart domain.CartInsert) error {
	return p.Cart.Insert(ctx, cart.UserId, cart.InventoryId, cart.Quantity)
}

// GetCart implements domain.IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userId int64, limit int) (*domain.CartSlices, error) {
	return p.Cart.GetCartByUserID(ctx, userId, limit)
}

// GetOrders implements domain.IPurchaseService.
func (p *Purchase) GetOrders(ctx context.Context, limit int) ([]domain.Orders, error) {
	panic("unimplemented")
}

// InsertOrders implements domain.IPurchaseService.
func (p *Purchase) InsertOrders(ctx context.Context, userId int64, inventoryId ...int64) (string, error) {
	panic("unimplemented")
}
