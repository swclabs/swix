package specifications

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ ISpecifications = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

// GetByInventoryID implements ISpecifications.
func (m *Mock) GetByInventoryID(ctx context.Context, inventoryID int64) ([]entity.Specifications, error) {
	args := m.Called(ctx, inventoryID)
	return args.Get(0).([]entity.Specifications), args.Error(1)
}

// Insert implements ISpecifications.
func (m *Mock) Insert(ctx context.Context, specs entity.Specifications) error {
	panic("unimplemented")
}
