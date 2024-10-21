package deliveries

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// IDeliveries is an interface for Delivery.
type IDeliveries interface {
	Create(ctx context.Context, delivery entity.Deliveries) (int64, error)
	GetByID(ctx context.Context, ID int64) (*entity.Deliveries, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Deliveries, error)
}
