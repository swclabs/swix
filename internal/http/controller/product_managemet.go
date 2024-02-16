package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/service"
	"github.com/swclabs/swipe-api/pkg/validator"
)

type IProductManagement interface {
	InsertCategory(c *gin.Context)
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
func (product *ProductManagement) InsertCategory(c *gin.Context) {
	var request domain.CategoriesRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	if _valid := validator.Validate(request); _valid != "" {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
		return
	}
	if err := product.services.InsertCategory(&domain.Categories{
		Name:        request.Name,
		Description: request.Description,
	}); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "category data invalid",
		})
		return
	}
	c.JSON(http.StatusCreated, domain.OK{
		Msg: "category has been created",
	})
}
