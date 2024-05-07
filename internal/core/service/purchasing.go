package service

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type Purchasing struct {
	// PurchasingTask
}

func NewPurchasingService() domain.IPurchasingService {
	return &Purchasing{}
}

// AddToCart implements domain.IPurchasingService.
func (p *Purchasing) AddToCart(ctx context.Context, cart domain.CartInfo) {
	panic("unimplemented")
}

// GetCart implements domain.IPurchasingService.
func (p *Purchasing) GetCart(ctx context.Context, limit int) ([]domain.Carts, error) {
	panic("unimplemented")
}

// GetOrders implements domain.IPurchasingService.
func (p *Purchasing) GetOrders(ctx context.Context, limit int) ([]domain.Orders, error) {
	panic("unimplemented")
}

// InsertOrders implements domain.IPurchasingService.
func (p *Purchasing) InsertOrders(ctx context.Context, order domain.Orders) error {
	panic("unimplemented")
}
