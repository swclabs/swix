package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

type IProducts interface {
	GetCategories(c echo.Context) error
	GetProductLimit(c echo.Context) error
	GetSupplier(c echo.Context) error
	InsertCategory(c echo.Context) error
	InsertSupplier(c echo.Context) error
	UploadProductImage(c echo.Context) error
	UploadProduct(c echo.Context) error
	GetProductAvailability(c echo.Context) error
	AddToWarehouse(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type Products struct {
	Services products.IProductService
}

func NewProducts(services products.IProductService) IProducts {
	return &Products{
		Services: services,
	}
}

// GetProductAvailability
// @Description Get product availability in warehouse
// @Tags products
// @Accept json
// @Produce json
// @Param pid query number true "product id"
// @Param ram query number true "ram"
// @Param ssd query number true "ssd"
// @Param color query string true "color"
// @Success 200 {object} domain.WarehouseSchema
// @Router /warehouse [GET]
func (p *Products) GetProductAvailability(c echo.Context) error {
	pid := c.QueryParam("pid")
	if pid == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "required 'limit' query params",
		})
	}
	ram := c.QueryParam("ram")
	ssd := c.QueryParam("ssd")
	color := c.QueryParam("color")

	product, err := p.Services.GetProductsInWarehouse(c.Request().Context(), pid, ram, ssd, color)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// GetCategories
// @Description Get categories
// @Tags products
// @Accept json
// @Produce json
// @Param limit query number true "limit number"
// @Success 200 {object} domain.CategorySlices
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
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.CategorySlices{
		Data: resp,
	})
}

// GetProductLimit
// @Description get product information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of products"
// @Success 200 {object} domain.ProductsRes
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
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.ProductsRes{
		Data: prd,
	})
}

// DeleteProduct
// @Description delete product by id
// @Tags products
// @Accept json
// @Produce json
// @Param pid query int true "product id"
// @Success 200 {object} domain.OK
// @Router /products [DELETE]
func (p *Products) DeleteProduct(c echo.Context) error {
	sId := c.QueryParam("pid")
	if sId == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing param 'pid' required",
		})
	}
	id, err := strconv.ParseInt(sId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "param 'pid' must be integer",
		})
	}
	if err := p.Services.DeleteProductById(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.Error{
		Msg: "your product has been deleted successfully",
	})
}

// InsertCategory
// @Description Insert new category
// @Tags products
// @Accept json
// @Produce json
// @Param login body domain.CategoriesSwagger true "Categories Request"
// @Success 201 {object} domain.OK
// @Router /categories [POST]
func (p *Products) InsertCategory(c echo.Context) error {
	var request domain.Categories
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
	if err := p.Services.InsertCategory(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
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
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	files := form.File["img"]
	// get id params
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	// call services
	if err := p.Services.UploadProductImage(c.Request().Context(), id, files); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
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
// @Accept json
// @Produce json
// @Param product body domain.ProductReq true "Product Request"
// @Success 200 {object} domain.UploadProductRes
// @Router /products [POST]
func (p *Products) UploadProduct(c echo.Context) error {
	// bind json to structure
	var productReq domain.ProductReq
	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// check validate struct
	if validate := valid.Validate(&productReq); validate != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: validate,
		})
	}
	// call services
	id, err := p.Services.UploadProduct(c.Request().Context(), productReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.UploadProductRes{
		Msg: "upload product successfully",
		Id:  id,
	})
}

// GetSupplier
// @Description get suppliers information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of suppliers"
// @Success 200 {object} domain.SupplierSlices
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
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.SupplierSlices{
		Data: _supp,
	})
}

// InsertSupplier
// @Description insert new suppliers information
// @Tags products
// @Accept json
// @Produce json
// @Param SuppliersReq body domain.SuppliersReq true "Suppliers Request"
// @Success 201 {object} domain.OK
// @Router /suppliers [POST]
func (p *Products) InsertSupplier(c echo.Context) error {
	var req domain.SuppliersReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if validate := valid.Validate(req); validate != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: validate,
		})
	}
	if err := p.Services.InsertSuppliers(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "suppliers created successfully",
	})
}

// AddToWarehouse
// @Description add product to warehouse
// @Tags products
// @Accept json
// @Produce json
// @Param WarehouseStruct body domain.WarehouseStruct true "Warehouse Request"
// @Success 201 {object} domain.OK
// @Router /warehouse [POST]
func (p *Products) AddToWarehouse(c echo.Context) error {
	var req domain.WarehouseStruct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if validate := valid.Validate(req); validate != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: validate,
		})
	}
	if err := p.Services.InsertIntoWarehouse(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "add product to warehouse created successfully",
	})
}
