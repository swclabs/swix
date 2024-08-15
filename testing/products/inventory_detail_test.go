package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repository/inventories"
	"swclabs/swix/internal/core/repository/specifications"
	"swclabs/swix/internal/core/service/products"
	"swclabs/swix/internal/webapi/controller"
	"swclabs/swix/pkg/lib/logger"
	"testing"

	productRepo "swclabs/swix/internal/core/repository/products"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var e = echo.New()

func TestGetInventory(t *testing.T) {
	var (
		spec = dtos.InvStorage{
			ID:  1,
			RAM: "8GB",
			SSD: "256GB",
		}
		bSpec, _  = json.Marshal(spec)
		inventory inventories.Mock
		product   productRepo.Mock
		specs     specifications.Mock
		service   = products.ProductService{
			Inventory: &inventory,
			Products:  &product,
			Specs:     &specs,
		}
		controller = controller.Products{
			Services: &service,
		}
	)

	specs.On("GetByInventoryID", context.Background(), int64(1)).Return([]entity.Specifications{
		{
			ID:          1,
			InventoryID: 1,
			Content:     string(bSpec),
		},
	}, nil)

	inventory.On("GetByID", context.Background(), int64(1)).Return(&entity.Inventories{
		ID:           1,
		ProductID:    1,
		Available:    1000,
		Price:        decimal.NewFromInt(10000),
		CurrencyCode: "VND",
		Status:       "active",
		Color:        "Black Titanium",
		ColorImg:     "https://example.com/black-titanium.jpg",
		Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
	}, nil)

	product.On("GetByID", context.Background(), int64(1)).Return(&entity.Products{
		Name: "iPhone 12",
	}, nil)

	e.GET("/inventories/details", controller.GetInvDetails)
	req := httptest.NewRequest(http.MethodGet, "/inventories/details?id=1", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body dtos.Inventory[dtos.InvStorage]
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fail()
	}

	file, err := os.Create("./inventory_detail_out.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	logger := logger.Write(file)
	logger.Info("Response body", zap.Any("body", body), zap.Int("status", rr.Code))
}
