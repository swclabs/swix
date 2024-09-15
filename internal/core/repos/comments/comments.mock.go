package comments

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ ICommentRepository = (*Mock)(nil)

// NewProductsMock creates a new Mock.
func NewCommentsMock() *Mock {
	return &Mock{}
}

// Mock is a mock for IProductRepository.
type Mock struct {
	mock.Mock
}

// Insert implements ICommentRepository.
func (p *Mock) Insert(ctx context.Context, cmt entity.Comments) (int64, error) {
	args := p.Called(ctx, cmt)
	return args.Get(0).(int64), args.Error(1)
}

// GetByID implements ICommentRepository.
func (p *Mock) GetByID(ctx context.Context, ID int64) (*entity.Comments, error) {
	args := p.Called(ctx, ID)
	return args.Get(0).(*entity.Comments), args.Error(1)
}

// Update implements ICommentRepository.
func (p *Mock) Update(ctx context.Context, cmt entity.Comments) error {
	args := p.Called(ctx, cmt)
	return args.Error(0)
}

// GetByProductID implements ICommentRepository.
func (p *Mock) GetByProductID(ctx context.Context, ID int64) ([]entity.Comments, error) {
	args := p.Called(ctx, ID)
	return args.Get(0).([]entity.Comments), args.Error(1)
}

// Delete implements ICommentRepository.
func (p *Mock) DeleteByID(ctx context.Context, ID int64) error {
	args := p.Called(ctx, ID)
	return args.Error(0)
}