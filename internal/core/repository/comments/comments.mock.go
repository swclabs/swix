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
func (p *Mock) Insert(ctx context.Context, cmt entity.Comment) (int64, error) {
	args := p.Called(ctx, cmt)
	return args.Get(0).(int64), args.Error(1)
}
