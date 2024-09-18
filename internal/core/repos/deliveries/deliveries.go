package deliveries

import (
	"context"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

var _ IDeliveries = (*Deliveries)(nil)
var _ = boot.Repos(Init)

// Init creates a new Deliveries object
func Init(_ cache.ICache, db db.IDatabase) IDeliveries {
	return &Deliveries{
		db: db,
	}
}

// Deliveries struct for delivery repos
type Deliveries struct {
	db db.IDatabase
}

// Create implements IDelivery.
func (d *Deliveries) Create(ctx context.Context, delivery entity.Deliveries) error {
	return d.db.SafeWrite(ctx, insert,
		delivery.UserID, delivery.AddressID, delivery.Status, delivery.Method, delivery.Note,
		delivery.SentDate, delivery.ReceivedDate)
}

// GetByID implements IDelivery.
func (d *Deliveries) GetByID(ctx context.Context, ID int64) (*entity.Deliveries, error) {
	raw, err := d.db.Query(ctx, selectByID, ID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectOneRow[entity.Deliveries](raw)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByUserID implements IDelivery.
func (d *Deliveries) GetByUserID(ctx context.Context, userID int64) ([]entity.Deliveries, error) {
	raw, err := d.db.Query(ctx, selectByUserID, userID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRows[entity.Deliveries](raw)
	if err != nil {
		return nil, err
	}
	return result, nil
}
