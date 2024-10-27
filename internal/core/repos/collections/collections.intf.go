package collections

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

// ICollections Bind form json data to structure data
type ICollections interface {
	Create(ctx context.Context, banner entity.Collection) (int64, error)
	GetMany(ctx context.Context, position string, limit int) ([]entity.Collection, error)
	UploadCollectionImage(ctx context.Context, collectionID string, url string) error
	AddHeadlineBanner(ctx context.Context, headline entity.Collection) error
}
