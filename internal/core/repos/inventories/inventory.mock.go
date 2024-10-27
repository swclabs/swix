package inventories

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"

	"github.com/stretchr/testify/mock"
)

var _ IInventories = (*Mock)(nil)

// Mock represents a mock for IInventoryRepository.
type Mock struct {
	mock.Mock
}

// UploadColorImage implements IInventories.
func (w *Mock) UploadColorImage(ctx context.Context, ID int, url string) error {
	panic("unimplemented")
}

// GetByColor implements IInventories.
func (w *Mock) GetByColor(ctx context.Context, productID int64, color string) ([]entity.Inventory, error) {
	args := w.Called(ctx, productID, color)
	return args.Get(0).([]entity.Inventory), args.Error(1)
}

// GetColor implements IInventories.
func (w *Mock) GetColor(ctx context.Context, productID int64) ([]model.ColorItem, error) {
	args := w.Called(ctx, productID)
	return args.Get(0).([]model.ColorItem), args.Error(1)
}

// Update implements IInventoryRepository.
func (w *Mock) Update(_ context.Context, _ entity.Inventory) error {
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
func (w *Mock) GetLimit(_ context.Context, _ int, _ int) ([]entity.Inventory, error) {
	panic("unimplemented")
}

// GetByProductID implements IInventoryRepository.
func (w *Mock) GetByProductID(ctx context.Context, ID int64) ([]entity.Inventory, error) {
	args := w.Called(ctx, ID)
	return args.Get(0).([]entity.Inventory), args.Error(1)
}

// GetByID implements IInventoryRepository.
func (w *Mock) GetByID(ctx context.Context, ID int64) (*entity.Inventory, error) {
	args := w.Called(ctx, ID)
	return args.Get(0).(*entity.Inventory), args.Error(1)
}

// InsertProduct implements IInventoryRepository.
func (w *Mock) InsertProduct(ctx context.Context, product entity.Inventory) (int64, error) {
	args := w.Called(ctx, product)
	return int64(args.Int(0)), args.Error(1)
}
