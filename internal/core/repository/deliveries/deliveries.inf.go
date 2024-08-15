package deliveries

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// IDelivery is an interface for Delivery.
type IDelivery interface {
	Create(ctx context.Context, delivery entity.Deliveries) error
	GetByID(ctx context.Context, ID int64) (*entity.Deliveries, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Deliveries, error)
}
