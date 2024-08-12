// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/workers/queue"
	"swclabs/swix/pkg/lib/worker"
)

var _ IPurchaseService = (*Task)(nil)

// UseTask use task for purchase service
func UseTask(service IPurchaseService) IPurchaseService {
	return &Task{
		service: service,
		worker:  worker.NewClient(config.RedisHost, config.RedisPort, config.RedisPassword),
	}
}

// Task struct for purchase service
type Task struct {
	worker  worker.IWorkerClient
	service IPurchaseService
}

// AddToCart implements IPurchaseService.
func (t *Task) AddToCart(_ context.Context, cart dtos.CartInsert) error {
	return t.worker.Exec(queue.CartQueue, worker.NewTask(
		worker.GetTaskName(t.AddToCart),
		cart,
	))
}

// CreateOrders implements IPurchaseService.
func (t *Task) CreateOrders(ctx context.Context, createOrder dtos.CreateOrderSchema) (string, error) {
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
