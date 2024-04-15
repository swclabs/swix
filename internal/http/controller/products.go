package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/service"

	"github.com/labstack/echo/v4"
)

type IProducts interface {
	GetNewsletter(c echo.Context) error
	GetCategories(c echo.Context) error
	GetProductLimit(c echo.Context) error
}

type Products struct {
	product domain.IProductService
}

func NewProducts() IProducts {
	return &Products{
		product: service.NewProductService(),
	}
}

// GetNewsletter
// @Description Get Product Newsletter
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of newsletter"
// @Success 200 {object} domain.NewsletterListResponse
// @Router /newsletter [GET]
func (p *Products) GetNewsletter(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	newsletter, err := p.product.GetNewsletter(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.NewsletterListResponse{
		Data: newsletter,
	})
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

	resp, err := p.product.GetCategoriesLimit(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.CategoriesList{
		Data: resp,
	})
}

// GetProductLimit
// @Description get product information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of products"
// @Success 200 {object} domain.ProductsListResponse
// @Router /products [GET]
func (p *Products) GetProductLimit(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	prd, err := p.product.GetProductsLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.ProductsListResponse{
		Data: prd,
	})
}
