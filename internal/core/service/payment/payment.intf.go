package payment

import "context"

// IPaymentService : Module for Payment.
// Actor: Admin & Customer (User)
type IPaymentService interface {
	// GetPayments retrieves payment information.
	// ctx is the context to manage the request's lifecycle.
	GetPayments(ctx context.Context)
}
