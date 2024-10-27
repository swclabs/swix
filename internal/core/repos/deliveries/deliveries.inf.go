package deliveries

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

// IDeliveries is an interface for Delivery.
type IDeliveries interface {
	Create(ctx context.Context, delivery entity.Delivery) (int64, error)
	GetByID(ctx context.Context, ID int64) (*entity.Delivery, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.Delivery, error)
}
