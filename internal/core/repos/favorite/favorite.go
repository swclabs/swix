package favorite

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(conn db.IDatabase) IFavorite {
	return &Favorite{db: conn}
}

type Favorite struct {
	db db.IDatabase
}

// Create implements IFavorite.
func (f *Favorite) Create(ctx context.Context, favorite entity.Favorite) error {
	return f.db.SafeWrite(ctx, insert, favorite.UserID, favorite.InventoryID)
}

// Delete implements IFavorite.
func (f *Favorite) Delete(ctx context.Context, favorite entity.Favorite) error {
	return f.db.SafeWrite(ctx, delete, favorite.UserID, favorite.InventoryID)
}

// GetByInventoryID implements IFavorite.
func (f *Favorite) GetByInventoryID(ctx context.Context, inventoryID int64, userId int64) (*entity.Favorite, error) {
	row, err := f.db.Query(ctx, getByInventoryID, inventoryID, userId)
	if err != nil {
		return nil, err
	}
	favorite, err := db.CollectRow[entity.Favorite](row)
	if err != nil {
		return nil, err
	}
	return &favorite, nil
}

// GetByUserID implements IFavorite.
func (f *Favorite) GetByUserID(ctx context.Context, userID int64) ([]entity.Favorite, error) {
	rows, err := f.db.Query(ctx, getByUserID, userID)
	if err != nil {
		return nil, err
	}
	favorites, err := db.CollectRows[entity.Favorite](rows)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
