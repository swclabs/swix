package specifications

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// ISpecifications represents the interface for Specifications repository.
type ISpecifications interface {
	GetByInventoryID(ctx context.Context, inventoryID int64) ([]entity.Specifications, error)
	Insert(ctx context.Context, specs entity.Specifications) error
	GetByID(ctx context.Context, ID int64) (*entity.Specifications, error)
}
