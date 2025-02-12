package deliveries

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/infra/db"
)

var _ IDeliveries = (*Deliveries)(nil)
var _ = app.Repos(Init)

// Init creates a new Deliveries object
func Init(_ cache.ICache, db db.IDatabase) IDeliveries {
	return &Deliveries{
		db: db,
	}
}

func New(db db.IDatabase) IDeliveries {
	return &Deliveries{
		db: db,
	}
}

// Deliveries struct for delivery repos
type Deliveries struct {
	db db.IDatabase
}

// Create implements IDelivery.
func (d *Deliveries) Create(ctx context.Context, delivery entity.Delivery) (int64, error) {
	return d.db.SafeWriteReturn(ctx, insert,
		delivery.UserID, delivery.AddressID, delivery.Status, delivery.Method, delivery.Note,
		delivery.SentDate)
}

// GetByID implements IDelivery.
func (d *Deliveries) GetByID(ctx context.Context, ID int64) (*entity.Delivery, error) {
	raw, err := d.db.Query(ctx, selectByID, ID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRow[entity.Delivery](raw)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByUserID implements IDelivery.
func (d *Deliveries) GetByUserID(ctx context.Context, userID int64) ([]entity.Delivery, error) {
	raw, err := d.db.Query(ctx, selectByUserID, userID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRows[entity.Delivery](raw)
	if err != nil {
		return nil, err
	}
	return result, nil
}
