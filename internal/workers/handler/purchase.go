// Package handler implements handler of worker
package handler

import (
	"context"
	"encoding/json"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/purchase"
	"swclabs/swix/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

// IPurchase is an interface for Purchase.
type IPurchase interface {
	HandleAddToCart() (string, worker.HandleFunc)
}

var _ IPurchase = (*Purchase)(nil)

// NewPurchase creates a new Purchase object
func NewPurchaseConsume(service purchase.IPurchaseService) IPurchase {
	return &Purchase{service: service}
}

// Purchase is a struct for Purchase.
type Purchase struct {
	purchase.Task
	service purchase.IPurchaseService
}

// HandleAddToCart implements IPurchase.
func (p *Purchase) HandleAddToCart() (string, worker.HandleFunc) {
	return worker.GetTaskName(p.AddToCart),
		func(_ context.Context, task *asynq.Task) error {
			var req dtos.CartInsert
			if err := json.Unmarshal(task.Payload(), &req); err != nil {
				return err
			}
			return p.service.AddToCart(context.Background(), req)
		}
}
