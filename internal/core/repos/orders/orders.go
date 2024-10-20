package orders

import (
	"context"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

// New creates a new Orders object
func New(conn db.IDatabase) IOrders {
	return &Orders{
		db: conn,
	}
}

var _ = app.Repos(Init)

// Init initializes the Orders object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IOrders {
	return useCache(cache, New(conn))
}

var _ IOrders = (*Orders)(nil)

// Orders represents the repos for orders
type Orders struct {
	db db.IDatabase
}

// GetOrderItemByCode implements IOrders.
func (orders *Orders) GetOrderItemByCode(ctx context.Context, orderCode string) ([]model.Order, error) {
	rows, err := orders.db.Query(ctx, getByOrderCode, orderCode)
	if err != nil {
		return nil, err
	}
	order, err := db.CollectRows[model.Order](rows)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetByUUID implements IOrders.
func (orders *Orders) GetByUUID(ctx context.Context, uuid string) (*entity.Orders, error) {
	rows, err := orders.db.Query(ctx, getByUUID, uuid)
	if err != nil {
		return nil, err
	}
	order, err := db.CollectRow[entity.Orders](rows)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// InsertProduct implements IOrdersRepository.
func (orders *Orders) InsertProduct(ctx context.Context, product entity.ProductInOrder) error {
	return orders.db.SafeWrite(ctx, insertProductToOrder,
		product.OrderID, product.InventoryID, product.Quantity, "VND",
		product.TotalAmount.String(), product.SpecsID,
	)
}

// Create implements IOrdersRepository.
func (orders *Orders) Create(ctx context.Context, order entity.Orders) (int64, error) {
	return orders.db.SafeWriteReturn(ctx, insertOrder,
		order.UUID, order.UserID, "active", order.TotalAmount.String(), order.DeliveryID,
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
