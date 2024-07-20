package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/internal/webapi/controller"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	productRepo "swclabs/swipecore/internal/core/repository/products"
)

var e = echo.New()

func TestGetProductAvailability(t *testing.T) {
	// repository layers
	specs, _ := json.Marshal(domain.InventorySpecsDetail{
		Color:      "black",
		RAM:        "16",
		Ssd:        "512",
		ColorImage: "",
		Image:      "",
	})
	inventoryRepos := inventories.Mock{}
	price, _ := decimal.NewFromString("10000")
	inventoryRepos.On("FindDevice", context.Background(),
		domain.InventoryDeviveSpecs{
			ProductID: "1",
			RAM:       "64",
			Ssd:       "512",
			Color:     "black",
		}).Return(&domain.Inventories{
		ProductID:    1,
		ID:           "1",
		Status:       "active",
		Available:    "100",
		Price:        price,
		Specs:        string(specs),
		CurrencyCode: "USD",
	}, nil)
	productRepos := productRepo.Mock{}
	productRepos.On("GetByID", context.Background(), int64(1)).Return(&domain.Products{
		Name: "iPhone 15 Pro Max",
	}, nil)

	// business logic layers
	services := products.ProductService{
		Inventory: &inventoryRepos,
		Products:  &productRepos,
	}
	// presenter layers
	controllers := controller.Products{
		Services: &services,
	}

	e.GET("/inventories", controllers.GetProductAvailability)

	req := httptest.NewRequest(http.MethodGet, "/inventories?pid=1&ram=64&ssd=512&color=black", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"id\":\"1\",\"product_name\":\"iPhone 15 Pro Max\",\"status\":\"active\",\"product_id\":\"1\",\"price\":\"10000\",\"available\":\"100\",\"currency_code\":\"USD\",\"specs\":{\"color\":\"black\",\"ram\":\"16\",\"ssd\":\"512\",\"color_image\":\"\",\"image\":\"\"}}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
