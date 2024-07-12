package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

// Mock is a mock type for ICategoriesRepository.
type Mock struct {
	mock.Mock
}

var _ ICategoriesRepository = (*Mock)(nil)

// NewCategoriesMock creates a new mock object.
func NewCategoriesMock() *Mock {
	return &Mock{}
}

// GetLimit implements ICategoriesRepository.
func (c *Mock) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	args := c.Called(ctx, limit)
	return args.Get(0).([]domain.Categories), args.Error(1)
}

// Insert implements ICategoriesRepository.
func (c *Mock) Insert(ctx context.Context, ctg domain.Categories) error {
	args := c.Called(ctx, ctg)
	return args.Error(0)
}
