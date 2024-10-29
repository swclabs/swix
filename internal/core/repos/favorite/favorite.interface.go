package favorite

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

type IFavorite interface {
	Save(ctx context.Context, favorite entity.Favorite) error
	Delete(ctx context.Context, favorite entity.Favorite) error
	GetByUserID(ctx context.Context, userID int64) ([]entity.Favorite, error)
	GetByInventoryID(ctx context.Context, inventoryID int64, userId int64) (*entity.Favorite, error)
}
