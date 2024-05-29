package delivery

import (
	"context"
)

// IDeliveryService : Module for Delivery.
// Actor: Admin & Customer (User)
type IDeliveryService interface {
	// GetDeliveryInfo retrieves delivery information.
	// ctx is the context to manage the request's lifecycle.
	GetDeliveryInfo(ctx context.Context)
}
