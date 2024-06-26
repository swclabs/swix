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
	order orders.IOrdersRepository,
	cart carts.ICartRepository,
) IPurchaseService {
	return &Purchase{
		cart:  cart,
		order: order,
	}
}

// DeleteItemFromCart implements domain.IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, userId int64, warehouseId int64) error {
	return p.cart.RemoveItem(ctx, userId, warehouseId)
}

// AddToCart implements domain.IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart domain.CartInsertReq) error {
	return p.cart.Insert(
		ctx, cart.UserId, cart.WarehouseId, cart.Quantity)
}

// GetCart implements domain.IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userId int64, limit int) (*domain.CartSchema, error) {
	return p.cart.GetCartByUserID(ctx, userId, limit)
}

// GetOrders implements domain.IPurchaseService.
func (p *Purchase) GetOrders(ctx context.Context, limit int) ([]domain.Orders, error) {
	panic("unimplemented")
}

// InsertOrders implements domain.IPurchaseService.
func (p *Purchase) InsertOrders(ctx context.Context, order domain.Orders) error {
	panic("unimplemented")
}
