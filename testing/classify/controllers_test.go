package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	classifyContainer "github.com/swclabs/swipex/internal/apis/container/classify"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/repos/suppliers"
	classifyService "github.com/swclabs/swipex/internal/core/service/classify"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var e = echo.New()

func TestGetSuppliers(t *testing.T) {
	// repos layers
	repos := suppliers.Mock{}
	repos.On("GetLimit", context.Background(), 10).Return([]entity.Supplier{
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

	expected := "{\"body\":[{\"id\":1,\"name\":\"apple\",\"email\":\"apple@example.com\"}]}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}
