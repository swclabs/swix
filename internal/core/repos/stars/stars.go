package stars

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/db"
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

// Save implements IStar.
func (s *Star) Save(ctx context.Context, star entity.Star) error {
	return s.db.SafeWrite(ctx, insertStar, star.ProductID, star.UserID)
}
