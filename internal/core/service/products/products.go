package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/errors"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/inventories"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/utils"

	"github.com/google/uuid"
)

// ProductService struct for product service
type ProductService struct {
	Blob       blob.IBlobStorage
	Categories categories.ICategoriesRepository
	Products   products.IProductRepository
	Suppliers  suppliers.ISuppliersRepository
	Inventory  inventories.IInventoryRepository
}

var _ IProductService = (*ProductService)(nil)

// New creates a new ProductService object
func New(
	blob blob.IBlobStorage,
	categories categories.ICategoriesRepository,
	products products.IProductRepository,
	suppliers suppliers.ISuppliersRepository,
	inventory inventories.IInventoryRepository,
) IProductService {
	return &ProductService{
		Blob:       blob,
		Categories: categories,
		Products:   products,
		Suppliers:  suppliers,
		Inventory:  inventory,
	}
}

// GetAllStock implements IProductService.
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
		stock.Stock = append(stock.Stock, domain.InventorySchema{
			ID:          _inventory.ID,
			ProductName: product.Name,
			InventoryStruct: domain.InventoryStruct{
				ProductID:    strconv.Itoa(int(_inventory.ProductID)),
				Price:        _inventory.Price.String(),
				Available:    _inventory.Available,
				CurrencyCode: _inventory.CurrencyCode,
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
func (s *ProductService) GetInventory(ctx context.Context, productID int64) ([]domain.Inventories, error) {
	return s.Inventory.GetByProductID(ctx, productID)
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

// FindDeviceInInventory implements IProductService.
func (s *ProductService) FindDeviceInInventory(
	ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.InventorySchema, error) {
	_inventory, err := s.Inventory.FindDevice(ctx, deviceSpecs)
	if err != nil {
		return nil, err
	}
	product, err := s.Products.GetByID(ctx, _inventory.ProductID)
	if err != nil {
		return nil, err
	}
	var inventoryRes = domain.InventorySchema{
		ID:          _inventory.ID,
		ProductName: product.Name,
		InventoryStruct: domain.InventoryStruct{
			ProductID:    _inventory.ID,
			Price:        _inventory.Price.String(),
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

// CreateSuppliers implements IProductService.
func (s *ProductService) CreateSuppliers(ctx context.Context, supplierReq domain.SupplierSchema) error {
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		supplier = domain.Suppliers{
			Name:  supplierReq.Name,
			Email: supplierReq.Email,
		}
		addr = domain.Addresses{
			City:     supplierReq.City,
			Ward:     supplierReq.Ward,
			District: supplierReq.District,
			Street:   supplierReq.Street,
		}
		supplierRepo = suppliers.New(tx)
		addressRepo  = addresses.New(tx)
	)
	if err := supplierRepo.Insert(ctx, supplier); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	supp, err := supplierRepo.GetByPhone(ctx, supplierReq.Email)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	addr.UUID = uuid.New().String()
	if err = addressRepo.Insert(ctx, addr); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	if err := supplierRepo.InsertAddress(ctx, domain.SuppliersAddress{
		SuppliersID: supp.ID,
		AddressUuiD: addr.UUID,
	}); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	return tx.Commit(ctx)
}

// DeleteProductByID implements IProductService.
func (s *ProductService) DeleteProductByID(ctx context.Context, productID int64) error {
	return s.Products.DeleteByID(ctx, productID)
}

// InsertIntoInventory implements IProductService.
func (s *ProductService) InsertIntoInventory(ctx context.Context, product domain.InventoryStruct) error {
	return s.Inventory.InsertProduct(ctx, product)
}

// GetCategoriesLimit implements IProductService.
func (s *ProductService) GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	return s.Categories.GetLimit(ctx, limit)
}

// GetProductsLimit implements IProductService.
func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]domain.ProductSchema, error) {
	products, err := s.Products.GetLimit(ctx, limit)
	if err != nil {
		return nil, err
	}
	var productResponse []domain.ProductSchema
	for _, p := range products {
		var spec domain.Specs
		if err := json.Unmarshal([]byte(p.Spec), &spec); err != nil {
			// don't find anything, just return empty object
			return nil, errors.Repository("json", err)
		}
		images := strings.Split(p.Image, ",")
		productResponse = append(productResponse,
			domain.ProductSchema{
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

// CreateCategory implements IProductService.
func (s *ProductService) CreateCategory(ctx context.Context, ctg domain.Categories) error {
	return s.Categories.Insert(ctx, ctg)
}

// GetSuppliersLimit implements IProductService.
func (s *ProductService) GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return s.Suppliers.GetLimit(ctx, limit)
}
