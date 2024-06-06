package collections

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"gorm.io/gorm"
)

type Collections struct {
	conn *gorm.DB
}

var _ ICollections = (*Collections)(nil)

func New(conn *gorm.DB) *Collections {
	return &Collections{
		conn: conn,
	}
}

func (collection *Collections) UploadCollectionImage(
	ctx context.Context, collectionId string, url string) error {
	return db.SafeWriteQuery(
		ctx, collection.conn, UpdateCollectionImage,
		url, collectionId,
	)
}

func (collection *Collections) AddCollection(
	ctx context.Context, collectionType domain.CollectionType) (int64, error) {
	_collection, err := json.Marshal(collectionType.Body)
	if err != nil {
		return -1, err
	}
	return db.SafeWriteQueryReturnId(
		ctx, collection.conn, InsertIntoCollections,
		collectionType.Position, collectionType.Headline, string(_collection),
	)
}

func (collection *Collections) SlicesOfCollections(
	ctx context.Context, position string, limit int) ([]domain.Collection, error) {
	var collections []domain.Collection
	if err := collection.conn.WithContext(ctx).
		Raw(SelectCollectionByPosition, position, limit).
		Scan(&collections).Error; err != nil {
		return nil, err
	}
	return collections, nil
}

// AddHeadlineBanner implements domain.IHeadlineBannerCollections.
func (collection *Collections) AddHeadlineBanner(
	ctx context.Context, headline domain.HeadlineBannerType) error {
	body, err := json.Marshal(headline.Body)
	if err != nil {
		return err
	}
	return db.SafeWriteQuery(
		ctx, collection.conn, InsertIntoCollections, headline.Position, "", string(body))
}
