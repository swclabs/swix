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
func (w *WarehouseMock) GetProducts(ctx context.Context, productID, ram, ssd, color string) (*domain.WarehouseRes, error) {
	args := w.Called(ctx, productID, ram, ssd, color)
	return args.Get(0).(*domain.WarehouseRes), args.Error(1)
}

// InsertProduct implements domain.IWarehouseRepository.
func (w *WarehouseMock) InsertProduct(ctx context.Context, product domain.WarehouseReq) error {
	panic("unimplemented")
}
