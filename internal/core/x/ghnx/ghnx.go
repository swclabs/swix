package ghnx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"swclabs/swix/app"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/core/domain/xdto"
)

type pID struct {
	ProvinceID int `json:"province_id"`
}

type dID struct {
	DistrictID int `json:"district_id"`
}

type Ghnx struct {
	client *http.Client
}

type IGhnx interface {
	// CreateOrder(ctx context.Context, order xdto.CreateOrderDTO) (*xdto.OrderDTO, error)
	Ward(ctx context.Context, districtID int) (*xdto.WardDTO, error)
	District(ctx context.Context, provinceID int) (*xdto.DistrictDTO, error)
	Province(ctx context.Context) (*xdto.ProvinceDTO, error)
}

func call[T any](client *http.Client, method string, url string, bodyReq io.Reader) (*T, error) {
	req, err := http.NewRequest(method, url, bodyReq)
	if err != nil {
		return nil, fmt.Errorf("error when create request %v", err)
	}
	req.Header.Set("Token", config.DeliveryTokenAPI)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when handle request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error when read response: %v", err)
	}
	var types T
	if err := json.Unmarshal(body, &types); err != nil {
		return nil, fmt.Errorf("error when unmarshal body to struct: %v", err)
	}
	return &types, nil
}

var New = app.Service(func() IGhnx {
	return &Ghnx{
		client: &http.Client{},
	}
})

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
