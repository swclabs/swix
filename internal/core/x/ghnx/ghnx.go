package ghnx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/config"
	"swclabs/swipex/internal/core/domain/x/ghn"
	"swclabs/swipex/pkg/lib/valid"
)

type IGhnx interface {
	CreateOrder(ctx context.Context, shopID int, order ghn.CreateOrderDTO) (*ghn.OrderDTO, error)
	OrderInfo(ctx context.Context, orderCode string) (*ghn.OrderInfoDTO, error)
}

var New = app.Service(func() IGhnx {
	return &Ghnx{
		client: &http.Client{},
	}
})

type Ghnx struct {
	client *http.Client
}

// OrderInfo implements IGhnx.
func (g *Ghnx) OrderInfo(ctx context.Context, OrderCode string) (*ghn.OrderInfoDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		OrderCode := orderCode{OrderCode: OrderCode}
		body, _ := json.Marshal(OrderCode)
		return call[ghn.OrderInfoDTO](g.client,
			"POST", config.DeliveryAddressAPI,
			bytes.NewBuffer(body),
		)
	}
}

// CreateOrder implements IGhnx.
func (g *Ghnx) CreateOrder(ctx context.Context, shopID int, order ghn.CreateOrderDTO) (*ghn.OrderDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if err := valid.Validate(&order); err != nil {
			return nil, fmt.Errorf("error when validate order: %v", err)
		}

		body, _ := json.Marshal(order)

		return call[ghn.OrderDTO](g.client,
			"POST", config.DeliveryAddressAPI,
			bytes.NewBuffer(body),
			header{key: "ShopId", value: fmt.Sprintf("%d", shopID)},
		)
	}
}
