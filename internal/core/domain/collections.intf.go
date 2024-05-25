package domain

import "context"

// ICollections Bind form json data to structure data
type ICollections interface {
	AddCollection(ctx context.Context, banner CollectionType) (int64, error)
	SlicesOfCollections(ctx context.Context, position string, limit int) ([]Collection, error)
	UploadCollectionImage(ctx context.Context, collectionId string, url string) error
	AddHeadlineBanner(ctx context.Context, headline HeadlineBannerType) error
}
