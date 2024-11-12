package ghnx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/swclabs/swipex/internal/config"
)

type pID struct {
	ProvinceID int `json:"province_id"`
}

type dID struct {
	DistrictID int `json:"district_id"`
}

type orderCode struct {
	OrderCode string `json:"order_code"`
}

type header struct {
	key   string
	value string
}

func call[T any](client *http.Client, method string, url string, bodyReq io.Reader, headers ...header) (*T, error) {
	req, err := http.NewRequest(method, url, bodyReq)
	if err != nil {
		return nil, fmt.Errorf("error when create request %v", err)
	}
	req.Header.Set("Token", config.DeliveryTokenAPI)
	req.Header.Set("Content-Type", "application/json")
	for _, header := range headers {
		req.Header.Set(header.key, header.value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error when handle request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error when read response: %v", err)
	}
	var types T
	if err := json.Unmarshal(body, &types); err != nil {
		return nil, fmt.Errorf("error when unmarshal body to struct: %v", err)
	}
	return &types, nil
}
