package paydeliver

import (
	"context"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repos/addresses"
	"swclabs/swix/internal/core/repos/deliveries"
	"time"
)

// New creates a new Classify object
func New(
	addr addresses.IAddress,
	del deliveries.IDeliveries,
) IPaymentDelivery {
	return &PaymentDelivery{
		Address:  addr,
		Delivery: del,
	}
}

// PaymentDelivery struct for classify service
type PaymentDelivery struct {
	Address  addresses.IAddress
	Delivery deliveries.IDeliveries
}

// CreateDelivery implements IPaymentDelivery.
func (p *PaymentDelivery) CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error {
	sendate, err := time.Parse(time.RFC3339, delivery.SentDate)
	if err != nil {
		sendate = time.Time{}
	}
	receivedate, err := time.Parse(time.RFC3339, delivery.ReceivedDate)
	if err != nil {
		receivedate = time.Time{}
	}
	return p.Delivery.Create(ctx, entity.Deliveries{
		UserID:       delivery.UserID,
		AddressID:    delivery.AddressID,
		Status:       delivery.Status,
		Method:       delivery.Method,
		Note:         delivery.Note,
		SentDate:     sendate,
		ReceivedDate: receivedate,
	})
}

// CreateDeliveryAddress implements IPaymentDelivery.
func (p *PaymentDelivery) CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error {
	return p.Address.Insert(ctx, entity.Addresses{
		UserID:   addr.UserID,
		Street:   addr.Street,
		City:     addr.City,
		Ward:     addr.Ward,
		District: addr.District,
	})
}

// GetDelivery implements IPaymentDelivery.
func (p *PaymentDelivery) GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error) {
	deliveries, err := p.Delivery.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var delivery = []dtos.Delivery{}
	for _, del := range deliveries {
		var (
			sentdate     string
			receiveddate string
		)
		if !del.SentDate.IsZero() {
			sentdate = del.SentDate.Format(time.RFC3339)
		}
		if !del.ReceivedDate.IsZero() {
			receiveddate = del.ReceivedDate.Format(time.RFC3339)
		}
		delivery = append(delivery, dtos.Delivery{
			ID:           del.ID,
			AddressID:    del.AddressID,
			UserID:       del.UserID,
			Status:       del.Status,
			Method:       del.Method,
			Note:         del.Note,
			SentDate:     sentdate,
			ReceivedDate: receiveddate,
		})
	}
	return delivery, nil
}

// GetDeliveryAddress implements IPaymentDelivery.
func (p *PaymentDelivery) GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error) {
	addrs, err := p.Address.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var addresses = []dtos.Address{}
	for _, addr := range addrs {
		addresses = append(addresses, dtos.Address{
			ID:       addr.ID,
			Street:   addr.Street,
			City:     addr.City,
			Ward:     addr.Ward,
			District: addr.District,
		})
	}
	return addresses, nil
}
