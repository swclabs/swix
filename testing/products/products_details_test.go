package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/repository/inventories"
	productRepo "swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/internal/webapi/controller"
	"swclabs/swipecore/pkg/lib/logger"
	"testing"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func TestProductDetails(t *testing.T) {
	var (
		productSpecs = dtos.ProductSpecs{
			Screen:  "6.1 inch",
			Display: "Super Retina XDR display",
			SSD:     []int{128, 256, 512},
			RAM:     []int{4, 8},
		}
		inventorySpec = dtos.InventorySpecification{
			RAM: "8GB",
			SSD: "256GB",
		}
		inventory inventories.Mock
		product   productRepo.Mock
		service   = products.ProductService{
			Inventory: &inventory,
			Products:  &product,
		}
		controller = controller.Products{
			Services: &service,
		}
	)

	sProductSpecs, _ := json.Marshal(productSpecs)
	sInventorySpec, _ := json.Marshal(inventorySpec)
	inventory.On("GetByProductID", context.Background(), int64(1)).Return([]entity.Inventories{
		{
			ID:           "1",
			ProductID:    1,
			Available:    "1000",
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "Black Titanium",
			ColorImg:     "https://example.com/black-titanium.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg,https://example.com/iphone-12-3.jpg",
			Specs:        string(sInventorySpec),
		},
		{
			ID:           "2",
			ProductID:    1,
			Available:    "1000",
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "White Ceramic",
			ColorImg:     "https://example.com/white-ceramic.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
			Specs:        string(sInventorySpec),
		},
		{
			ID:           "3",
			ProductID:    1,
			Available:    "1000",
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "Blue Titanium",
			ColorImg:     "https://example.com/blue-titanium.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
			Specs:        string(sInventorySpec),
		},
		{
			ID:           "4",
			ProductID:    1,
			Available:    "1000",
			Price:        decimal.NewFromInt(10000),
			CurrencyCode: "VND",
			Status:       "active",
			Color:        "Red Titanium",
			ColorImg:     "https://example.com/red-titanium.jpg",
			Image:        "https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg",
			Specs:        string(sInventorySpec),
		},
	}, nil)
	product.On("GetByID", context.Background(), int64(1)).Return(&entity.Products{
		Name:   "iPhone 12",
		Image:  "/img/shop/iphone-15-pro/unselect/iphone-15-pro-model-unselect-gallery-1-202309.jpg,/img/shop/iphone-15-pro/unselect/iphone-15-pro-model-unselect-gallery-2-202309.jpg,/img/shop/iphone-15-pro/iphone-15-pro-finish-select.jpg",
		Price:  "17.000.000 - 18.000.000",
		Spec:   string(sProductSpecs),
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
