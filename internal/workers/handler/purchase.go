// Package handler implements handler of worker
package handler

import (
	"context"
	"encoding/json"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/purchase"
	"swclabs/swix/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

var _ IPurchase = (*Purchase)(nil)
var _ = boot.Controller(NewPurchase)

// NewPurchase creates a new Purchase object
func NewPurchase(service purchase.IPurchase) IPurchase {
	return &Purchase{service: service}
}

// IPurchase is an interface for Purchase.
type IPurchase interface {
	HandleAddToCart() (string, worker.HandleFunc)
}

// Purchase is a struct for Purchase.
type Purchase struct {
	purchase.Task
	service purchase.IPurchase
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
