package specifications

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
	"swclabs/swipecore/pkg/infra/db"
)

func Init(conn db.IDatabase, cache cache.ICache) ISpecifications {
	return &Specifications{
		db: conn,
	}
}

type Specifications struct {
	db db.IDatabase
}

// GetByInventoryID implements ISpecifications.
func (s *Specifications) GetByInventoryID(ctx context.Context, inventoryID int64) ([]entity.Specifications, error) {
	rows, err := s.db.Query(ctx, getByInventoryID, inventoryID)
	if err != nil {
		return nil, err
	}
	specs, err := db.CollectRows[entity.Specifications](rows)
	if err != nil {
		return nil, err
	}
	return specs, nil
}

// Insert implements ISpecifications.
func (s *Specifications) Insert(ctx context.Context, specs entity.Specifications) error {
	return s.db.SafeWrite(ctx, insert, specs.InventoryID, specs.Content)
}
