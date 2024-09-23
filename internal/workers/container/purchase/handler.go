// Package purchase implements handler of worker
package purchase

import (
	"context"
	"encoding/json"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/purchase"
	"swclabs/swix/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

var _ IHandler = (*Handler)(nil)
var _ = app.Controller(NewHandler)

// NewHandler creates a new Purchase object
func NewHandler(service purchase.IPurchase) IHandler {
	return &Handler{service: service}
}

// IHandler is an interface for Purchase.
type IHandler interface {
	HandleAddToCart() (string, worker.HandleFunc)
}

// Handler is a struct for Handler.
type Handler struct {
	purchase.Task
	service purchase.IPurchase
}

// HandleAddToCart implements IPurchase.
func (p *Handler) HandleAddToCart() (string, worker.HandleFunc) {
	return worker.GetTaskName(p.AddToCart),
		func(_ context.Context, task *asynq.Task) error {
			var req dtos.CartInsertDTO
			if err := json.Unmarshal(task.Payload(), &req); err != nil {
				return err
			}
			return p.service.AddToCart(context.Background(), req)
		}
}
