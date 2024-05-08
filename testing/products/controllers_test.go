package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/internal/http/controller"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var e = echo.New()

func TestGetProductAvailability(t *testing.T) {
	// repository layers
	repos := repository.WarehouseMock{}
	repos.On("GetProducts", context.Background(), "1", "64", "512").Return(&domain.Warehouse{
		Id:        "1",
		ProductID: "1",
		Ram:       "64",
		Ssd:       "512",
		Model:     "iPhone 15 Pro Max",
		Available: "100",
		Price:     "$1000",
	}, nil)

	// bussiness logic layers
	service := service.ProductService{
		Warehouse: &repos,
	}

	// presenter layers
	controllers := controller.Products{
		Services: &service,
	}

	e.GET("/warehouse", controllers.GetProductAvailability)

	req := httptest.NewRequest(http.MethodGet, "/warehouse?pid=1&ram=64&ssd=512", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"id\":\"1\",\"product_id\":\"1\",\"price\":\"$1000\",\"ram\":\"64\",\"ssd\":\"512\",\"model\":\"iPhone 15 Pro Max\",\"available\":\"100\"}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
