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

func TestGetSuppliers(t *testing.T) {
	// repository layers
	repos := repository.SuppliersMock{}
	repos.On("GetLimit", context.Background(), 10).Return([]domain.Suppliers{
		{
			Id:    "1",
			Name:  "apple",
			Email: "apple@example.com",
		},
	}, nil)

	// business logic layers
	services := service.ProductService{
		Suppliers: &repos,
	}

	// presenter layers
	controllers := controller.Products{
		Services: &services,
	}

	e.GET("/suppliers", controllers.GetSupplier)

	req := httptest.NewRequest(http.MethodGet, "/suppliers?limit=10", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"data\":[{\"id\":\"1\",\"name\":\"apple\",\"email\":\"apple@example.com\"}]}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
