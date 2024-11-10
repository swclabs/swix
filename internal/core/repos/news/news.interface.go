package news

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

type INews interface {
	Create(ctx context.Context, news entity.News) (int64, error)
	GetMany(ctx context.Context, category string, limit int) ([]entity.News, error)
	UploadNewsImage(ctx context.Context, newsID int64, url string) error
	CreateHeadline(ctx context.Context, headline entity.News) error
}
