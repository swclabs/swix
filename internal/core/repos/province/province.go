package province

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(db db.IDatabase) IProvince {
	return &Province{db: db}
}

type Province struct {
	db db.IDatabase
}

// GetAll implements IProvince.
func (p *Province) GetAll(ctx context.Context) ([]entity.Province, error) {
	rows, err := p.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}

	provinces, err := db.CollectRows[entity.Province](rows)
	if err != nil {
		return nil, err
	}
	return provinces, nil
}
