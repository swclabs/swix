// Package purchase implements handler of worker
package purchase

import (
	"context"
	"encoding/json"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/service/purchase"
	"swclabs/swipex/pkg/lib/worker"
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
func (p *Handler) AddToCart(c worker.Context) error {
	var req dtos.CartInsertDTO
	if err := json.Unmarshal(c.Payload(), &req); err != nil {
		return err
	}
	return p.service.AddToCart(context.Background(), req)
}
