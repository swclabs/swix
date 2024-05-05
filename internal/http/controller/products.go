package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/pkg/tools/valid"
	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

type IProducts interface {
	GetCategories(c echo.Context) error
	GetProductLimit(c echo.Context) error
	GetSupplier(c echo.Context) error
	InsertCategory(c echo.Context) error
	InsertSupplier(c echo.Context) error
	UploadProductImage(c echo.Context) error
	UploadProduct(c echo.Context) error
}

type Products struct {
	Services domain.IProductService
}

func NewProducts() IProducts {
	return &Products{
		Services: service.NewProductService(),
	}
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

	resp, err := p.Services.GetCategoriesLimit(c.Request().Context(), limit)
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
	prd, err := p.Services.GetProductsLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.ProductsListResponse{
		Data: prd,
	})
}

// InsertCategory
// @Description Insert new category
// @Tags products
// @Accept json
// @Produce json
// @Param login body domain.CategoriesRequest true "Categories Request"
// @Success 200 {object} domain.LoginResponse
// @Router /categories [POST]
func (p *Products) InsertCategory(c echo.Context) error {
	var request domain.CategoriesRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(request); _valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
	}
	if err := p.Services.InsertCategory(c.Request().Context(), &domain.Categories{
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
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param id query string true "id of product"
// @Param img formData file true "image of product"
// @Success 200 {object} domain.OK
// @Failure 400 {object} domain.Error
// @Router /products/img [POST]
func (p *Products) UploadProductImage(c echo.Context) error {
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// get id params
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	// call services
	if err := p.Services.UploadProductImage(c.Request().Context(), id, file); err != nil {
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
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param img formData file true "image of product"
// @Param product formData domain.ProductRequest true "Product Request"
// @Success 200 {object} domain.OK
// @Router /products [POST]
func (p *Products) UploadProduct(c echo.Context) error {
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
	if valid := valid.Validate(&productReq); valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: valid,
		})
	}
	// call services
	if err := p.Services.UploadProduct(c.Request().Context(), file, productReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "upload product successfully",
	})
}

// GetSupplier
// @Description get suppliers information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of suppliers"
// @Success 200 {object} domain.SuppliersListResponse
// @Router /suppliers [GET]
func (p *Products) GetSupplier(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	_supp, err := p.Services.GetSuppliersLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.SuppliersListResponse{
		Data: _supp,
	})
}

// InsertSupplier
// @Description insert new suppliers information
// @Tags products
// @Accept json
// @Produce json
// @Param SuppliersRequest body domain.SuppliersRequest true "Suppliers Request"
// @Success 200 {object} domain.OK
// @Router /suppliers [POST]
func (p *Products) InsertSupplier(c echo.Context) error {
	var req domain.SuppliersRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if valid := valid.Validate(req); valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: valid,
		})
	}
	if err := p.Services.InsertSuppliers(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.OK{
		Msg: "suppliers created successfully",
	})
}
