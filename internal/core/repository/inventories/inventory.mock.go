package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IInventoryRepository = (*Mock)(nil)

// GetLimit implements IInventoryRepository.
func (w *Mock) GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error) {
	panic("unimplemented")
}

func (w *Mock) GetByProductId(ctx context.Context, productId int64) ([]domain.Inventories, error) {
	//TODO implement me
	panic("implement me")
}

// GetById implements IInventoryRepository.
func (w *Mock) GetById(ctx context.Context, inventoryId int64) (*domain.Inventories, error) {
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
