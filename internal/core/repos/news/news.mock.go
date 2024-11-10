package news

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ INews = (*Mock)(nil)

// Mock is a mock type for INews.
type Mock struct {
	mock.Mock
}

// Create implements INews.
func (m *Mock) Create(ctx context.Context, news entity.News) (int64, error) {
	panic("unimplemented")
}

// CreateHeadline implements INews.
func (m *Mock) CreateHeadline(ctx context.Context, headline entity.News) error {
	panic("unimplemented")
}

// GetMany implements INews.
func (m *Mock) GetMany(ctx context.Context, category string, limit int) ([]entity.News, error) {
	args := m.Called(ctx, category, limit)
	return args.Get(0).([]entity.News), args.Error(1)
}

// UploadNewsImage implements INews.
func (m *Mock) UploadNewsImage(ctx context.Context, newsID int64, url string) error {
	panic("unimplemented")
}
