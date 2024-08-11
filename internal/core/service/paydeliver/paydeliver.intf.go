// Package paydeliver define payment and delivery methods
package paydeliver

import (
	"context"
	"swclabs/swix/internal/core/domain/dtos"
)

// IPaymentDelivery : Module for Payment and Delivery.
type IPaymentDelivery interface {
	CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error
	GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error)
	CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error
	GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error)
}
