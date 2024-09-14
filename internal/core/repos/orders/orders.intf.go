package orders

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// IOrders interface for orders repos
type IOrders interface {
	Create(ctx context.Context, order entity.Orders) (int64, error)
	Get(ctx context.Context, userID int64, limit int) ([]entity.Orders, error)

	InsertProduct(ctx context.Context, product entity.ProductInOrder) error
	GetProductByOrderID(ctx context.Context, orderID int64) ([]entity.ProductInOrder, error)
}
