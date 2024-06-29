package products

import (
	"context"
	"encoding/json"
	"errors"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/inventory"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/pkg/blob"
)

type ProductService struct {
	Categories categories.ICategoriesRepository
	Products   products.IProductRepository
	Suppliers  suppliers.ISuppliersRepository
	Inventory  inventory.IInventoryRepository
}

var _ IProductService = (*ProductService)(nil)

func New(
	categories categories.ICategoriesRepository,
	products products.IProductRepository,
	suppliers suppliers.ISuppliersRepository,
	inventory inventory.IInventoryRepository,
) IProductService {
	return &ProductService{
		Categories: categories,
		Products:   products,
		Suppliers:  suppliers,
		Inventory:  inventory,
	}
}

// DeleteProductById implements domain.IProductService.
func (s *ProductService) DeleteProductById(ctx context.Context, productId int64) error {
	return s.Products.DeleteById(ctx, productId)
}

// GetProductsInInventory implements domain.IProductService.
func (s *ProductService) GetProductsInInventory(
	ctx context.Context, productID, ram, ssd, color string) (*domain.InventorySchema, error) {
	_inventory, err := s.Inventory.GetProducts(ctx, productID, ram, ssd, color)
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

// InsertIntoInventory implements domain.IProductService.
func (s *ProductService) InsertIntoInventory(
	ctx context.Context, product domain.InventoryStruct) error {
	return s.Inventory.InsertProduct(ctx, product)
}

func (s *ProductService) GetCategoriesLimit(
	ctx context.Context, limit string) ([]domain.Categories, error) {
	return s.Categories.GetLimit(ctx, limit)
}

func (s *ProductService) GetProductsLimit(
	ctx context.Context, limit int) ([]domain.ProductRes, error) {
	return s.Products.GetLimit(ctx, limit)
}

func (s *ProductService) InsertCategory(
	ctx context.Context, ctg domain.Categories) error {
	return s.Categories.Insert(ctx, ctg)
}

func (s *ProductService) GetSuppliersLimit(
	ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return s.Suppliers.GetLimit(ctx, limit)
}

func (s *ProductService) UploadProductImage(
	ctx context.Context, Id int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return errors.New("missing image file")
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

func (s *ProductService) UploadProduct(
	ctx context.Context, products domain.ProductReq) (int64, error) {
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

func (s *ProductService) InsertSuppliers(
	ctx context.Context, supplierReq domain.SuppliersReq) error {
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
