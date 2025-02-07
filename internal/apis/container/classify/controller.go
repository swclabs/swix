package classify

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/service/classify"
	"github.com/swclabs/swipex/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewController)

// NewController creates a new Classify object
func NewController(service classify.IClassify) IController {
	return &Controller{
		Service: service,
	}
}

// IController interface for classify controller
type IController interface {
	GetSupplier(c echo.Context) error
	InsertSupplier(c echo.Context) error

	GetCategories(c echo.Context) error
	InsertCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
}

// Controller struct implementation of IClassify
type Controller struct {
	Service classify.IClassify
}

// GetCategories .
// @Description get categories
// @Tags categories
// @Accept json
// @Produce json
// @Param limit query number true "limit number"
// @Success 200 {object} dtos.Slices[entity.Category]
// @Router /categories [GET]
func (classify *Controller) GetCategories(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "required 'limit' query params",
		})
	}

	resp, err := classify.Service.GetCategoriesLimit(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.Slices[entity.Category]{
		Body: resp,
	})
}

// GetSupplier .
// @Description get suppliers information
// @Tags suppliers
// @Accept json
// @Produce json
// @Param limit query int true "limit number of suppliers"
// @Success 200 {object} dtos.Slices[entity.Supplier]
// @Router /suppliers [GET]
func (classify *Controller) GetSupplier(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	_supp, err := classify.Service.GetSuppliersLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.Slices[entity.Supplier]{
		Body: _supp,
	})
}

// InsertCategory .
// @Description insert new category
// @Tags categories
// @Accept json
// @Produce json
// @Param login body entity.Category true "Categories Request"
// @Success 201 {object} dtos.OK
// @Router /categories [POST]
func (classify *Controller) InsertCategory(c echo.Context) error {
	var request entity.Category
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	if err := classify.Service.CreateCategory(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: fmt.Sprintf("category data invalid, %v", err),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "category has been created",
	})
}

// InsertSupplier .
// @Description insert new suppliers information
// @Tags suppliers
// @Accept json
// @Produce json
// @Param Supplier body dtos.Supplier true "Suppliers Request"
// @Success 201 {object} dtos.OK
// @Router /suppliers [POST]
func (classify *Controller) InsertSupplier(c echo.Context) error {
	var req dtos.Supplier
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if validate := valid.Validate(&req); validate != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: validate.Error(),
		})
	}
	if err := classify.Service.CreateSuppliers(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "suppliers created successfully",
	})
}

// DeleteCategory .
// @Description delete category by ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "category ID"
// @Success 200 {object} dtos.OK
// @Router /categories/{id} [DELETE]
func (classify *Controller) DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "invalid category ID",
		})
	}
	if err := classify.Service.DelCategoryByID(c.Request().Context(), int64(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "category has been deleted",
	})
}

// UpdateCategory .
// @Description update category information
// @Tags categories
// @Accept json
// @Category json
// @Param id path int true "category ID"
// @Param category body dtos.UpdateCategories true "Category Request"
// @Success 200 {object} dtos.OK
// @Router /categories/{id} [PUT]
func (classify *Controller) UpdateCategory(c echo.Context) error {
	var payload dtos.UpdateCategories
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&payload); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}

	if err := classify.Service.UpdateCategoryInfo(c.Request().Context(), dtos.UpdateCategories{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "your product has been updated successfully",
	})
}
