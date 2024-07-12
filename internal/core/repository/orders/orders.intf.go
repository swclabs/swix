package orders

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IOrdersRepository interface for orders repository
type IOrdersRepository interface {
	Create(ctx context.Context, order domain.Orders) (int64, error)
	InsertProduct(ctx context.Context, product domain.ProductInOrder) error
}
