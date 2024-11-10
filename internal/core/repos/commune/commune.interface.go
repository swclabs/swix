package commune

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

type ICommune interface {
	GetByDistrictID(ctx context.Context, districtID string) ([]entity.Commune, error)
}
