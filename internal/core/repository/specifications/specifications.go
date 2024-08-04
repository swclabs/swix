// Package specifications contains the Specifications repository.
package specifications

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
	"swclabs/swipecore/pkg/infra/db"
)

// Init initializes the Specifications repository.
func Init(conn db.IDatabase, _ cache.ICache) ISpecifications {
	return &Specifications{
		db: conn,
	}
}

// Specifications represents the Specifications repository.
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
