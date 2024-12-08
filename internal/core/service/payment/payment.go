package payment

import (
	"context"
	"log"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/config"
	"github.com/swclabs/swipex/pkg/gen/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ = app.Service(New)

func New() *Payment {
	conn, err := grpc.NewClient(config.PaymentService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	return &Payment{
		payment.NewVNPayClient(conn),
	}
}

type Payment struct {
	client payment.VNPayClient
}

// CheckStatus implements payment.VNPayClient.
func (p *Payment) CheckStatus(ctx context.Context, in *payment.StatusRequest, opts ...grpc.CallOption) (*payment.StatusResponse, error) {
	return p.client.CheckStatus(ctx, in, opts...)
}

// ProcessPayment implements payment.VNPayClient.
func (p *Payment) ProcessPayment(ctx context.Context, in *payment.PaymentRequest, opts ...grpc.CallOption) (*payment.PaymentResponse, error) {
	return p.client.ProcessPayment(ctx, in, opts...)
}

// ProcessPaymentReturn implements payment.VNPayClient.
func (p *Payment) ProcessPaymentReturn(ctx context.Context, in *payment.PaymentReturnRequest, opts ...grpc.CallOption) (*payment.PaymentReturnResponse, error) {
	return p.client.ProcessPaymentReturn(ctx, in, opts...)
}
