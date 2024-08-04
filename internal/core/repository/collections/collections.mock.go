package collections

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ ICollections = (*Mock)(nil)

// Mock is a mock type for ICollections.
type Mock struct {
	mock.Mock
}

// AddHeadlineBanner implements ICollections.
func (m *Mock) AddHeadlineBanner(_ context.Context, _ entity.Collection) error {
	panic("unimplemented")
}

// Create implements ICollections.
func (m *Mock) Create(_ context.Context, _ entity.Collection) (int64, error) {
	panic("unimplemented")
}

// GetMany implements ICollections.
func (m *Mock) GetMany(ctx context.Context, position string, limit int) ([]entity.Collection, error) {
	args := m.Called(ctx, position, limit)
	return args.Get(0).([]entity.Collection), args.Error(1)
}

// UploadCollectionImage implements ICollections.
func (m *Mock) UploadCollectionImage(_ context.Context, _ string, _ string) error {
	panic("unimplemented")
}
