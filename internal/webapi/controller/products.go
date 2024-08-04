// Package controller This file includes all the product controller functions.
package controller

import (
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/enum"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

// IProducts interface for products controller
type IProducts interface {
	GetProductLimit(c echo.Context) error
	UploadProductImage(c echo.Context) error
	CreateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	UpdateProductInfo(c echo.Context) error
	GetProductDetails(c echo.Context) error
	GetProductView(c echo.Context) error

	GetInventoryDetails(c echo.Context) error
	AddToInventory(c echo.Context) error
	DeleteInventory(c echo.Context) error
	UploadInventoryImage(c echo.Context) error
	GetStock(c echo.Context) error
	UpdateInventory(c echo.Context) error
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

// GetProductView .
// @Description get product view
// @Tags products
// @Accept json
// @Produce json
// @Param type query string true "type of product"
// @Success 200 {object} []dtos.ProductView
// @Router /products/view [GET]
func (p *Products) GetProductView(c echo.Context) error {
	var types enum.Category
	if err := types.Load(c.QueryParam("type")); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	product, err := p.Services.ViewDataOf(c.Request().Context(), types, 0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// GetProductDetails .
// @Description get product details
// @Tags products
// @Accept json
// @Produce json
// @Param id query number true "product id"
// @Success 200 {object} dtos.ProductDetail
// @Router /products/details [GET]
func (p *Products) GetProductDetails(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}

	product, err := p.Services.ProductDetailOf(c.Request().Context(), int64(ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// UpdateInventory .
// @Description update inventory
// @Tags inventories
// @Accept json
// @Produce json
// @Param inventory body dtos.UpdateInventory true "Inventory Request"
// @Success 200 {object} dtos.OK
// @Router /inventories [PUT]
func (p *Products) UpdateInventory(c echo.Context) error {
	var inventory dtos.UpdateInventory
	if err := c.Bind(&inventory); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&inventory); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	if err := p.Services.UpdateInventory(c.Request().Context(), inventory); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "your inventory has been updated successfully",
	})
}

// UploadInventoryImage .
// @Description update inventory image
// @Tags inventories
// @Accept json
// @Produce json
// @Param image formData file true "stock image"
// @Success 200 {object} dtos.OK
// @Router /inventories/image [PUT]
func (p *Products) UploadInventoryImage(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	files := form.File["image"]
	// get id params
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	if err := p.Services.UploadStockImage(c.Request().Context(), id, files); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.Error{
		Msg: "your inventory image has been uploaded successfully",
	})
}

// DeleteInventory .
// @Description delete inventory by id
// @Tags inventories
// @Accept json
// @Produce json
// @Param id query int true "inventory id"
// @Success 200 {object} dtos.OK
// @Router /inventories [DELETE]
func (p *Products) DeleteInventory(c echo.Context) error {
	iID := c.QueryParam("id")
	if iID == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing param 'id' required",
		})
	}
	id, err := strconv.ParseInt(iID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "param 'id' must be integer",
		})
	}
	if err := p.Services.DeleteInventoryByID(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.Error{
		Msg: "your inventory has been deleted successfully",
	})
}

// GetStock .
// @Description get all product from inventory
// @Tags inventories
// @Accept json
// @Produce json
// @Param page query number true "page"
// @Param limit query number true "limit"
// @Success 200 {object} dtos.StockInInventory
// @Router /inventories [GET]
func (p *Products) GetStock(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'page' or 'page' is not a number",
		})
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'limit' or 'limit' is not a number",
		})
	}
	stock, err := p.Services.GetAllStock(c.Request().Context(), page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
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
// @Param product body dtos.UpdateProductInfo true "Product Information Request"
// @Success 200 {object} dtos.OK
// @Router /products [PUT]
func (p *Products) UpdateProductInfo(c echo.Context) error {
	var payload dtos.UpdateProductInfo
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
	if err := p.Services.UpdateProductInfo(c.Request().Context(), payload); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "your product has been updated successfully",
	})
}

// GetInventoryDetails .
// @Description get product availability in inventories
// @Tags inventories
// @Accept json
// @Produce json
// @Param id query number true "inventory id"
// @Success 200 {object} dtos.Inventory
// @Router /inventories/details [GET]
func (p *Products) GetInventoryDetails(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}

	product, err := p.Services.GetInventoryByID(c.Request().Context(), int64(ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// GetProductLimit .
// @Description get product information
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int true "limit number of products"
// @Success 200 {object} dtos.Slices[dtos.ProductResponse]
// @Router /products [GET]
func (p *Products) GetProductLimit(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	prd, err := p.Services.GetProductsLimit(c.Request().Context(), _limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.Slices[dtos.ProductResponse]{
		Body: prd,
	})
}

// DeleteProduct .
// @Description delete product by id
// @Tags products
// @Accept json
// @Produce json
// @Param pid query int true "product id"
// @Success 200 {object} dtos.OK
// @Router /products [DELETE]
func (p *Products) DeleteProduct(c echo.Context) error {
	sID := c.QueryParam("pid")
	if sID == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing param 'pid' required",
		})
	}
	id, err := strconv.ParseInt(sID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "param 'pid' must be integer",
		})
	}
	if err := p.Services.DeleteProductByID(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.Error{
		Msg: "your product has been deleted successfully",
	})
}

// UploadProductImage .
// @Description insert new product image
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param id query string true "id of product"
// @Param img formData file true "image of product"
// @Success 200 {object} dtos.OK
// @Failure 400 {object} dtos.Error
// @Router /products/img [POST]
func (p *Products) UploadProductImage(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	files := form.File["img"]
	// get id params
	id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	// call services
	if err := p.Services.UploadProductImage(c.Request().Context(), id, files); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "update successfully",
	})
}

// CreateProduct .
// @Description create new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body dtos.ProductRequest true "Product Request"
// @Success 200 {object} dtos.CreateProduct
// @Router /products [POST]
func (p *Products) CreateProduct(c echo.Context) error {
	// bind json to structure
	var productReq dtos.ProductRequest
	if err := c.Bind(&productReq); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	// check validate struct
	if validate := valid.Validate(&productReq); validate != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: validate.Error(),
		})
	}
	// call services
	product := dtos.Product{
		Price:       productReq.Price,
		Description: productReq.Description,
		Name:        productReq.Name,
		SupplierID:  productReq.SupplierID,
		CategoryID:  productReq.CategoryID,
		Status:      productReq.Status,
	}
	id, err := p.Services.CreateProduct(c.Request().Context(), product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.CreateProduct{
		Msg: "upload product successfully",
		ID:  id,
	})
}

// AddToInventory .
// @Description add product to inventories
// @Tags inventories
// @Accept json
// @Produce json
// @Param InventoryDetail body dtos.InventoryDetail true "Inventories Request"
// @Success 201 {object} dtos.OK
// @Router /inventories [POST]
func (p *Products) AddToInventory(c echo.Context) error {
	var req dtos.InventoryDetail
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
	if err := p.Services.InsertIntoInventory(c.Request().Context(),
		dtos.Inventory{
			ProductID:    req.ProductID,
			Price:        req.Price,
			Available:    req.Available,
			CurrencyCode: req.CurrencyCode,
			ColorImg:     req.ColorImg,
			Color:        req.Color,
			Status:       req.Status,
			Image:        req.Image,
			Specs:        req.Specs,
		}); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "add product to inventories created successfully",
	})
}
