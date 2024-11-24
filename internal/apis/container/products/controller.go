package products

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/enum"
	"github.com/swclabs/swipex/internal/core/service/products"
	"github.com/swclabs/swipex/pkg/lib/crypto"
	"github.com/swclabs/swipex/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewController)

// NewController creates a new Products object
func NewController(services products.IProducts) IController {
	return &Controller{
		service: services,
	}
}

// IController interface for products controller
type IController interface {
	Rating(c echo.Context) error

	Search(c echo.Context) error
	SearchDetails(c echo.Context) error

	GetProductLimit(c echo.Context) error
	GetProductInfo(c echo.Context) error
	UploadProductImage(c echo.Context) error
	UploadProductShopImage(c echo.Context) error
	CreateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
	UpdateProductInfo(c echo.Context) error
	GetProductDetails(c echo.Context) error
	GetProductByType(c echo.Context) error

	GetInvDetails(c echo.Context) error
	InsertInv(c echo.Context) error
	DeleteInv(c echo.Context) error
	UploadInvImage(c echo.Context) error
	UploadInvColorImage(c echo.Context) error
	GetItems(c echo.Context) error
	UpdateInv(c echo.Context) error

	AddBookmark(c echo.Context) error
	GetBookmark(c echo.Context) error
}

// Controller struct implementation of IProducts
type Controller struct {
	service products.IProducts
}

// GetProductInfo .
// @Description get product information
// @Tags products
// @Accept json
// @Produce json
// @Param id query int true "products id"
// @Success 200 {object} dtos.ProductResponse
// @Router /products/info [GET]
func (p *Controller) GetProductInfo(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}
	product, err := p.service.GetProductInfo(c.Request().Context(), int64(ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// AddBookmark .
// @Description add product to favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Param id path number true "inventory id"
// @Success 200 {object} dtos.OK
// @Router /favorite/{id} [POST]
func (p *Controller) AddBookmark(c echo.Context) error {
	userID, _, _ := crypto.Authenticate(c)
	inventoryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' param",
		})
	}
	if err := p.service.AddBookmark(c.Request().Context(), userID, inventoryID); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "save favorite state successfully",
	})
}

// GetBookmark .
// @Description get product from favorite
// @Tags favorite
// @Accept json
// @Produce json
// @Success 200 {object} []dtos.Bookmark
// @Router /favorite [GET]
func (p *Controller) GetBookmark(c echo.Context) error {
	userID, _, _ := crypto.Authenticate(c)
	bookmarks, err := p.service.GetBookmarks(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, bookmarks)
}

// Rating .
// @Description update inventory image
// @Tags ratings
// @Accept json
// @Produce json
// @Param star query number true "id of product"
// @Success 200 {object} dtos.OK
// @Router /rating/{id} [PUT]
func (p *Controller) Rating(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query",
		})
	}
	star, err := strconv.Atoi(c.QueryParam("star"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'star' query parameter",
		})
	}
	userID, _, _ := crypto.Authenticate(c)
	if err := p.service.Rating(c.Request().Context(), userID, id, star); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{Msg: "your rating has been updated successfully"})
}

// UploadInvColorImage .
// @Description update inventory image
// @Tags inventories
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "stock image"
// @Success 200 {object} dtos.OK
// @Router /inventories/image/color [PUT]
func (p *Controller) UploadInvColorImage(c echo.Context) error {
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
			Msg: "Invalid 'id' query parameter",
		})
	}
	if err := p.service.UploadItemColorImage(c.Request().Context(), id, files); err != nil {
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

// UploadProductShopImage .
// @Description insert new product image
// @Tags products
// @Accept multipart/form-data
// @Produce json
// @Param id query string true "id of product"
// @Param img formData file true "image of product"
// @Success 200 {object} dtos.OK
// @Failure 400 {object} dtos.Error
// @Router /products/images [PUT]
func (p *Controller) UploadProductShopImage(c echo.Context) error {
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
			Msg: "Invalid 'id' query parameter",
		})
	}
	// call services
	if err := p.service.UploadProductShopImage(c.Request().Context(), id, files); err != nil {
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

// SearchDetails .
// @Description get product
// @Tags search
// @Accept json
// @Produce json
// @Param key query string true "keyword"
// @Success 200 {object} []dtos.ProductDetail
// @Router /search/details [GET]
func (p *Controller) SearchDetails(c echo.Context) error {
	keyword := c.QueryParam("key")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'keyword' query parameter",
		})
	}
	userID, _, _ := crypto.Authenticate(c)
	product, err := p.service.SearchDetails(c.Request().Context(), userID, keyword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

// Search .
// @Description get product
// @Tags search
// @Accept json
// @Produce json
// @Param keyword query string true "keyword"
// @Success 200 {object} []dtos.ProductResponse
// @Router /search [GET]
func (p *Controller) Search(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'keyword' query parameter",
		})
	}
	product, err := p.service.Search(c.Request().Context(), keyword)
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

