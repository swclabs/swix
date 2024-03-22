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
// @Tags category
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
