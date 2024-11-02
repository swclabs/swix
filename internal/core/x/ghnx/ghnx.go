package ghnx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/x/ghn"
	"swclabs/swipex/pkg/lib/valid"
)

type IGhnx interface {
	Ward(ctx context.Context, districtID int) (*ghn.WardDTO, error)
	District(ctx context.Context, provinceID int) (*ghn.DistrictDTO, error)
	Province(ctx context.Context) (*ghn.ProvinceDTO, error)
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
			"POST", "https://dev-online-gateway.ghn.vn/shiip/public-api/v2/shipping-order/detail",
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
			"POST", "https://dev-online-gateway.ghn.vn/shiip/public-api/v2/shipping-order/create",
			bytes.NewBuffer(body),
			header{key: "ShopId", value: fmt.Sprintf("%d", shopID)},
		)
	}
}

// Ward implements IGhnx.
func (g *Ghnx) Ward(ctx context.Context, districtID int) (*ghn.WardDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		ID := dID{DistrictID: districtID}
		body, _ := json.Marshal(ID)
		return call[ghn.WardDTO](g.client,
			"GET", "https://dev-online-gateway.ghn.vn/shiip/public-api/master-data/ward?district_id",
			bytes.NewBuffer(body),
		)
	}
}

// District implements IGhnx.
func (g *Ghnx) District(ctx context.Context, provinceID int) (*ghn.DistrictDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		ID := pID{ProvinceID: provinceID}
		body, _ := json.Marshal(ID)
		return call[ghn.DistrictDTO](g.client,
			"GET", "https://dev-online-gateway.ghn.vn/shiip/public-api/master-data/district",
			bytes.NewBuffer(body),
		)
	}
}

// Province implements IGhnx.
func (g *Ghnx) Province(ctx context.Context) (*ghn.ProvinceDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return call[ghn.ProvinceDTO](g.client,
			"GET", "https://dev-online-gateway.ghn.vn/shiip/public-api/master-data/province", nil)
	}
}
