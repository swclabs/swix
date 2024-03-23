package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipe-api/internal/core/domain"
	"github.com/swclabs/swipe-api/internal/core/service"
	"github.com/swclabs/swipe-api/pkg/tools"
)

type IProductManagement interface {
	InsertCategory(c echo.Context) error
	UploadProductImage(c echo.Context) error
}

type ProductManagement struct {
	services domain.IProductManagementService
}

func NewProductManagement() IProductManagement {
	return &ProductManagement{
		services: service.NewProductManagement(),
	}
}

// InsertCategory
// @Description Insert new category
// @Tags product_management
// @Accept json
// @Produce json
// @Param login body domain.CategoriesRequest true "Categories Request"
// @Success 200 {object} domain.LoginResponse
// @Router /categories [POST]
func (product *ProductManagement) InsertCategory(c echo.Context) error {
	var request domain.CategoriesRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if _valid := tools.Validate(request); _valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
	}
	if err := product.services.InsertCategory(&domain.Categories{
		Name:        request.Name,
		Description: request.Description,
	}); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "category data invalid",
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "category has been created",
	})
}

// UploadProductImage
// @Description Insert new product image
// @Tags product_management
// @Accept json
// @Produce json
// @Param id path string true "id of product"
// @Param img formData file true "image of product"
// @Success 200 {object} domain.OK
// @Failure 400 {object} domain.Error
// @Router /products/img/:id [POST]
func (product *ProductManagement) UploadProductImage(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing param 'id' in yours request",
		})
	}
	if err := product.services.UploadImage(id, file); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "update successfully",
	})
}
