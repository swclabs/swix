// Package purchase implements the purchase interface
package purchase

import (
	"context"

	"github.com/swclabs/swipex/internal/config"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/x/ghn"
	"github.com/swclabs/swipex/internal/workers/queue"
	"github.com/swclabs/swipex/pkg/lib/worker"
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

// DeleteCoupon implements IPurchase.
func (t *Task) DeleteCoupon(ctx context.Context, code string) error {
	return t.service.DeleteCoupon(ctx, code)
}

// UpdateOrderStatus implements IPurchase.
func (t *Task) UpdateOrderStatus(ctx context.Context, orderCode string, status string) error {
	return t.service.UpdateOrderStatus(ctx, orderCode, status)
}

func (t *Task) GetUsersByAdmin(ctx context.Context, limit int) ([]dtos.OrderInfo, error) {
	return t.service.GetUsersByAdmin(ctx, limit)
}

// AddressDistrict implements IPurchase.
func (t *Task) AddressDistrict(ctx context.Context, provinceID string) ([]entity.District, error) {
	return t.service.AddressDistrict(ctx, provinceID)
}

// AddressProvince implements IPurchase.
func (t *Task) AddressProvince(ctx context.Context) ([]entity.Province, error) {
	return t.service.AddressProvince(ctx)
}

// AddressWard implements IPurchase.
func (t *Task) AddressWard(ctx context.Context, districtID string) ([]entity.Commune, error) {
	return t.service.AddressWard(ctx, districtID)
}

// CreateCoupon implements IPurchase.
func (t *Task) CreateCoupon(ctx context.Context, coupon dtos.CreateCoupon) (code string, err error) {
	return t.service.CreateCoupon(ctx, coupon)
}

// GetCoupon implements IPurchase.
func (t *Task) GetCoupon(ctx context.Context) (coupons []dtos.Coupon, err error) {
	return t.service.GetCoupon(ctx)
}

// GetOrderByCode implements IPurchase.
func (t *Task) GetOrderByCode(ctx context.Context, orderCode string) (*dtos.OrderInfo, error) {
	return t.service.GetOrderByCode(ctx, orderCode)
}

// CreateOrderForm implements IPurchase.
func (t *Task) CreateOrderForm(ctx context.Context, order dtos.OrderForm) (string, error) {
	return t.service.CreateOrderForm(ctx, order)
}

// CreateDeliveryOrder implements IPurchase.
func (t *Task) CreateDeliveryOrder(ctx context.Context, shopID int, order ghn.CreateOrderDTO) (*ghn.OrderDTO, error) {
	return t.service.CreateDeliveryOrder(ctx, shopID, order)
}

// DeliveryOrderInfo implements IPurchase.
func (t *Task) DeliveryOrderInfo(ctx context.Context, orderCode string) (*ghn.OrderInfoDTO, error) {
	return t.service.DeliveryOrderInfo(ctx, orderCode)
}

// CreateDelivery implements IPurchase.
func (t *Task) CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error {
	return t.service.CreateDelivery(ctx, delivery)
}

// CreateDeliveryAddress implements IPurchase.
func (t *Task) CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error {
	return t.service.CreateDeliveryAddress(ctx, addr)
}

// GetDelivery implements IPurchase.
func (t *Task) GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error) {
	return t.service.GetDelivery(ctx, userID)
}

// GetDeliveryAddress implements IPurchase.
func (t *Task) GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error) {
	return t.service.GetDeliveryAddress(ctx, userID)
}

// AddToCart implements IPurchaseService.
func (t *Task) AddToCart(ctx context.Context, cart dtos.CartInsertDTO) error {
	return t.worker.Exec(ctx,
		queue.CartQueue,
		worker.NewTask("purchase.AddToCart", cart),
	)
}

// CreateOrders implements IPurchaseService.
func (t *Task) CreateOrders(ctx context.Context, userID int64, createOrder dtos.Order) (string, error) {
	return t.service.CreateOrders(ctx, userID, createOrder)
}

// DeleteItemFromCart implements IPurchaseService.
func (t *Task) DeleteItemFromCart(ctx context.Context, inventoryID int64, userID int64) error {
	return t.service.DeleteItemFromCart(ctx, inventoryID, userID)
}

// GetCart implements IPurchaseService.
func (t *Task) GetCart(ctx context.Context, userID int64, limit int) (*dtos.Carts, error) {
	return t.service.GetCart(ctx, userID, limit)
}

// GetOrdersByUserID implements IPurchaseService.
func (t *Task) GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderInfo, error) {
	return t.service.GetOrdersByUserID(ctx, userID, limit)
}
