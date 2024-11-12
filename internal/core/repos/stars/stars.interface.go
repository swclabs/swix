package stars

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
)

type IStar interface {
	Save(ctx context.Context, star entity.Star) error
}
