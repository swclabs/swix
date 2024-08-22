package categories

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

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

// GetByID implements ICategoriesRepository.
func (c *Mock) GetByID(ctx context.Context, ID int64) (*entity.Categories, error) {
	args := c.Called(ctx, ID)
	return args.Get(0).(*entity.Categories), args.Error(1)
}

// GetLimit implements ICategoriesRepository.
func (c *Mock) GetLimit(ctx context.Context, limit string) ([]entity.Categories, error) {
	args := c.Called(ctx, limit)
	return args.Get(0).([]entity.Categories), args.Error(1)
}

// Insert implements ICategoriesRepository.
func (c *Mock) Insert(ctx context.Context, ctg entity.Categories) error {
	args := c.Called(ctx, ctg)
	return args.Error(0)
}

// DeleteByID implements ICategoriesRepository.
func (c *Mock) DeleteByID(ctx context.Context, ID int64) error {
	args := c.Called(ctx, ID)
	return args.Error(0)
}

// Update implements IProductRepository.
func (c *Mock) Update(ctx context.Context, category entity.Categories) error {
	args := c.Called(ctx, category)
	return args.Error(0)
}
