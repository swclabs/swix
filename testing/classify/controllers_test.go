package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repository/suppliers"
	"swclabs/swix/internal/core/service/classify"
	"swclabs/swix/internal/webapi/controller"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var e = echo.New()

func TestGetSuppliers(t *testing.T) {
	// repository layers
	repos := suppliers.Mock{}
	repos.On("GetLimit", context.Background(), 10).Return([]entity.Suppliers{
		{
			ID:    "1",
			Name:  "apple",
			Email: "apple@example.com",
		},
	}, nil)

	// business logic layers
	services := classify.Classify{
		Supplier: &repos,
	}

	// presenter layers
	controllers := controller.Classify{
		Service: &services,
	}

	e.GET("/suppliers", controllers.GetSupplier)

	req := httptest.NewRequest(http.MethodGet, "/suppliers?limit=10", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"body\":[{\"id\":\"1\",\"name\":\"apple\",\"email\":\"apple@example.com\"}]}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
