package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/service/classify"
	"swclabs/swix/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

// IClassify interface for classify controller
type IClassify interface {
	GetSupplier(c echo.Context) error
	InsertSupplier(c echo.Context) error

	GetCategories(c echo.Context) error
	InsertCategory(c echo.Context) error
}

// NewClassify creates a new Classify object
func NewClassify(service classify.IClassify) IClassify {
	return &Classify{
		Service: service,
	}
}

// Classify struct implementation of IClassify
type Classify struct {
	Service classify.IClassify
}

// GetCategories .
// @Description get categories
// @Tags categories
// @Accept json
// @Produce json
// @Param limit query number true "limit number"
// @Success 200 {object} dtos.Slices[entity.Categories]
// @Router /categories [GET]
func (classify *Classify) GetCategories(c echo.Context) error {
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

	return c.JSON(http.StatusOK, dtos.Slices[entity.Categories]{
		Body: resp,
	})
}

// GetSupplier .
// @Description get suppliers information
// @Tags suppliers
// @Accept json
// @Produce json
// @Param limit query int true "limit number of suppliers"
// @Success 200 {object} dtos.Slices[entity.Suppliers]
// @Router /suppliers [GET]
func (classify *Classify) GetSupplier(c echo.Context) error {
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
	return c.JSON(http.StatusOK, dtos.Slices[entity.Suppliers]{
		Body: _supp,
	})
}

// InsertCategory .
// @Description insert new category
// @Tags categories
// @Accept json
// @Produce json
// @Param login body entity.Categories true "Categories Request"
// @Success 201 {object} dtos.OK
// @Router /categories [POST]
func (classify *Classify) InsertCategory(c echo.Context) error {
	var request entity.Categories
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
func (classify *Classify) InsertSupplier(c echo.Context) error {
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
