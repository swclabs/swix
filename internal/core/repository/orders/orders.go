package orders

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

type Orders struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IOrdersRepository {
	return &Orders{
		db: conn,
	}
}

var _ IOrdersRepository = (*Orders)(nil)

// InsertProduct implements IOrdersRepository.
func (orders *Orders) InsertProduct(ctx context.Context, product domain.ProductInOrder) error {
	return orders.db.SafeWrite(ctx, insertProductToOrder,
		product.OrderId, product.InventoryId, product.Quantity, "VND",
		product.TotalAmount.String(),
	)
}

// Create implements IOrdersRepository.
func (orders *Orders) Create(ctx context.Context, order domain.Orders) (int64, error) {
	return orders.db.SafeWriteReturn(ctx, insertOrder,
		order.Uuid, order.UserId, "active", order.TotalAmount.String(),
	)
}
