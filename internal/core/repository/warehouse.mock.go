package repository

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type WarehouseMock struct {
	mock.Mock
}

var _ domain.IWarehouseRepository = (*WarehouseMock)(nil)

// GetProducts implements domain.IWarehouseRepository.
func (w *WarehouseMock) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Warehouse, error) {
	args := w.Called(ctx, productID, ram, ssd, color)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

// InsertProduct implements domain.IWarehouseRepository.
func (w *WarehouseMock) InsertProduct(ctx context.Context, product domain.WarehouseStructure) error {
	panic("unimplemented")
}
