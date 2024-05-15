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
	repos.On("GetProducts", context.Background(), "1", "64", "512", "black").Return(&domain.WarehouseRes{
		Id: "1",
		WarehouseReq: domain.WarehouseReq{ProductID: "1",
			Model:     "iPhone 15 Pro Max",
			Available: "100",
			Price:     "$1000",
			Specs: domain.SpecsDetail{
				Color:      "black",
				Ram:        "16",
				Ssd:        "512",
				ColorImage: "",
				Image:      "",
			}},
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

	req := httptest.NewRequest(http.MethodGet, "/warehouse?pid=1&ram=64&ssd=512&color=black", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"id\":\"1\",\"product_id\":\"1\",\"price\":\"$1000\",\"model\":\"iPhone 15 Pro Max\",\"specs\":{\"color\":\"black\",\"ram\":\"16\",\"ssd\":\"512\",\"color_image\":\"\",\"image\":\"\"},\"available\":\"100\"}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
