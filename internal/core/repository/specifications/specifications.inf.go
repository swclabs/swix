package specifications

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
)

// ISpecifications represents the interface for Specifications repository.
type ISpecifications interface {
	GetByInventoryID(ctx context.Context, inventoryID int64) ([]entity.Specifications, error)
	Insert(ctx context.Context, specs entity.Specifications) error
}
