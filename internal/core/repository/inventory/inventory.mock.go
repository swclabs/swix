package inventory

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IInventoryRepository = (*Mock)(nil)

// GetById implements IInventoryRepository.
func (w *Mock) GetById(ctx context.Context, inventoryId int64) (*domain.Inventory, error) {
	panic("unimplemented")
}

// GetProducts implements IInventoryRepository.
func (w *Mock) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Inventory, error) {
	args := w.Called(ctx, productID, ram, ssd, color)
	return args.Get(0).(*domain.Inventory), args.Error(1)
}

// InsertProduct implements IInventoryRepository.
func (w *Mock) InsertProduct(ctx context.Context, product domain.InventoryStruct) error {
	args := w.Called(ctx, product)
	return args.Error(0)
}
