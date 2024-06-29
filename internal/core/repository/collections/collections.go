package collections

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Collections struct {
	db db.IDatabase
}

var _ ICollections = (*Collections)(nil)

func New(conn db.IDatabase) ICollections {
	return useCache(&Collections{
		db: conn,
	})
}

func (collection *Collections) UploadCollectionImage(
	ctx context.Context, collectionId string, url string) error {
	return collection.db.SafeWrite(
		ctx, updateCollectionImage,
		url, collectionId,
	)
}

func (collection *Collections) AddCollection(
	ctx context.Context, collectionType domain.CollectionSchema) (int64, error) {
	_collection, err := json.Marshal(collectionType.Body)
	if err != nil {
		return -1, err
	}
	return collection.db.SafeWriteReturn(
		ctx, insertIntoCollections,
		collectionType.Position, collectionType.Headline, string(_collection),
	)
}

func (collection *Collections) SlicesOfCollections(
	ctx context.Context, position string, limit int) ([]domain.Collection, error) {
	rows, err := collection.db.Query(ctx, selectCollectionByPosition, position, limit)
	if err != nil {
		return nil, err
	}
	collections, err := db.CollectRows[domain.Collection](rows)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// AddHeadlineBanner implements domain.IHeadlineBannerCollections.
func (collection *Collections) AddHeadlineBanner(
	ctx context.Context, headline domain.HeadlineBannerSchema) error {
	body, err := json.Marshal(headline.Body)
	if err != nil {
		return err
	}
	return collection.db.SafeWrite(
		ctx, insertIntoCollections, headline.Position, "", string(body))
}
