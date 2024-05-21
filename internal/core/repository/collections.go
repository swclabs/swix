package repository

import (
	"context"
	"encoding/json"
	"log"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/db/queries"

	"gorm.io/gorm"
)

type Collections struct {
	conn *gorm.DB
}

var _ domain.ICollections = (*Collections)(nil)

func NewCardBannerCollection() domain.ICollections {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Collections{
		conn: _conn,
	}
}

func (collection *Collections) UploadCollectionImage(ctx context.Context, collectionId string, url string) error {
	return db.SafeWriteQuery(
		ctx,
		collection.conn,
		queries.UpdateCollectionImage,
		url, collectionId,
	)
}

func (collection *Collections) AddCollection(ctx context.Context, collectionType domain.CollectionType) (int64, error) {
	_collection, err := json.Marshal(collectionType.Body)
	if err != nil {
		return -1, err
	}
	return db.SafeWriteQueryReturnId(
		ctx,
		collection.conn,
		queries.InsertIntoCollections,
		collectionType.Position, collectionType.Headline, string(_collection),
	)
}

func (collection *Collections) SlicesOfCollections(ctx context.Context, position string, limit int) (*domain.Collections, error) {
	var collections []domain.Collection
	if err := collection.conn.WithContext(ctx).
		Raw(queries.SelectCollectionByPosition, position, limit).Scan(&collections).Error; err != nil {
		return nil, err
	}
	var _collections domain.Collections
	_collections.Position = collections[0].Position
	_collections.Headline = collections[0].Headline

	for _, _collection := range collections {
		var body domain.CollectionBody
		if err := json.Unmarshal([]byte(_collection.Body), &body); err != nil {
			return nil, err
		}
		_collections.CardBanner = append(_collections.CardBanner, domain.CollectionsBody{
			CollectionBody: body,
			Id:             _collection.Id,
			Created:        _collection.Created,
		})
	}
	return &_collections, nil
}
