package inventories

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"

	"github.com/stretchr/testify/mock"
)

var _ IInventories = (*Mock)(nil)

// Mock represents a mock for IInventoryRepository.
type Mock struct {
	mock.Mock
}

// GetByColor implements IInventories.
func (w *Mock) GetByColor(ctx context.Context, productID int64, color string) ([]entity.Inventories, error) {
	args := w.Called(ctx, productID, color)
	return args.Get(0).([]entity.Inventories), args.Error(1)
}

// GetColor implements IInventories.
func (w *Mock) GetColor(ctx context.Context, productID int64) ([]model.ColorItem, error) {
	args := w.Called(ctx, productID)
	return args.Get(0).([]model.ColorItem), args.Error(1)
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
func (w *Mock) GetByProductID(ctx context.Context, ID int64) ([]entity.Inventories, error) {
	args := w.Called(ctx, ID)
	return args.Get(0).([]entity.Inventories), args.Error(1)
}

// GetByID implements IInventoryRepository.
func (w *Mock) GetByID(ctx context.Context, ID int64) (*entity.Inventories, error) {
	args := w.Called(ctx, ID)
	return args.Get(0).(*entity.Inventories), args.Error(1)
}

// InsertProduct implements IInventoryRepository.
func (w *Mock) InsertProduct(ctx context.Context, product entity.Inventories) (int64, error) {
	args := w.Called(ctx, product)
	return int64(args.Int(0)), args.Error(1)
}
