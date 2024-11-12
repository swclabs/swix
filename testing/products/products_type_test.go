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
	"github.com/swclabs/swipex/internal/core/domain/enum"
	"github.com/swclabs/swipex/internal/core/domain/model"
	productRepo "github.com/swclabs/swipex/internal/core/repos/products"
	productService "github.com/swclabs/swipex/internal/core/service/products"
	"github.com/swclabs/swipex/pkg/lib/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func TestProductType(t *testing.T) {
	var (
		product    productRepo.Mock
		service    = productService.Products{Products: &product}
		controller = productContainer.NewController(&service)
		specs, _   = json.Marshal(dtos.ProductSpecs{})
		products   = []model.ProductXCategory{
			{
				ID:           1,
				Name:         "iPhone 12",
				Description:  "iPhone 12",
				Price:        "1.000.000 - 2.000.000",
				Image:        "https://example.com/iphone-12.jpg",
				CategoryName: enum.Phone.String(),
				Specs:        string(specs),
			},
			{
				ID:           2,
				Name:         "iPhone 12 Pro",
				Description:  "iPhone 12 Pro",
				Price:        "2.000.000 - 3.000.000",
				Image:        "https://example.com/iphone-12-pro.jpg",
				CategoryName: enum.Phone.String(),
				Specs:        string(specs),
			},
			{
				ID:           3,
				Name:         "iPhone 12 Mini",
				Description:  "iPhone 12 Mini",
				Price:        "1.000.000 - 2.000.000",
				Image:        "https://example.com/iphone-12-mini.jpg",
				CategoryName: enum.Phone.String(),
				Specs:        string(specs),
			},
		}
	)

	product.On("GetByCategory", context.Background(), enum.Phone, 0).Return(products, nil)

	var e = echo.New()
	e.GET("/products/:type", controller.GetProductByType)
	req := httptest.NewRequest(http.MethodGet, "/products/phone", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body []dtos.ProductDTO
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("./products_type_out.json")
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

func TestProductTypeAccessory(t *testing.T) {
	var (
		product    productRepo.Mock
		service    = productService.Products{Products: &product}
		controller = productContainer.NewController(&service)
		specs, _   = json.Marshal(dtos.ProductSpecs{
			SSD: []int{},
			RAM: []int{},
		})
		products = []model.ProductXCategory{
			{
				ID:           1,
				Name:         "Apple iPhone Adapter",
				Description:  "Apple iPhone Adapter 20W",
				Price:        "500.000",
				Image:        "https://example.com/apple-iphone-adapter.jpg",
				Specs:        string(specs),
				CategoryName: enum.Accessories.String(),
			},
			{
				ID:           2,
				Name:         "Apple iPhone Case",
				Description:  "Apple iPhone Case",
				Price:        "500.000",
				Image:        "https://example.com/apple-iphone-case.jpg",
				Specs:        string(specs),
				CategoryName: enum.Accessories.String(),
			},
			{
				ID:           3,
				Name:         "Apple iPhone Screen Protector",
				Description:  "Apple iPhone Screen Protector",
				Price:        "500.000",
				Image:        "https://example.com/apple-iphone-screen-protector.jpg",
				Specs:        string(specs),
				CategoryName: enum.Accessories.String(),
			},
		}
	)

	product.On("GetByCategory", context.Background(), enum.Accessories, 0).Return(products, nil)

	var e = echo.New()
	e.GET("/products/:type", controller.GetProductByType)
	req := httptest.NewRequest(http.MethodGet, "/products/accessories", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	responseBody := rr.Body.Bytes()
	var body []dtos.ProductDTO
	if err := json.Unmarshal(responseBody, &body); err != nil {
		t.Fatal(rr.Body.String())
	}

	file, err := os.Create("./products_type_accessory_out.json")
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
