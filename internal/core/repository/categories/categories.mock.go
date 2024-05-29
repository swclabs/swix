package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type CategoriesMock struct {
	mock.Mock
}

var _ ICategoriesRepository = (*CategoriesMock)(nil)

func NewCategoriesMock() *CategoriesMock {
	return &CategoriesMock{}
}

// GetLimit implements domain.ICategoriesRepository.
func (c *CategoriesMock) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	panic("unimplemented")
}

// Insert implements domain.ICategoriesRepository.
func (c *CategoriesMock) Insert(ctx context.Context, ctg *domain.Categories) error {
	panic("unimplemented")
}
