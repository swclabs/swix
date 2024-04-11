package controller

import (
	"net/http"
	"strconv"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/service"
	"swclabs/swipe-api/pkg/tools"
	"swclabs/swipe-api/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type IProductManagement interface {
	InsertCategory(c echo.Context) error
	UploadProductImage(c echo.Context) error
	UploadProduct(c echo.Context) error
	UploadNewsletter(c echo.Context) error
	GetSupplier(c echo.Context) error
	NewSuppliers(c echo.Context) error
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
	if err := product.services.InsertCategory(c.Request().Context(), &domain.Categories{
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
// @Accept multipart/form-data
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
	// get id params
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing param 'id' in yours request",
		})
	}
	// call services
	if err := product.services.UploadImage(id, file); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "update successfully",
	})
}

// UploadProduct
// @Description Create new product
// @Tags product_management
// @Accept multipart/form-data
// @Produce json
// @Param img formData file true "image of product"
// @Param product formData domain.ProductRequest true "Product Request"
// @Success 200 {object} domain.OK
// @Router /products [POST]
func (product *ProductManagement) UploadProduct(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	formData, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// bind json to structure
	var productReq domain.ProductRequest
	if err := mapstructure.Decode(utils.NxN2Nx1(formData.Value), &productReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// check validate struct
	if valid := tools.Validate(&productReq); valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: valid,
		})
	}
	// call services
	if err := product.services.UploadProduct(c.Request().Context(), file, productReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "upload product successfully",
	})
}

// UploadNewsletter
// @Description Create newsletter
// @Tags product_management
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param img formData file true "image of newsletter"
// @Param product formData domain.Newsletter true "Newsletter Request"
// @Success 200 {object} domain.OK
// @Router /newsletters [POST]
func (product *ProductManagement) UploadNewsletter(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	formData, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// bind json to structure
	var newsletter domain.Newsletter

	if err := mapstructure.Decode(utils.NxN2Nx1(formData.Value), &newsletter); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// check validate struct
	if valid := tools.Validate(&newsletter); valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: valid,
		})
	}
	// call services
	if err := product.services.UploadNewsletter(c.Request().Context(), newsletter, file); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "upload newsletter successfully",
	})
}

// GetSupplier
// @Description get suppliers information
// @Tags product_management
// @Accept json
// @Produce json
// @Param limit query int true "limit number of suppliers"
// @Success 200 {object} domain.SuppliersListResponse
// @Router /suppliers [GET]
func (product *ProductManagement) GetSupplier(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	_supp, err := product.services.GetSuppliersLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.SuppliersListResponse{
		Data: _supp,
	})
}

// NewSuppliers
// @Description insert new suppliers information
// @Tags product_management
// @Accept json
// @Produce json
// @Param SuppliersRequest body domain.SuppliersRequest true "Suppliers Request"
// @Success 200 {object} domain.OK
// @Router /suppliers [POST]
func (product *ProductManagement) NewSuppliers(c echo.Context) error {
	var req domain.SuppliersRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if valid := tools.Validate(req); valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: valid,
		})
	}
	if err := product.services.InsertSuppliers(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.OK{
		Msg: "suppliers created successfully",
	})
}
