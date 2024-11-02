package news

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(New)

func New(conn db.IDatabase) INews {
	return &News{
		db: conn,
	}
}

type News struct {
	db db.IDatabase
}

// Create implements INews.
func (n *News) Create(ctx context.Context, news entity.News) (int64, error) {
	return n.db.SafeWriteReturn(
		ctx, insertIntoNews,
		news.Category, news.Header, news.Body,
	)
}

// CreateHeadline implements INews.
func (n *News) CreateHeadline(ctx context.Context, headline entity.News) error {
	return n.db.SafeWrite(
		ctx, insertIntoNews, headline.Category, headline.Header, "")
}

// GetMany implements INews.
func (n *News) GetMany(ctx context.Context, category string, limit int) ([]entity.News, error) {
	rows, err := n.db.Query(ctx, selectByCategory, category, limit)
	if err != nil {
		return nil, err
	}
	newss, err := db.CollectRows[entity.News](rows)
	if err != nil {
		return nil, err
	}
	return newss, nil
}

// UploadNewsImage implements INews.
func (n *News) UploadNewsImage(ctx context.Context, newsID int64, url string) error {
	return n.db.SafeWrite(
		ctx, updateNewsImage,
		url, newsID,
	)
}
