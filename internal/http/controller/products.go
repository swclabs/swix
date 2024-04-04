package controller

import (
	"net/http"
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/service"

	"github.com/labstack/echo/v4"
)

type IProducts interface {
	GetNewsletter(c echo.Context) error
	GetCategories(c echo.Context) error
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
// @Router /newsletter [GET]
func (p *Products) GetNewsletter(c echo.Context) error {
	return nil
}

// GetCategories
// @Description Get categories
// @Tags products
// @Accept json
// @Produce json
// @Param limit query number true "limit number"
// @Success 200 {object} domain.CategoriesList
// @Router /categories [GET]
func (p *Products) GetCategories(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "required 'limit' query params",
		})
	}

	resp, err := p.service.GetCategoriesLimit(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.CategoriesList{
		Data: resp,
	})
}
