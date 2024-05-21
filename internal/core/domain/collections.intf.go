package domain

import "context"

type Any interface{}

type IUtilsJsonDocuments[T any] interface {
	Bind(in []byte, out *T) error
}

// ICollections Bind form json data to structure data
type ICollections interface {
	AddCollection(ctx context.Context, banner CollectionType) (int64, error)
	SlicesOfCollections(ctx context.Context, position string, limit int) (*Collections, error)
	UploadCollectionImage(ctx context.Context, collectionId string, url string) error
}
