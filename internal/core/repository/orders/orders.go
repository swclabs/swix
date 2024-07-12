package orders

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

// Orders represents the repository for orders
type Orders struct {
	db db.IDatabase
}

// New creates a new Orders object
func New(conn db.IDatabase) IOrdersRepository {
	return &Orders{
		db: conn,
	}
}

var _ IOrdersRepository = (*Orders)(nil)

// InsertProduct implements IOrdersRepository.
func (orders *Orders) InsertProduct(ctx context.Context, product domain.ProductInOrder) error {
	return orders.db.SafeWrite(ctx, insertProductToOrder,
		product.OrderID, product.InventoryID, product.Quantity, "VND",
		product.TotalAmount.String(),
	)
}

// Create implements IOrdersRepository.
func (orders *Orders) Create(ctx context.Context, order domain.Orders) (int64, error) {
	return orders.db.SafeWriteReturn(ctx, insertOrder,
		order.UUID, order.UserID, "active", order.TotalAmount.String(),
	)
}
