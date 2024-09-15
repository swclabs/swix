package collections

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

// Collections struct for collections
type Collections struct {
	db db.IDatabase
}

var _ ICollections = (*Collections)(nil)

// New creates a new Collections object
func New(conn db.IDatabase) ICollections {
	return &Collections{
		db: conn,
	}
}

// Init initializes the Collections object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) ICollections {
	return useCache(cache, New(conn))
}

// UploadCollectionImage implements domain.ICollections.
func (collection *Collections) UploadCollectionImage(
	ctx context.Context, collectionID string, url string) error {
	return collection.db.SafeWrite(
		ctx, updateCollectionImage,
		url, collectionID,
	)
}

// Create implements domain.ICollections.
func (collection *Collections) Create(
	ctx context.Context, collectionType entity.Collection) (int64, error) {
	return collection.db.SafeWriteReturn(
		ctx, insertIntoCollections,
		collectionType.Position, collectionType.Headline, collectionType.Body,
	)
}

// GetMany implements domain.ICollections.
func (collection *Collections) GetMany(
	ctx context.Context, position string, limit int) ([]entity.Collection, error) {
	rows, err := collection.db.Query(ctx, selectCollectionByPosition, position, limit)
	if err != nil {
		return nil, err
	}
	collections, err := db.CollectRows[entity.Collection](rows)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// AddHeadlineBanner implements domain.IHeadlineBannerCollections.
func (collection *Collections) AddHeadlineBanner(
	ctx context.Context, headline entity.Collection) error {
	return collection.db.SafeWrite(
		ctx, insertIntoCollections, headline.Position, "", headline.Body)
}
