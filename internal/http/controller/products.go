package controller

import (
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/service"

	"github.com/labstack/echo/v4"
)

type IProducts interface {
	GetNewsletter(c echo.Context) error
}

type Products struct {
	service domain.IProductService
}

func NewProducts() IProducts {
	return &Products{
		service: service.NewProductService(),
	}
}

// GetNewsletter
// @Description Get Product Newsletter
// @Tags products
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "Login"
// @Success 200 {object} domain.LoginResponse
// @Router /auth/login [POST]
func (p *Products) GetNewsletter(c echo.Context) error {
	return nil
}
