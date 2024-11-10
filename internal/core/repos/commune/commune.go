package commune

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(db db.IDatabase) ICommune {
	return &Commune{db: db}
}

type Commune struct {
	db db.IDatabase
}

// GetByDistrictID implements ICommune.
func (c *Commune) GetByDistrictID(ctx context.Context, districtID string) ([]entity.Commune, error) {
	rows, err := c.db.Query(ctx, getByDistrictID, districtID)
	if err != nil {
		return nil, err
	}

	communes, err := db.CollectRows[entity.Commune](rows)
	if err != nil {
		return nil, err
	}
	return communes, nil
}
