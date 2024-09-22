package ghnx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/xdto"
	"swclabs/swix/pkg/lib/valid"
)

type IGhnx interface {
	Ward(ctx context.Context, districtID int) (*xdto.WardDTO, error)
	District(ctx context.Context, provinceID int) (*xdto.DistrictDTO, error)
	Province(ctx context.Context) (*xdto.ProvinceDTO, error)
	CreateOrder(ctx context.Context, shopID int, order xdto.CreateOrderDTO) (*xdto.OrderDTO, error)
	OrderInfo(ctx context.Context, orderCode string) (*xdto.OrderInfoDTO, error)
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
func (g *Ghnx) OrderInfo(ctx context.Context, OrderCode string) (*xdto.OrderInfoDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		OrderCode := orderCode{OrderCode: OrderCode}
		body, _ := json.Marshal(OrderCode)
		return call[xdto.OrderInfoDTO](g.client,
			"POST", "https://dev-online-gateway.ghn.vn/shiip/public-api/v2/shipping-order/detail",
			bytes.NewBuffer(body),
		)
	}
}

// CreateOrder implements IGhnx.
func (g *Ghnx) CreateOrder(ctx context.Context, shopID int, order xdto.CreateOrderDTO) (*xdto.OrderDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if err := valid.Validate(&order); err != nil {
			return nil, fmt.Errorf("error when validate order: %v", err)
		}
		body, _ := json.Marshal(order)
		return call[xdto.OrderDTO](g.client,
			"POST", "https://dev-online-gateway.ghn.vn/shiip/public-api/v2/shipping-order/create",
			bytes.NewBuffer(body),
			header{key: "ShopId", value: fmt.Sprintf("%d", shopID)},
		)
	}
}

// Ward implements IGhnx.
func (g *Ghnx) Ward(ctx context.Context, districtID int) (*xdto.WardDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		ID := dID{DistrictID: districtID}
		body, _ := json.Marshal(ID)
		return call[xdto.WardDTO](g.client,
			"GET", "https://dev-online-gateway.ghn.vn/shiip/public-api/master-data/ward?district_id",
			bytes.NewBuffer(body),
		)
	}
}

// District implements IGhnx.
func (g *Ghnx) District(ctx context.Context, provinceID int) (*xdto.DistrictDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		ID := pID{ProvinceID: provinceID}
		body, _ := json.Marshal(ID)
		return call[xdto.DistrictDTO](g.client,
			"GET", "https://dev-online-gateway.ghn.vn/shiip/public-api/master-data/district",
			bytes.NewBuffer(body),
		)
	}
}

// Province implements IGhnx.
func (g *Ghnx) Province(ctx context.Context) (*xdto.ProvinceDTO, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return call[xdto.ProvinceDTO](g.client,
			"GET", "https://dev-online-gateway.ghn.vn/shiip/public-api/master-data/province", nil)
	}
}
