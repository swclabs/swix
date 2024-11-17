package orders

import (
	"context"
	"fmt"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/infra/db"
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

// UpdateStatus implements IOrders.
func (orders *Orders) UpdateStatus(ctx context.Context, orderCode string, status string) error {
	return orders.db.SafeWrite(ctx, updateStatus, status, orderCode)
}

func (orders *Orders) GetLimit(ctx context.Context, limit int) ([]entity.Order, error) {
	rows, err := orders.db.Query(ctx, getLimit, limit)
	if err != nil {
		return nil, err
	}
	ords, err := db.CollectRows[entity.Order](rows)
	if err != nil {
		return nil, err
	}
	return ords, nil
}

// GetItemByCode implements IOrders.
func (orders *Orders) GetItemByCode(ctx context.Context, orderCode string) ([]model.Order, error) {
	rows, err := orders.db.Query(ctx, getByOrderCode, orderCode)
	if err != nil {
		return nil, fmt.Errorf("error getting order items by order code: %w", err)
	}

	order, err := db.CollectRows[model.Order](rows)
	if err != nil {
		return nil, fmt.Errorf("error collecting order items by order code: %w", err)
	}
	return order, nil
}

// GetByUUID implements IOrders.
func (orders *Orders) GetByUUID(ctx context.Context, uuid string) (*entity.Order, error) {
	rows, err := orders.db.Query(ctx, getByUUID, uuid)
	if err != nil {
		return nil, err
	}
	order, err := db.CollectRow[entity.Order](rows)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

// InsertProduct implements IOrdersRepository.
func (orders *Orders) InsertProduct(ctx context.Context, product entity.ProductInOrder) error {
	return orders.db.SafeWrite(ctx, insertProductToOrder,
		product.OrderID, product.InventoryID, product.Quantity, "VND",
		product.TotalAmount.String(),
	)
}

// Create implements IOrdersRepository.
func (orders *Orders) Create(ctx context.Context, order entity.Order) (int64, error) {
	return orders.db.SafeWriteReturn(ctx, insertOrder,
		order.UUID, order.UserID, "active", order.TotalAmount.String(), order.DeliveryID,
	)
}

// GetByUserID implements IOrdersRepository.
func (orders *Orders) GetByUserID(ctx context.Context, userID int64, limit int) ([]entity.Order, error) {
	rows, err := orders.db.Query(ctx, getOrder, userID, limit)
	if err != nil {
		return nil, err
	}
	_orders, err := db.CollectRows[entity.Order](rows)
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
