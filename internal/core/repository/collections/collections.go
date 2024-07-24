package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/db"
)

// Collections struct for collections
type Collections struct {
	db db.IDatabase
}

var _ ICollections = (*Collections)(nil)

// New creates a new Collections object
func New(conn db.IDatabase) ICollections {
	return useCache(&Collections{
		db: conn,
	})
}

// UploadCollectionImage implements domain.ICollections.
func (collection *Collections) UploadCollectionImage(
	ctx context.Context, collectionID string, url string) error {
	return collection.db.SafeWrite(
		ctx, updateCollectionImage,
		url, collectionID,
	)
}

// AddCollection implements domain.ICollections.
func (collection *Collections) AddCollection(
	ctx context.Context, collectionType entity.Collection) (int64, error) {
	return collection.db.SafeWriteReturn(
		ctx, insertIntoCollections,
		collectionType.Position, collectionType.Headline, collectionType.Body,
	)
}

// SlicesOfCollections implements domain.ICollections.
func (collection *Collections) SlicesOfCollections(
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
