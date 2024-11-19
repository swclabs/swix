package stars

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
)

type IStar interface {
	Save(ctx context.Context, star entity.Star) (int64, error)
	GetByProductID(ctx context.Context, productID int64) ([]entity.Star, error)
}
