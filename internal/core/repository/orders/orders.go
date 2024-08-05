package orders

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
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

// Init initializes the Orders object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IOrdersRepository {
	return useCache(cache, New(conn))
}

var _ IOrdersRepository = (*Orders)(nil)

// InsertProduct implements IOrdersRepository.
func (orders *Orders) InsertProduct(ctx context.Context, product entity.ProductInOrder) error {
	return orders.db.SafeWrite(ctx, insertProductToOrder,
		product.OrderID, product.InventoryID, product.Quantity, "VND",
		product.TotalAmount.String(),
	)
}

// Create implements IOrdersRepository.
func (orders *Orders) Create(ctx context.Context, order entity.Orders) (int64, error) {
	return orders.db.SafeWriteReturn(ctx, insertOrder,
		order.UUID, order.UserID, "active", order.TotalAmount.String(),
	)
}

// Get implements IOrdersRepository.
func (orders *Orders) Get(ctx context.Context, userID int64, limit int) ([]entity.Orders, error) {
	rows, err := orders.db.Query(ctx, getOrder, userID, limit)
	if err != nil {
		return nil, err
	}
	_orders, err := db.CollectRows[entity.Orders](rows)
	if err != nil {
		return nil, err
	}
	return _orders, nil
}

// GetProductByOrderID implements IOrdersRepository.
func (orders *Orders) GetProductByOrderID(ctx context.Context, orderID int64) ([]entity.ProductInOrder, error) {
	rows, err := orders.db.Query(ctx, getProductByOrderID, orderID)
	if err != nil {
		return nil, err
	}
	_products, err := db.CollectRows[entity.ProductInOrder](rows)
	if err != nil {
		return nil, err
	}
	return _products, nil
}
