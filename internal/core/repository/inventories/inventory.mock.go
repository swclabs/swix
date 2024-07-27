package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

var _ IInventoryRepository = (*Mock)(nil)

// Mock represents a mock for IInventoryRepository.
type Mock struct {
	mock.Mock
}

// Update implements IInventoryRepository.
func (w *Mock) Update(_ context.Context, _ entity.Inventories) error {
	panic("unimplemented")
}

// UploadImage implements IInventoryRepository.
func (w *Mock) UploadImage(_ context.Context, _ int, _ string) error {
	panic("unimplemented")
}

// DeleteByID implements IInventoryRepository.
func (w *Mock) DeleteByID(_ context.Context, _ int64) error {
	panic("unimplemented")
}

// GetLimit implements IInventoryRepository.
func (w *Mock) GetLimit(_ context.Context, _ int, _ int) ([]entity.Inventories, error) {
	panic("unimplemented")
}

// GetByProductID implements IInventoryRepository.
func (w *Mock) GetByProductID(_ context.Context, _ int64) ([]entity.Inventories, error) {
	//TODO implement me
	panic("implement me")
}

// GetByID implements IInventoryRepository.
func (w *Mock) GetByID(_ context.Context, _ int64) (*entity.Inventories, error) {
	panic("unimplemented")
}

// FindDevice implements IInventoryRepository.
func (w *Mock) FindDevice(ctx context.Context, device dtos.InventoryDeviceSpecs) (*entity.Inventories, error) {
	args := w.Called(ctx, device)
	return args.Get(0).(*entity.Inventories), args.Error(1)
}

// InsertProduct implements IInventoryRepository.
func (w *Mock) InsertProduct(ctx context.Context, product entity.Inventories) error {
	args := w.Called(ctx, product)
	return args.Error(0)
}
