// Package controller This file includes all the product controller functions.
package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/enum"
	"swclabs/swix/internal/core/service/products"
	"swclabs/swix/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

// IProducts interface for products controller
type IProducts interface {
	Search(c echo.Context) error

	GetProductLimit(c echo.Context) error
	UploadProductImage(c echo.Context) error
	CreateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	UpdateProductInfo(c echo.Context) error
	GetProductStorageDetails(c echo.Context) error
	GetProductView(c echo.Context) error

	GetInvDetails(c echo.Context) error
	AddInv(c echo.Context) error
	DeleteInv(c echo.Context) error
	UploadInvImage(c echo.Context) error
	GetStock(c echo.Context) error
	UpdateInv(c echo.Context) error
	InsertInvSpecs(c echo.Context) error
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

// Search .
// @Description get product
// @Tags search
// @Accept json
// @Produce json
// @Param keyword query string true "keyword"
// @Success 200 {object} []dtos.ProductResponse
// @Router /search [GET]
func (p *Products) Search(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'keyword' query parameter",
		})
	}
	product, err := p.Services.Search(c.Request().Context(), keyword)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// InsertInvSpecs .
// @Description create new specification for inventory
// @Tags inventories
// @Accept json
// @Produce json
// @Param type path string true "Specification type (storage or wireless)"
// @Param spec body dtos.Object true "Storage Specification"
// @Success 201 {object} dtos.OK
// @Failure 400 {object} dtos.Error "Bad Request"
// @Failure 500 {object} dtos.Error "Internal Server Error"
// @Router /inventories/specs/{type} [POST]
func (p *Products) InsertInvSpecs(c echo.Context) error {
	types := c.Param("type")
	switch types {
	case "storage":
		var invSpec dtos.Storage
		if err := c.Bind(&invSpec); err != nil {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		if _valid := valid.Validate(&invSpec); _valid != nil {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: _valid.Error(),
			})
		}
		if err := p.Services.InsertSpecStorage(c.Request().Context(), invSpec); err != nil {
			if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
				return c.JSON(http.StatusBadRequest, dtos.Error{
					Msg: err.Error(),
				})
			}
			return c.JSON(http.StatusInternalServerError, dtos.Error{
				Msg: err.Error(),
			})
		}
	case "wireless":
		var invSpec dtos.Wireless
		if err := c.Bind(&invSpec); err != nil {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		if _valid := valid.Validate(&invSpec); _valid != nil {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: _valid.Error(),
			})
		}
		if err := p.Services.InsertSpecWireless(c.Request().Context(), invSpec); err != nil {
			if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
				return c.JSON(http.StatusBadRequest, dtos.Error{
					Msg: err.Error(),
				})
			}
			return c.JSON(http.StatusInternalServerError, dtos.Error{
				Msg: err.Error(),
			})
		}
	default:
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "invalid type",
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "your specification has been created successfully",
	})
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
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// GetProductStorageDetails .
// @Description get product details
// @Tags products
// @Accept json
// @Produce json
// @Param id query number true "product id"
// @Success 200 {object} dtos.ProductDetail[dtos.DetailStorage]
// @Router /products/details [GET]
func (p *Products) GetProductStorageDetails(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}

	product, err := p.Services.ProductDetail(c.Request().Context(), int64(ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// UpdateInv .
// @Description update inventory
// @Tags inventories
// @Accept json
// @Produce json
// @Param inventory body dtos.InvUpdate true "Inventory Request"
// @Success 200 {object} dtos.OK
// @Router /inventories [PUT]
func (p *Products) UpdateInv(c echo.Context) error {
	var inventory dtos.InvUpdate
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
	if err := p.Services.UpdateInv(c.Request().Context(), inventory); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "your inventory has been updated successfully",
	})
}

// UploadInvImage .
// @Description update inventory image
// @Tags inventories
// @Accept json
// @Produce json
// @Param image formData file true "stock image"
// @Success 200 {object} dtos.OK
// @Router /inventories/image [PUT]
func (p *Products) UploadInvImage(c echo.Context) error {
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
	if err := p.Services.UploadInvImage(c.Request().Context(), id, files); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.Error{
		Msg: "your inventory image has been uploaded successfully",
	})
}

// DeleteInv .
// @Description delete inventory by id
// @Tags inventories
// @Accept json
// @Produce json
// @Param id query int true "inventory id"
// @Success 200 {object} dtos.OK
// @Router /inventories [DELETE]
func (p *Products) DeleteInv(c echo.Context) error {
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
	if err := p.Services.DeleteInvByID(c.Request().Context(), id); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
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
// @Success 200 {object} dtos.InvStock[dtos.InvStorage]
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
	stock, err := p.Services.GetAllInv(c.Request().Context(), page, limit)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
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
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "your product has been updated successfully",
	})
}

// GetInvDetails .
// @Description get product availability in inventories
// @Tags inventories
// @Accept json
// @Produce json
// @Param id query number true "inventory id"
// @Success 200 {object} dtos.Inventory[dtos.InvStorage]
// @Router /inventories/details [GET]
func (p *Products) GetInvDetails(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}

	product, err := p.Services.GetInvByID(c.Request().Context(), int64(ID))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
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
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
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
	if err := p.Services.DelProductByID(c.Request().Context(), id); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
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
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
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
		Specs:       productReq.Specs,
		Price:       productReq.Price,
		Description: productReq.Description,
		Name:        productReq.Name,
		SupplierID:  productReq.SupplierID,
		CategoryID:  productReq.CategoryID,
		Status:      productReq.Status,
	}
	ID, err := p.Services.CreateProduct(c.Request().Context(), product)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.CreateProduct{
		Msg: "upload product successfully",
		ID:  ID,
	})
}

// AddInv .
// @Description add product to inventories
// @Tags inventories
// @Accept json
// @Produce json
// @Param type path string true "Specification type (storage or wireless)"
// @Param InvDetail body dtos.InvDetail[dtos.Object] true "Inventories Request"
// @Success 201 {object} dtos.OK
// @Router /inventories/{type} [POST]
func (p *Products) AddInv(c echo.Context) error {

	var (
		types = c.Param("type")
		specs []interface{}
		inv   dtos.Inventory[interface{}]
	)
	switch types {
	case "storage":
		var req dtos.InvDetail[dtos.StorageReq]
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
		specs = make([]interface{}, len(req.Specs))
		for i, s := range req.Specs {
			specs[i] = s
		}
		inv = dtos.Inventory[interface{}]{
			ProductID:    req.ProductID,
			Price:        req.Price,
			Available:    req.Available,
			CurrencyCode: req.CurrencyCode,
			ColorImg:     req.ColorImg,
			Color:        req.Color,
			Status:       req.Status,
			Image:        req.Image,
			Specs:        specs,
		}

	case "wireless":
		var req dtos.InvDetail[dtos.WirelessReq]
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
		specs := make([]interface{}, len(req.Specs))
		for i, s := range req.Specs {
			specs[i] = s
		}
		inv = dtos.Inventory[interface{}]{
			ProductID:    req.ProductID,
			Price:        req.Price,
			Available:    req.Available,
			CurrencyCode: req.CurrencyCode,
			ColorImg:     req.ColorImg,
			Color:        req.Color,
			Status:       req.Status,
			Image:        req.Image,
			Specs:        specs,
		}
	default:
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "invalid type",
		})
	}
	if err := p.Services.InsertInv(c.Request().Context(), inv); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("[code %d]", http.StatusBadRequest)) {
			return c.JSON(http.StatusBadRequest, dtos.Error{
				Msg: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "add product to inventories created successfully",
	})
}
