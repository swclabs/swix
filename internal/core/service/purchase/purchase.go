package purchase

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type Purchase struct {
	// PurchaseTask
}

func NewPurchasingService() IPurchaseService {
	return &Purchase{}
}

// AddToCart implements domain.IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart domain.CartInfo) {
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
