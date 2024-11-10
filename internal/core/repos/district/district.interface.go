package district

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

type IDistrict interface {
	GetByProvinceID(ctx context.Context, provinceID string) ([]entity.District, error)
}
