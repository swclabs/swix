package province

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

type IProvince interface {
	GetAll(ctx context.Context) ([]entity.Province, error)
}
