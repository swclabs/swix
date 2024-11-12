package district

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(db db.IDatabase) IDistrict {
	return &District{db: db}
}

type District struct {
	db db.IDatabase
}

// GetByProvinceID implements IDistrict.
func (d *District) GetByProvinceID(ctx context.Context, provinceID string) ([]entity.District, error) {
	rows, err := d.db.Query(ctx, getByProvinceID, provinceID)
	if err != nil {
		return nil, err
	}

	districts, err := db.CollectRows[entity.District](rows)
	if err != nil {
		return nil, err
	}
	return districts, nil
}
