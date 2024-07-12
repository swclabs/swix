package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICollections Bind form json data to structure data
type ICollections interface {
	AddCollection(ctx context.Context, banner domain.CollectionSchema) (int64, error)
	SlicesOfCollections(ctx context.Context, position string, limit int) ([]domain.Collection, error)
	UploadCollectionImage(ctx context.Context, collectionID string, url string) error
	AddHeadlineBanner(ctx context.Context, headline domain.HeadlineBannerSchema) error
}
