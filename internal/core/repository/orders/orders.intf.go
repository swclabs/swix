package orders

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IOrdersRepository interface for orders repository
type IOrdersRepository interface {
	Create(ctx context.Context, order domain.Orders) (int64, error)
	Get(ctx context.Context, userID int64, limit int) ([]domain.Orders, error)

	InsertProduct(ctx context.Context, product domain.ProductInOrder) error
	GetProductByOrderID(ctx context.Context, orderID int64) ([]domain.ProductInOrder, error)
}
