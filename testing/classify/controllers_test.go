package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	classifyContainer "swclabs/swix/internal/apis/container/classify"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repos/suppliers"
	classifyService "swclabs/swix/internal/core/service/classify"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var e = echo.New()

func TestGetSuppliers(t *testing.T) {
	// repos layers
	repos := suppliers.Mock{}
	repos.On("GetLimit", context.Background(), 10).Return([]entity.Suppliers{
		{
			ID:    1,
			Name:  "apple",
			Email: "apple@example.com",
		},
	}, nil)

	// business logic layers
	services := classifyService.New(nil, &repos)

	// presenter layers
	controllers := classifyContainer.NewController(services)

	e.GET("/suppliers", controllers.GetSupplier)

	req := httptest.NewRequest(http.MethodGet, "/suppliers?limit=10", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"body\":[{\"id\":\"1\",\"name\":\"apple\",\"email\":\"apple@example.com\"}]}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
