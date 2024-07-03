package products

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/errors"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/pkg/blob"
	"swclabs/swipecore/pkg/utils"
)

type ProductService struct {
	Categories categories.ICategoriesRepository
	Products   products.IProductRepository
	Suppliers  suppliers.ISuppliersRepository
	Inventory  inventories.IInventoryRepository
}

var _ IProductService = (*ProductService)(nil)

func New(
	categories categories.ICategoriesRepository,
	products products.IProductRepository,
	suppliers suppliers.ISuppliersRepository,
	inventory inventories.IInventoryRepository,
) IProductService {
	return &ProductService{
		Categories: categories,
		Products:   products,
		Suppliers:  suppliers,
		Inventory:  inventory,
	}
}

// GetStock implements IProductService.
func (s *ProductService) GetAllStock(ctx context.Context, page int, limit int) (*domain.InventoryStockSchema, error) {
	inventories, err := s.Inventory.GetLimit(ctx, limit, page)
	if err != nil {
		return nil, errors.Service("get stock", err)
	}
	var (
		stock domain.InventoryStockSchema
		specs domain.InventorySpecsDetail
	)

	for _, _inventory := range inventories {
		if err := json.Unmarshal([]byte(_inventory.Specs), &specs); err != nil {
			return nil, errors.Service("json unmarshal error", err)
		}
		switch _inventory.Status {
		case "active":
			stock.Active++
		case "draft":
			stock.Draft++
		case "archived":
			stock.Active++
		}
		stock.Stock = append(stock.Stock, domain.InventorySchema{
			Id: _inventory.Id,
			InventoryStruct: domain.InventoryStruct{
				ProductID:    strconv.Itoa(int(_inventory.ProductID)),
				Price:        _inventory.Price.String(),
				Model:        _inventory.Model,
				Available:    _inventory.Available,
				CurrencyCode: _inventory.CurrencyCode,
				Specs:        specs,
			},
		})
	}

	stock.Page = page
	stock.Limit = limit
	stock.All = len(inventories)

	return &stock, nil
}

// GetInventory implements IProductService.
func (s *ProductService) GetInventory(ctx context.Context, productId int64) ([]domain.Inventories, error) {
	return s.Inventory.GetByProductId(ctx, productId)
}

// Search implements IProductService.
func (s *ProductService) Search(ctx context.Context, keyword string) ([]domain.ProductSchema, error) {
	_products, err := s.Products.Search(ctx, keyword)
	if err != nil {
		return nil, errors.Service("keyword error", err)
	}
	var (
		productSchema []domain.ProductSchema
		specs         domain.Specs
	)
	for _, p := range _products {
		err := json.Unmarshal([]byte(p.Spec), &specs)
		if err != nil {
			return nil, errors.Service("failed to unmarshal", err)
		}
		productSchema = append(productSchema, domain.ProductSchema{
			ID:          p.ID,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Status:      p.Status,
			Spec:        specs,
			Image:       strings.Split(p.Image, ",")[1:],
			Created:     utils.HanoiTimezone(p.Created),
		})
	}
	return productSchema, nil
}

// UpdateProductInfo implements IProductService.
func (s *ProductService) UpdateProductInfo(ctx context.Context, product domain.UpdateProductInfo) error {
	spec, err := json.Marshal(product.Product.Specs)
	if err != nil {
		return errors.Service("update product info", err)
	}
	return s.Products.Update(ctx,
		domain.Products{
			ID:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			SupplierID:  product.SupplierID,
			CategoryID:  product.CategoryID,
			Status:      product.Status,
			Spec:        string(spec),
		})
}

// FindDeviceInInventory implements IProductService.
func (s *ProductService) FindDeviceInInventory(
	ctx context.Context, productID, ram, ssd, color string) (*domain.InventorySchema, error) {
	_inventory, err := s.Inventory.FindDevice(ctx, productID, ram, ssd, color)
	if err != nil {
		return nil, err
	}
	var inventoryRes = domain.InventorySchema{
		Id: _inventory.Id,
		InventoryStruct: domain.InventoryStruct{
			ProductID:    _inventory.Id,
			Price:        _inventory.Price.String(),
			Model:        _inventory.Model,
			Available:    _inventory.Available,
			CurrencyCode: _inventory.CurrencyCode,
		},
	}
	if err := json.Unmarshal([]byte(_inventory.Specs), &inventoryRes.Specs); err != nil {
		return &inventoryRes, nil // don't find anything, just return empty object
	}
	if inventoryRes.Available == "" {
		inventoryRes.Available = "0"
		return &inventoryRes, nil
	}
	return &inventoryRes, nil
}

func (s *ProductService) UploadProductImage(ctx context.Context, Id int, fileHeader []*multipart.FileHeader) error {

	if fileHeader == nil {
		return fmt.Errorf("missing image file")
	}
	for _, fileheader := range fileHeader {
		file, err := fileheader.Open()
		if err != nil {
			return err
		}
		resp, err := blob.UploadImages(blob.Connection(), file)
		if err != nil {
			return err
		}
		if err := s.Products.UploadNewImage(ctx, resp.SecureURL, Id); err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProductService) CreateProduct(ctx context.Context, products domain.Product) (int64, error) {
	specs, err := json.Marshal(domain.Specs{
		Screen:  products.Screen,
		Display: products.Display,
		SSD:     products.SSD,
		RAM:     products.RAM,
	})
	if err != nil {
		return -1, err
	}
	var prd = domain.Products{
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

func (s *ProductService) CreateSuppliers(ctx context.Context, supplierReq domain.SupplierSchema) error {
	supplier := domain.Suppliers{
		Name:  supplierReq.Name,
		Email: supplierReq.Email,
	}
	addr := domain.Addresses{
		City:     supplierReq.City,
		Ward:     supplierReq.Ward,
		District: supplierReq.District,
		Street:   supplierReq.Street,
	}
	return s.Suppliers.Insert(ctx, supplier, addr)
}

// DeleteProductById implements IProductService.
func (s *ProductService) DeleteProductById(ctx context.Context, productId int64) error {
	return s.Products.DeleteById(ctx, productId)
}

// InsertIntoInventory implements IProductService.
func (s *ProductService) InsertIntoInventory(ctx context.Context, product domain.InventoryStruct) error {
	return s.Inventory.InsertProduct(ctx, product)
}

func (s *ProductService) GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	return s.Categories.GetLimit(ctx, limit)
}

func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]domain.ProductSchema, error) {
	return s.Products.GetLimit(ctx, limit)
}

func (s *ProductService) CreateCategory(ctx context.Context, ctg domain.Categories) error {
	return s.Categories.Insert(ctx, ctg)
}

func (s *ProductService) GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return s.Suppliers.GetLimit(ctx, limit)
}
