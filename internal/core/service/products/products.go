package products

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/domain/enum"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/lib/errors"
	"swclabs/swipecore/pkg/utils"

	"github.com/shopspring/decimal"
)

var _ IProductService = (*ProductService)(nil)

// New creates a new ProductService object
func New(
	blob blob.IBlobStorage,
	products products.IProductRepository,
	inventory inventories.IInventoryRepository,
) IProductService {
	return &ProductService{
		Blob:      blob,
		Products:  products,
		Inventory: inventory,
	}
}

// ProductService struct for product service
type ProductService struct {
	Blob      blob.IBlobStorage
	Products  products.IProductRepository
	Inventory inventories.IInventoryRepository
}

// ViewDataOf implements IProductService.
func (s *ProductService) ViewDataOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductView, error) {
	products, err := s.Products.GetByCategory(ctx, types, offset)
	if err != nil {
		return nil, err
	}
	var productView []dtos.ProductView
	for _, p := range products {
		var specs dtos.ProductSpecs
		if err := json.Unmarshal([]byte(p.Spec), &specs); err != nil {
			return nil, err
		}
		productView = append(productView, dtos.ProductView{
			ID:    p.ID,
			Price: p.Price,
			Desc:  p.Description,
			Name:  p.Name,
			Image: p.Image,
			Specs: specs,
		})
	}
	return productView, nil
}

// GetInventoryByID implements IProductService.
func (s *ProductService) GetInventoryByID(ctx context.Context, inventoryID int64) (*dtos.Inventory, error) {
	stock, err := s.Inventory.GetByID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	product, err := s.Products.GetByID(ctx, stock.ProductID)
	if err != nil {
		return nil, err
	}
	var (
		specs  dtos.InventorySpecification
		result = dtos.Inventory{
			ID:          stock.ID,
			ProductName: product.Name,
			InventoryDetail: dtos.InventoryDetail{
				ProductID:    strconv.Itoa(int(stock.ProductID)),
				Price:        stock.Price.String(),
				Available:    stock.Available,
				CurrencyCode: stock.CurrencyCode,
				Status:       stock.Status,
				Color:        stock.Color,
				ColorImg:     stock.ColorImg,
				Image:        strings.Split(stock.Image, ","),
				Specs:        nil,
			},
		}
	)
	if err := json.Unmarshal([]byte(stock.Specs), &specs); err == nil {
		result.Specs = specs
	}
	return &result, nil

}

// ProductDetailOf implements IProductService.
func (s *ProductService) ProductDetailOf(ctx context.Context, productID int64) (*dtos.ProductDetail, error) {
	var (
		stocks       []dtos.Inventory
		productSpecs dtos.ProductSpecs
		details      dtos.ProductDetail
	)

	rawStocks, err := s.Inventory.GetByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	rawProduct, err := s.Products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	for _, stock := range rawStocks {
		var (
			inventory = dtos.Inventory{
				ID:          stock.ID,
				ProductName: rawProduct.Name,
				InventoryDetail: dtos.InventoryDetail{
					ProductID:    strconv.Itoa(int(stock.ProductID)),
					Price:        stock.Price.String(),
					Available:    stock.Available,
					CurrencyCode: stock.CurrencyCode,
					Status:       stock.Status,
					Color:        stock.Color,
					ColorImg:     stock.ColorImg,
					Image:        strings.Split(stock.Image, ","),
					Specs:        nil,
				}}
			specification dtos.InventorySpecification
		)
		if err := json.Unmarshal([]byte(stock.Specs), &specification); err == nil {
			inventory.Specs = specification
		}
		stocks = append(stocks, inventory)
	}

	if err := json.Unmarshal([]byte(rawProduct.Spec), &productSpecs); err != nil {
		return nil, err
	}

	details.Name = rawProduct.Name
	details.Screen = productSpecs.Screen
	details.Display = productSpecs.Display
	details.Image = strings.Split(rawProduct.Image, ",")

	for _, stock := range stocks {
		details.Color = append(details.Color, dtos.DetailColor{
			Name:    stock.Color,
			Img:     stock.ColorImg,
			Product: stock.Image,
		})
		if stock.Specs != nil {
			details.SSD = append(details.SSD, dtos.DetailSSD{
				Value: stock.Specs.(dtos.InventorySpecification).SSD,
				Price: stock.Price,
			})
		}
	}

	return &details, nil
}

// UpdateInventory implements IProductService.
func (s *ProductService) UpdateInventory(ctx context.Context, inventory dtos.UpdateInventory) error {
	pid, _ := strconv.Atoi(inventory.ProductID)
	specs, _ := json.Marshal(inventory.Specs)
	price, _ := decimal.NewFromString(inventory.Price)
	return s.Inventory.Update(ctx, entity.Inventories{
		Price:        price,
		ProductID:    int64(pid),
		ID:           inventory.ID,
		Specs:        string(specs),
		Status:       inventory.Status,
		Available:    inventory.Available,
		CurrencyCode: inventory.CurrencyCode,
	})
}

