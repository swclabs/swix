// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/xdto"
	"swclabs/swix/internal/workers/queue"
	"swclabs/swix/pkg/lib/worker"
)

var _ IPurchase = (*Task)(nil)

// UseTask use task for purchase service
func UseTask(service IPurchase) IPurchase {
	return &Task{
		service: service,
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
	}
}

// Task struct for purchase service
type Task struct {
	worker  worker.IWorkerClient
	service IPurchase
}

// CreateDeliveryOrder implements IPurchase.
func (t *Task) CreateDeliveryOrder(ctx context.Context, shopID int, order xdto.CreateOrderDTO) (*xdto.OrderDTO, error) {
	panic("unimplemented")
}

// DeliveryOrderInfo implements IPurchase.
func (t *Task) DeliveryOrderInfo(ctx context.Context, orderCode string) (*xdto.OrderInfoDTO, error) {
	panic("unimplemented")
}

// AddressDistrict implements IPurchase.
func (t *Task) AddressDistrict(ctx context.Context, provinceID int) (*xdto.DistrictDTO, error) {
	panic("unimplemented")
}

// AddressProvince implements IPurchase.
func (t *Task) AddressProvince(ctx context.Context) (*xdto.ProvinceDTO, error) {
	panic("unimplemented")
}

// AddressWard implements IPurchase.
func (t *Task) AddressWard(ctx context.Context, districtID int) (*xdto.WardDTO, error) {
	panic("unimplemented")
}

// CreateDelivery implements IPurchase.
func (t *Task) CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error {
	panic("unimplemented")
}

// CreateDeliveryAddress implements IPurchase.
func (t *Task) CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error {
	panic("unimplemented")
}

// GetDelivery implements IPurchase.
func (t *Task) GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error) {
	panic("unimplemented")
}

// GetDeliveryAddress implements IPurchase.
func (t *Task) GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error) {
	panic("unimplemented")
}

// AddToCart implements IPurchaseService.
func (t *Task) AddToCart(ctx context.Context, cart dtos.CartInsertDTO) error {
	return t.worker.Exec(ctx, queue.CartQueue, worker.NewTask(
		worker.GetTaskName(t.AddToCart),
		cart,
	))
}

// CreateOrders implements IPurchaseService.
func (t *Task) CreateOrders(ctx context.Context, createOrder dtos.CreateOrderDTO) (string, error) {
	return t.service.CreateOrders(ctx, createOrder)
}

// DeleteItemFromCart implements IPurchaseService.
func (t *Task) DeleteItemFromCart(ctx context.Context, cartID int64) error {
	return t.service.DeleteItemFromCart(ctx, cartID)
}

// GetCart implements IPurchaseService.
func (t *Task) GetCart(ctx context.Context, userID int64, limit int) (*dtos.CartSlices, error) {
	return t.service.GetCart(ctx, userID, limit)
}

// GetOrdersByUserID implements IPurchaseService.
func (t *Task) GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderSchema, error) {
	return t.service.GetOrdersByUserID(ctx, userID, limit)
}
