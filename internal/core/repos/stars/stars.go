package stars

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(conn db.IDatabase) IStar {
	return &Star{
		db: conn,
	}
}

type Star struct {
	db db.IDatabase
}

// GetByProductID implements IStar.
func (s *Star) GetByProductID(ctx context.Context, productID int64) ([]entity.Star, error) {
	rows, err := s.db.Query(ctx, getByProductID, productID)
	if err != nil {
		return nil, err
	}
	stars, err := db.CollectRows[entity.Star](rows)
	if err != nil {
		return nil, err
	}
	return stars, nil
}

// Save implements IStar.
func (s *Star) Save(ctx context.Context, star entity.Star) (int64, error) {
	return s.db.SafeWriteReturn(ctx, insertStar, star.ProductID, star.UserID, star.Star)
}
