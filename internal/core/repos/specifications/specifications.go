// Package specifications contains the Specifications repos.
package specifications

import (
	"context"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

var _ = boot.Repos(Init)

// Init initializes the Specifications repos.
func Init(conn db.IDatabase, _ cache.ICache) ISpecifications {
	return &Specifications{
		db: conn,
	}
}

// New creates a new Specifications object.
func New(conn db.IDatabase) ISpecifications {
	return &Specifications{
		db: conn,
	}
}

// Specifications represents the Specifications repos.
type Specifications struct {
	db db.IDatabase
}

// GetByID implements ISpecifications.
func (s *Specifications) GetByID(ctx context.Context, ID int64) (*entity.Specifications, error) {
	row, err := s.db.Query(ctx, getByID, ID)
	if err != nil {
		return nil, err
	}
	specs, err := db.CollectOneRow[entity.Specifications](row)
	if err != nil {
		return nil, err
	}
	return &specs, nil
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
