package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
)

// ICollections Bind form json data to structure data
type ICollections interface {
	AddCollection(ctx context.Context, banner entity.Collection) (int64, error)
	SlicesOfCollections(ctx context.Context, position string, limit int) ([]entity.Collection, error)
	UploadCollectionImage(ctx context.Context, collectionID string, url string) error
	AddHeadlineBanner(ctx context.Context, headline entity.Collection) error
}
