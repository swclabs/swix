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

func TestGetNewsletters(t *testing.T) {
	// repository layers
	repos := repository.NewsletterMock{}
	repos.On("Get", context.Background(), 10).Return([]domain.Newsletters{
		{
			Id: "1",
			Newsletter: domain.Newsletter{
				Title:       "KHOAN THANH TOAN HANG THANG THAP",
				SubTitle:    "Tra gop hang thang voi Momo",
				Description: "Tra dan, thoi han toi 24 thang va chi tra truoc 20%.",
				Image:       "/img/store/iphone2.jpg",
				TextColor:   "text-black",
				Type:        "product-page",
			},
		},
	}, nil)

	// bussiness logic layers
	service := service.Posts{
		Newsletter: &repos,
	}

	// presenter layers
	controllers := controller.Posts{
		Services: &service,
	}

	e.GET("/newsletters", controllers.GetNewsletter)

	req := httptest.NewRequest(http.MethodGet, "/newsletters?limit=10", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"data\":[{\"id\":\"1\",\"type\":\"product-page\",\"title\":\"KHOAN THANH TOAN HANG THANG THAP\",\"subtitle\":\"Tra gop hang thang voi Momo\",\"description\":\"Tra dan, thoi han toi 24 thang va chi tra truoc 20%.\",\"image\":\"/img/store/iphone2.jpg\",\"textcolor\":\"text-black\"}]}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}

func TestGetSuppliers(t *testing.T) {
	// repository layers
	repos := repository.SuppliersMock{}
	repos.On("GetLimit", context.Background(), 10).Return([]domain.Suppliers{
		{
			Id:          "1",
			Name:        "apple",
			PhoneNumber: "",
			Email:       "apple@example.com",
		},
	}, nil)

	// bussiness logic layers
	service := service.ProductService{
		Suppliers: &repos,
	}

	// presenter layers
	controllers := controller.Products{
		Services: &service,
	}

	e.GET("/suppliers", controllers.GetSupplier)

	req := httptest.NewRequest(http.MethodGet, "/suppliers?limit=10", nil)
	rr := httptest.NewRecorder()

	e.ServeHTTP(rr, req)

	expected := "{\"data\":[{\"id\":\"1\",\"name\":\"apple\",\"phone_number\":\"\",\"email\":\"apple@example.com\"}]}\n"
	assert.Equal(t, expected, rr.Body.String(), "response body should match expected")
}

func TestUploadNewsletters(t *testing.T) {

}
