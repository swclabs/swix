// Package controller This file includes all the product controller functions.
package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain/dto"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

// IProducts interface for products controller
type IProducts interface {
	GetCategories(c echo.Context) error
	GetProductLimit(c echo.Context) error
	GetSupplier(c echo.Context) error
	InsertCategory(c echo.Context) error
	InsertSupplier(c echo.Context) error
	UploadProductImage(c echo.Context) error
	CreateProduct(c echo.Context) error
	GetProductAvailability(c echo.Context) error
	AddToInventory(c echo.Context) error
	DeleteProduct(c echo.Context) error
	UpdateProductInfo(c echo.Context) error
	GetStock(c echo.Context) error
	DeleteInventory(c echo.Context) error
	UploadInventoryImage(c echo.Context) error
}

// NewProducts creates a new Products object
func NewProducts(services products.IProductService) IProducts {
	return &Products{
		Services: services,
	}
}

// Products struct implementation of IProducts
type Products struct {
	Services products.IProductService
}

// UploadInventoryImage .
// @Description update inventory image
// @Tags products
// @Accept json
// @Produce json
// @Param image formData file true "stock image"
// @Success 200 {object} dto.OK
// @Router /inventory/image [PUT]
func (p *Products) UploadInventoryImage(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	files := form.File["image"]
	// get id params
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	if err := p.Services.UploadStockImage(c.Request().Context(), id, files); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Error{
		Msg: "your inventory image has been uploaded successfully",
	})
}

// DeleteInventory .
// @Description delete inventory by id
// @Tags products
// @Accept json
// @Produce json
// @Param id query int true "inventory id"
// @Success 200 {object} dto.OK
// @Router /inventories [DELETE]
func (p *Products) DeleteInventory(c echo.Context) error {
	iID := c.QueryParam("id")
	if iID == "" {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "missing param 'id' required",
		})
	}
	id, err := strconv.ParseInt(iID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "param 'id' must be integer",
		})
	}
	if err := p.Services.DeleteInventoryByID(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Error{
		Msg: "your inventory has been deleted successfully",
	})
}

// GetStock .
// @Description get all product from inventory
// @Tags products
// @Accept json
// @Produce json
// @Param page query number true "page"
// @Param limit query number true "limit"
// @Success 200 {object} dto.StockInInventory
// @Router /inventories [GET]
func (p *Products) GetStock(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "missing 'page' or 'page' is not a number",
		})
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "missing 'limit' or 'limit' is not a number",
		})
	}
	stock, err := p.Services.GetAllStock(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, stock)
}

// UpdateProductInfo .
// @Description update product information
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.UpdateProductInfo true "Product Information Request"
// @Success 200 {object} dto.OK
// @Router /products [PUT]
func (p *Products) UpdateProductInfo(c echo.Context) error {
	var payload dto.UpdateProductInfo
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&payload); _valid != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: _valid.Error(),
		})
	}
	if err := p.Services.UpdateProductInfo(c.Request().Context(), payload); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.OK{
		Msg: "your product has been updated successfully",
	})
}

// GetProductAvailability .
// @Description get product availability in inventories
// @Tags products
// @Accept json
// @Produce json
// @Param pid query number true "product id"
// @Param ram query number true "ram"
// @Param ssd query number true "ssd"
// @Param color query string true "color"
// @Success 200 {object} dto.Inventory
// @Router /inventories/details [GET]
func (p *Products) GetProductAvailability(c echo.Context) error {
	var (
		pid   = c.QueryParam("pid")
		ram   = c.QueryParam("ram")
		ssd   = c.QueryParam("ssd")
		color = c.QueryParam("color")
	)
	if pid == "" {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "required 'limit' query params",
		})
	}

	product, err := p.Services.FindDeviceInInventory(c.Request().Context(),
		dto.InventoryDeviceSpecs{
			ProductID: pid,
			RAM:       ram,
			Ssd:       ssd,
			Color:     color,
		})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// GetCategories .
