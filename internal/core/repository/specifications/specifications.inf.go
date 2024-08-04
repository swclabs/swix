package specifications

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
)

type ISpecifications interface {
	GetByInventoryID(ctx context.Context, inventoryID int64) ([]entity.Specifications, error)
	Insert(ctx context.Context, specs entity.Specifications) error
}
