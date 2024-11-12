package province

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
)

type IProvince interface {
	GetAll(ctx context.Context) ([]entity.Province, error)
}
