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
	productRepo "swclabs/swix/internal/core/repository/products"
	"swclabs/swix/internal/core/repository/specifications"
	"swclabs/swix/internal/core/service/products"
	"swclabs/swix/internal/webapi/controller"
	"swclabs/swix/pkg/lib/logger"
	"testing"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func TestProductDetails(t *testing.T) {
	var (
		productSpecs = dtos.ProductSpecs{
			Screen:  "6.1 inch",
			Display: "Super Retina XDR display",
		}
		inventorySpec = dtos.InvSpecification{
			RAM: "8GB",
			SSD: "256GB",
		}
		bProductSpecs, _ = json.Marshal(productSpecs)
		bInvSpec, _      = json.Marshal(inventorySpec)
		inventory        inventories.Mock
		product          productRepo.Mock
		specs            specifications.Mock
		service          = products.ProductService{
			Inventory: &inventory,
			Products:  &product,
			Specs:     &specs,
		}
		controller = controller.Products{
			Services: &service,
		}
	)

	for i := 1; i <= 4; i++ {
		specs.On("GetByInventoryID", context.Background(), int64(i)).Return([]entity.Specifications{
			{
				ID:          1,
				InventoryID: int64(i),
				Content:     string(bInvSpec),
			},
			{
				ID:          2,
				InventoryID: int64(i),
				Content:     string(bInvSpec),
			},
			{
				ID:          3,
				InventoryID: int64(i),
				Content:     string(bInvSpec),
			},
		}, nil)
	}

	// sInventorySpec, _ := json.Marshal(inventorySpec)
	inventory.On("GetByProductID", context.Background(), int64(1)).Return([]entity.Inventories{
		{
			ID:           1,
			ProductID:    1,
			Available:    1000,
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "Black Titanium",
			ColorImg:     "https://example.com/black-titanium.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg,https://example.com/iphone-12-3.jpg",
		},
		{
			ID:           2,
			ProductID:    1,
			Available:    1000,
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "White Ceramic",
			ColorImg:     "https://example.com/white-ceramic.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
		},
		{
			ID:           3,
			ProductID:    1,
			Available:    1000,
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "Blue Titanium",
			ColorImg:     "https://example.com/blue-titanium.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
		},
		{
			ID:           4,
			ProductID:    1,
			Available:    1000,
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "Red Titanium",
			ColorImg:     "https://example.com/red-titanium.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
		},
	}, nil)
	product.On("GetByID", context.Background(), int64(1)).Return(&entity.Products{
		Name:   "iPhone 12",
		Image:  "/img/shop/iphone-15-pro/unselect/iphone-15-pro-model-unselect-gallery-1-202309.jpg,/img/shop/iphone-15-pro/unselect/iphone-15-pro-model-unselect-gallery-2-202309.jpg,/img/shop/iphone-15-pro/iphone-15-pro-finish-select.jpg",
		Price:  "17.000.000 - 18.000.000",
		Specs:  string(bProductSpecs),
		Status: "active",
	}, nil)

	e.GET("/products/details", controller.GetProductDetails)
	req := httptest.NewRequest(http.MethodGet, "/products/details?id=1", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body dtos.ProductDetail
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fail()
	}

	file, err := os.Create("./products_detail_out.json")
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
