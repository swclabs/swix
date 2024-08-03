package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ ICollections = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

// AddHeadlineBanner implements ICollections.
func (m *Mock) AddHeadlineBanner(ctx context.Context, headline entity.Collection) error {
	panic("unimplemented")
}

// Create implements ICollections.
func (m *Mock) Create(ctx context.Context, banner entity.Collection) (int64, error) {
	panic("unimplemented")
}

// GetMany implements ICollections.
func (m *Mock) GetMany(ctx context.Context, position string, limit int) ([]entity.Collection, error) {
	args := m.Called(ctx, position, limit)
	return args.Get(0).([]entity.Collection), args.Error(1)
}

// UploadCollectionImage implements ICollections.
func (m *Mock) UploadCollectionImage(ctx context.Context, collectionID string, url string) error {
	panic("unimplemented")
}