// @Description get categories
// @Tags products
// @Accept json
// @Produce json
// @Param limit query number true "limit number"
// @Success 200 {object} dto.Slices[entity.Categories]
// @Router /categories [GET]
func (p *Products) GetCategories(c echo.Context) error {
	limit := c.QueryParam("limit")
	if limit == "" {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "required 'limit' query params",
		})
	}

	resp, err := p.Services.GetCategoriesLimit(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.Slices[entity.Categories]{
		Body: resp,
	})
}

// GetProductLimit .
// @Description get product information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of products"
// @Success 200 {object} dto.Slices[dto.ProductSchema]
// @Router /products [GET]
func (p *Products) GetProductLimit(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	prd, err := p.Services.GetProductsLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Slices[dto.ProductSchema]{
		Body: prd,
	})
}

// DeleteProduct .
// @Description delete product by id
// @Tags products
// @Accept json
// @Produce json
// @Param pid query int true "product id"
// @Success 200 {object} dto.OK
// @Router /products [DELETE]
func (p *Products) DeleteProduct(c echo.Context) error {
	sID := c.QueryParam("pid")
	if sID == "" {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "missing param 'pid' required",
		})
	}
	id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "param 'pid' must be integer",
		})
	}
	if err := p.Services.DeleteProductByID(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Error{
		Msg: "your product has been deleted successfully",
	})
}

// InsertCategory .
// @Description insert new category
// @Tags products
// @Accept json
// @Produce json
// @Param login body entity.Categories true "Categories Request"
// @Success 201 {object} dto.OK
// @Router /categories [POST]
func (p *Products) InsertCategory(c echo.Context) error {
	var request entity.Categories
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: _valid.Error(),
		})
	}
	if err := p.Services.CreateCategory(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: "category data invalid",
		})
	}
	return c.JSON(http.StatusCreated, dto.OK{
		Msg: "category has been created",
	})
}

// UploadProductImage .
// @Description insert new product image
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param id query string true "id of product"
// @Param img formData file true "image of product"
// @Success 200 {object} dto.OK
// @Failure 400 {object} dto.Error
// @Router /products/img [POST]
func (p *Products) UploadProductImage(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	files := form.File["img"]
	// get id params
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	// call services
	if err := p.Services.UploadProductImage(c.Request().Context(), id, files); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dto.OK{
		Msg: "update successfully",
	})
}

// CreateProduct .
// @Description create new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.Product true "Product Request"
// @Success 200 {object} dto.CreateProduct
// @Router /products [POST]
func (p *Products) CreateProduct(c echo.Context) error {
	// bind json to structure
	var productReq dto.Product
	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	// check validate struct
	if validate := valid.Validate(&productReq); validate != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: validate.Error(),
		})
	}
	// call services
	id, err := p.Services.CreateProduct(c.Request().Context(), productReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dto.CreateProduct{
		Msg: "upload product successfully",
		ID:  id,
	})
}

// GetSupplier .
// @Description get suppliers information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of suppliers"
// @Success 200 {object} dto.Slices[entity.Suppliers]
// @Router /suppliers [GET]
func (p *Products) GetSupplier(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	_supp, err := p.Services.GetSuppliersLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.Slices[entity.Suppliers]{
		Body: _supp,
	})
}

// InsertSupplier .
// @Description insert new suppliers information
// @Tags products
// @Accept json
// @Produce json
// @Param Supplier body dto.Supplier true "Suppliers Request"
// @Success 201 {object} dto.OK
// @Router /suppliers [POST]
func (p *Products) InsertSupplier(c echo.Context) error {
	var req dto.Supplier
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if validate := valid.Validate(&req); validate != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: validate.Error(),
		})
	}
	if err := p.Services.CreateSuppliers(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dto.OK{
		Msg: "suppliers created successfully",
	})
}

// AddToInventory .
// @Description add product to inventories
// @Tags products
// @Accept json
// @Produce json
// @Param InventoryDetail body dto.InventoryDetail true "Inventories Request"
// @Success 201 {object} dto.OK
// @Router /inventories [POST]
func (p *Products) AddToInventory(c echo.Context) error {
	var req dto.InventoryDetail
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if validate := valid.Validate(&req); validate != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: validate.Error(),
		})
	}
	if err := p.Services.InsertIntoInventory(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dto.OK{
		Msg: "add product to inventories created successfully",
	})
}
