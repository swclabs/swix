package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

var _ IInventoryRepository = (*Mock)(nil)

// Mock represents a mock for IInventoryRepository.
type Mock struct {
	mock.Mock
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
func (w *Mock) GetLimit(_ context.Context, _ int, _ int) ([]domain.Inventories, error) {
	panic("unimplemented")
}

// GetByProductID implements IInventoryRepository.
func (w *Mock) GetByProductID(_ context.Context, _ int64) ([]domain.Inventories, error) {
	//TODO implement me
	panic("implement me")
}

// GetByID implements IInventoryRepository.
func (w *Mock) GetByID(_ context.Context, _ int64) (*domain.Inventories, error) {
	panic("unimplemented")
}

// FindDevice implements IInventoryRepository.
func (w *Mock) FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error) {
	args := w.Called(ctx, deviceSpecs)
	return args.Get(0).(*domain.Inventories), args.Error(1)
}

// InsertProduct implements IInventoryRepository.
func (w *Mock) InsertProduct(ctx context.Context, product domain.InventoryStruct) error {
	args := w.Called(ctx, product)
	return args.Error(0)
}
