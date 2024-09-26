// Package purchase implements handler of worker
package purchase

import (
	"context"
	"encoding/json"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/purchase"

	"github.com/hibiken/asynq"
)

var _ = app.Controller(NewHandler)

// NewHandler creates a new Purchase object
func NewHandler(service purchase.IPurchase) *Handler {
	return &Handler{service: service}
}

// Handler is a struct for Handler.
type Handler struct {
	service purchase.IPurchase
}

// HandleAddToCart implements IPurchase.
func (p *Handler) AddToCart(_ context.Context, task *asynq.Task) error {
	var req dtos.CartInsertDTO
	if err := json.Unmarshal(task.Payload(), &req); err != nil {
		return err
	}
	return p.service.AddToCart(context.Background(), req)
}
