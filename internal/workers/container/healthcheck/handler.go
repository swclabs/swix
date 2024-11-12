// Package handler implements handler of worker
package healthcheck

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/service/healthcheck"
	"github.com/swclabs/swipex/pkg/lib/worker"
)

var _ = app.Controller(NewHandler)

// NewHandler creates a new base object
func NewHandler(_base healthcheck.IService) *Handler {
	return &Handler{
		handler: _base,
	}
}

// Handler struct define the base object
type Handler struct {
	handler healthcheck.IService // create handler for services
}

// HandleHealthCheck handle health check
func (base *Handler) WorkerCheckResult(c worker.Context) error {
	var num int64
	if err := json.Unmarshal(c.Payload(), &num); err != nil {
		return err
	}
	result, err := base.handler.WorkerCheckResult(context.Background(), num)
	if err != nil {
		return err
	}
	return c.Return([]byte(fmt.Sprintf("ID: %s", result)))
}
