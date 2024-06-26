package warehouse

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IWarehouseRepository = (*Mock)(nil)

// GetById implements IWarehouseRepository.
func (w *Mock) GetById(ctx context.Context, warehouseId int64) (*domain.Warehouse, error) {
	panic("unimplemented")
}

// GetProducts implements domain.IWarehouseRepository.
func (w *Mock) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Warehouse, error) {
	args := w.Called(ctx, productID, ram, ssd, color)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

// InsertProduct implements domain.IWarehouseRepository.
func (w *Mock) InsertProduct(ctx context.Context, product domain.WarehouseStruct) error {
	args := w.Called(ctx, product)
	return args.Error(0)
}