// UploadStockImage implements IProductService.
func (s *ProductService) UploadStockImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("missing image file")
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := s.Blob.UploadImages(file)
		if err == nil {
			if err = s.Inventory.UploadImage(ctx, ID, resp.SecureURL); err == nil {
				if err = file.Close(); err != nil {
					return err
				}
			}
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteInventoryByID implements IProductService.
func (s *ProductService) DeleteInventoryByID(ctx context.Context, inventoryID int64) error {
	return s.Inventory.DeleteByID(ctx, inventoryID)
}

// GetAllStock implements IProductService.
func (s *ProductService) GetAllStock(ctx context.Context, page int, limit int) (*dtos.StockInInventory, error) {
	inventories, err := s.Inventory.GetLimit(ctx, limit, page)
	if err != nil {
		return nil, errors.Service("get stock", err)
	}
	var (
		stock dtos.StockInInventory
		specs dtos.InventorySpecification
	)

	for _, _inventory := range inventories {
		if err := json.Unmarshal([]byte(_inventory.Specs), &specs); err != nil {
			return nil, errors.Service("json unmarshal error", err)
		}
		product, err := s.Products.GetByID(ctx, _inventory.ProductID)
		if err != nil {
			return nil, err
		}
		switch _inventory.Status {
		case "active":
			stock.Header.Active++
		case "draft":
			stock.Header.Draft++
		case "archived":
			stock.Header.Active++
		}
		stock.Stock = append(stock.Stock, dtos.Inventory{
			ID:          _inventory.ID,
			ProductName: product.Name,
			InventoryDetail: dtos.InventoryDetail{
				ProductID:    strconv.Itoa(int(_inventory.ProductID)),
				Price:        _inventory.Price.String(),
				Available:    _inventory.Available,
				CurrencyCode: _inventory.CurrencyCode,
				Status:       _inventory.Status,
				Specs:        specs,
			},
		})
	}

	stock.Page = page
	stock.Limit = limit
	stock.Header.All = len(inventories)

	return &stock, nil
}

// GetInventory implements IProductService.
func (s *ProductService) GetInventory(ctx context.Context, productID int64) ([]entity.Inventories, error) {
	return s.Inventory.GetByProductID(ctx, productID)
}

// Search implements IProductService.
func (s *ProductService) Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error) {
	_products, err := s.Products.Search(ctx, keyword)
	if err != nil {
		return nil, errors.Service("keyword error", err)
	}
	var (
		productSchema []dtos.ProductResponse
		specs         dtos.ProductSpecs
	)
	for _, p := range _products {
		err := json.Unmarshal([]byte(p.Spec), &specs)
		if err != nil {
			return nil, errors.Service("failed to unmarshal", err)
		}
		productSchema = append(productSchema, dtos.ProductResponse{
			ID:          p.ID,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Status:      p.Status,
			Spec:        specs,
			Image:       strings.Split(p.Image, ","),
			Created:     utils.HanoiTimezone(p.Created),
		})
	}
	return productSchema, nil
}

// UpdateProductInfo implements IProductService.
func (s *ProductService) UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error {
	spec, err := json.Marshal(product.ProductRequest.ProductSpecs)
	if err != nil {
		return errors.Service("update product info", err)
	}
	return s.Products.Update(ctx,
		entity.Products{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			SupplierID:  product.SupplierID,
			CategoryID:  product.CategoryID,
			Status:      product.Status,
			Spec:        string(spec),
		})
}

// UploadProductImage implements IProductService.
func (s *ProductService) UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {

	if fileHeader == nil {
		return fmt.Errorf("missing image file")
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := s.Blob.UploadImages(file)
		if err != nil {
			return err
		}
		if err := s.Products.UploadNewImage(ctx, resp.SecureURL, ID); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}

// CreateProduct implements IProductService.
func (s *ProductService) CreateProduct(ctx context.Context, products dtos.ProductRequest) (int64, error) {
	specs, err := json.Marshal(dtos.ProductSpecs{
		Screen:  products.Screen,
		Display: products.Display,
		SSD:     products.SSD,
		RAM:     products.RAM,
	})
	if err != nil {
		return -1, err
	}
	var prd = entity.Products{
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Status:      products.Status,
		Spec:        string(specs),
	}
	return s.Products.Insert(ctx, prd)
}

// DeleteProductByID implements IProductService.
func (s *ProductService) DeleteProductByID(ctx context.Context, productID int64) error {
	return s.Products.DeleteByID(ctx, productID)
}

// InsertIntoInventory implements IProductService.
func (s *ProductService) InsertIntoInventory(ctx context.Context, product dtos.InventoryDetail) error {
	pid, _ := strconv.Atoi(product.ProductID)
	specs, _ := json.Marshal(product.Specs)
	price, _ := decimal.NewFromString(product.Price)
	return s.Inventory.InsertProduct(ctx, entity.Inventories{
		ProductID:    int64(pid),
		Specs:        string(specs),
		Price:        price,
		Available:    product.Available,
		CurrencyCode: product.CurrencyCode,
		Status:       "active",
	})
}

// GetProductsLimit implements IProductService.
func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]dtos.ProductResponse, error) {
	products, err := s.Products.GetLimit(ctx, limit)
	if err != nil {
		return nil, err
	}
	var productResponse []dtos.ProductResponse
	for _, p := range products {
		var spec dtos.ProductSpecs
		if err := json.Unmarshal([]byte(p.Spec), &spec); err != nil {
			// don't find anything, just return empty object
			return nil, errors.Repository("json", err)
		}
		images := strings.Split(p.Image, ",")
		productResponse = append(productResponse,
			dtos.ProductResponse{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     utils.HanoiTimezone(p.Created),
				Image:       images[1:],
				Spec:        spec,
			})
	}
	return productResponse, nil
}
