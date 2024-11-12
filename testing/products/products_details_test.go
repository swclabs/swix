package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	productContainer "github.com/swclabs/swipex/internal/apis/container/products"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/internal/core/repos/categories"
	"github.com/swclabs/swipex/internal/core/repos/inventories"
	productRepo "github.com/swclabs/swipex/internal/core/repos/products"
	productService "github.com/swclabs/swipex/internal/core/service/products"
	"github.com/swclabs/swipex/pkg/lib/logger"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func TestProductDetails(t *testing.T) {
	var (
		productSpecs = dtos.ProductSpecs{
			Screen:  "6.1 inch",
			Display: "Super Retina XDR display",
		}
		bProductSpecs, _ = json.Marshal(productSpecs)
		inventory        inventories.Mock
		product          productRepo.Mock
		category         = categories.Mock{}
		service          = productService.Products{
			Inventory: &inventory,
			Products:  &product,
			Category:  &category,
		}
		controller = productContainer.NewController(&service)
	)

	specs, _ := json.Marshal(dtos.Specs{
		SSD: "256",
		RAM: "128",
	})

	category.On("GetByID", context.Background(), int64(1)).Return(&entity.Category{
		ID:          int64(1),
		Name:        "phone",
		Description: "phone",
	}, nil)

	inventory.On("GetColor", context.Background(), int64(1)).Return([]model.ColorItem{
		{
			Color: "Black Titanium",
		},
	}, nil)

	// sInventorySpec, _ := json.Marshal(inventorySpec)
	inventory.On("GetByColor", context.Background(), int64(1), "Black Titanium").Return([]entity.Inventory{
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
			Specs:        string(specs),
		},
	}, nil)

	product.On("GetByID", context.Background(), int64(1)).Return(&entity.Product{
		Name:       "iPhone 12",
		Image:      "/img/shop/iphone-15-pro/unselect/iphone-15-pro-model-unselect-gallery-1-202309.jpg,/img/shop/iphone-15-pro/unselect/iphone-15-pro-model-unselect-gallery-2-202309.jpg,/img/shop/iphone-15-pro/iphone-15-pro-finish-select.jpg",
		Price:      "17.000.000 - 18.000.000",
		Specs:      string(bProductSpecs),
		Status:     "active",
		CategoryID: 1,
	}, nil)

	var e = echo.New()
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