// GetProductByType .
// @Description get product view
// @Tags products
// @Accept json
// @Produce json
// @Param type path string true "product type"
// @Success 200 {object} []dtos.ProductDTO
// @Router /products/{type} [GET]
func (p *Controller) GetProductByType(c echo.Context) error {
	var types enum.Category
	if err := types.Load(c.Param("type")); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	product, err := p.service.ProductType(c.Request().Context(), types, 0)
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

// GetProductDetails .
// @Description get product details
// @Tags products
// @Accept json
// @Produce json
// @Param id query number true "product id"
// @Success 200 {object} dtos.ProductDetail
// @Router /products/details [GET]
func (p *Controller) GetProductDetails(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}
	userID, _, _ := crypto.Authenticate(c)
	product, err := p.service.Detail(c.Request().Context(), userID, int64(ID))
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
func (p *Controller) UpdateInv(c echo.Context) error {
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
	if err := p.service.UpdateItem(c.Request().Context(), inventory); err != nil {
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
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "stock image"
// @Success 200 {object} dtos.OK
// @Router /inventories/image [PUT]
func (p *Controller) UploadInvImage(c echo.Context) error {
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
			Msg: "Invalid 'id' query parameter",
		})
	}
	if err := p.service.UploadItemImage(c.Request().Context(), id, files); err != nil {
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
func (p *Controller) DeleteInv(c echo.Context) error {
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
	if err := p.service.DeleteItem(c.Request().Context(), id); err != nil {
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

// GetItems .
// @Description get all product from inventory
// @Tags inventories
// @Accept json
// @Produce json
// @Param page query number true "page"
// @Param limit query number true "limit"
// @Success 200 {object} dtos.InventoryItems
// @Router /inventories [GET]
func (p *Controller) GetItems(c echo.Context) error {
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
	items, err := p.service.GetInvItems(c.Request().Context(), page, limit)
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
	return c.JSON(http.StatusOK, items)
}

// UpdateProductInfo .
// @Description update product information
// @Tags products
// @Accept json
// @Produce json
// @Param product body dtos.UpdateProductInfo true "Product Information Request"
// @Success 200 {object} dtos.OK
// @Router /products [PUT]
func (p *Controller) UpdateProductInfo(c echo.Context) error {
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
	if err := p.service.UpdateProductInfo(c.Request().Context(), payload); err != nil {
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
// @Success 200 {object} dtos.Inventory
// @Router /inventories/details [GET]
func (p *Controller) GetInvDetails(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'id' query parameter",
		})
	}

	product, err := p.service.GetItem(c.Request().Context(), int64(ID))
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
func (p *Controller) GetProductLimit(c echo.Context) error {
	_limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "Invalid 'limit' query parameter",
		})
	}
	prd, err := p.service.GetProducts(c.Request().Context(), _limit)
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
func (p *Controller) DeleteProduct(c echo.Context) error {
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
	if err := p.service.DelProduct(c.Request().Context(), id); err != nil {
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
// @Router /products/thumbnail [PUT]
func (p *Controller) UploadProductImage(c echo.Context) error {
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
	if err := p.service.UploadProductImage(c.Request().Context(), id, files); err != nil {
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
func (p *Controller) CreateProduct(c echo.Context) error {
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
	ID, err := p.service.CreateProduct(c.Request().Context(), product)
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

// InsertInv .
// @Description add product to inventories
// @Tags inventories
// @Accept json
// @Produce json
// @Param InvDetail body dtos.InventoryDetail true "Inventories Request"
// @Success 201 {object} dtos.OK
// @Router /inventories [POST]
func (p *Controller) InsertInv(c echo.Context) error {
	var (
		inv dtos.Inventory
		req dtos.InventoryDetail
	)
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
	inv = dtos.Inventory{
		ProductID:    req.ProductID,
		Price:        req.Price,
		Available:    req.Available,
		CurrencyCode: req.CurrencyCode,
		ColorImg:     req.ColorImg,
		Color:        req.Color,
		Status:       req.Status,
		Image:        req.Image,
		Specs:        req.Specs,
	}
	if err := p.service.InsertItem(c.Request().Context(), inv); err != nil {
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
