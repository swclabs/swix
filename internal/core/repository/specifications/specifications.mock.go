package specifications

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ ISpecifications = (*Mock)(nil)

// Mock is a mock type for ISpecifications.
type Mock struct {
	mock.Mock
}

// GetByID implements ISpecifications.
func (m *Mock) GetByID(_ context.Context, _ int64) (*entity.Specifications, error) {
	panic("unimplemented")
}

// GetByInventoryID implements ISpecifications.
func (m *Mock) GetByInventoryID(ctx context.Context, inventoryID int64) ([]entity.Specifications, error) {
	args := m.Called(ctx, inventoryID)
	return args.Get(0).([]entity.Specifications), args.Error(1)
}

// Insert implements ISpecifications.
func (m *Mock) Insert(_ context.Context, _ entity.Specifications) error {
	panic("unimplemented")
}
