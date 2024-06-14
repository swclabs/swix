package collections

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Collections struct {
	conn *pgx.Conn
}

var _ ICollections = (*Collections)(nil)

func New(conn *pgx.Conn) *Collections {
	return &Collections{
		conn: conn,
	}
}

func (collection *Collections) UploadCollectionImage(
	ctx context.Context, collectionId string, url string) error {
	return db.SafePgxWriteQuery(
		ctx, collection.conn, UpdateCollectionImage,
		url, collectionId,
	)
}

func (collection *Collections) AddCollection(
	ctx context.Context, collectionType domain.CollectionSchema) (int64, error) {
	_collection, err := json.Marshal(collectionType.Body)
	if err != nil {
		return -1, err
	}
	return db.SafePgxWriteQueryReturnId(
		ctx, collection.conn, InsertIntoCollections,
		collectionType.Position, collectionType.Headline, string(_collection),
	)
}

func (collection *Collections) SlicesOfCollections(
	ctx context.Context, position string, limit int) ([]domain.Collection, error) {
	rows, err := collection.conn.Query(ctx, SelectCollectionByPosition, position, limit)
	if err != nil {
		return nil, err
	}
	collections, err := pgx.CollectRows[domain.Collection](rows, pgx.RowToStructByName[domain.Collection])
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
	return db.SafePgxWriteQuery(
		ctx, collection.conn, InsertIntoCollections, headline.Position, "", string(body))
}
