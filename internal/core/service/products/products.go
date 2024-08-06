package products

import (
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/enum"
	"swclabs/swix/internal/core/repository/categories"
	"swclabs/swix/internal/core/repository/inventories"
	"swclabs/swix/internal/core/repository/products"
	"swclabs/swix/internal/core/repository/specifications"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/lib/errors"
	"swclabs/swix/pkg/utils"

	"github.com/shopspring/decimal"
)

var _ IProductService = (*ProductService)(nil)

// New creates a new ProductService object
func New(
	blob blob.IBlobStorage,
	products products.IProductRepository,
	inventory inventories.IInventoryRepository,
	category categories.ICategoriesRepository,
	Spec specifications.ISpecifications,
) IProductService {
	return &ProductService{
		Blob:      blob,
		Products:  products,
		Inventory: inventory,
		Category:  category,
		Specs:     Spec,
	}
}

// ProductService struct for product service
type ProductService struct {
	Blob      blob.IBlobStorage
	Products  products.IProductRepository
	Inventory inventories.IInventoryRepository
	Category  categories.ICategoriesRepository
	Specs     specifications.ISpecifications
}

// InsertSpecs implements IProductService.
func (s *ProductService) InsertSpecs(ctx context.Context, specification dtos.Specifications) error {
	inventory, err := s.Inventory.GetByID(ctx, specification.InventoryID)
	if err != nil {
		return err
	}
	product, err := s.Products.GetByID(ctx, inventory.ProductID)
	if err != nil {
		return err
	}
	category, err := s.Category.GetByID(ctx, product.CategoryID)
	if err != nil {
		return err
	}
	var types enum.Category
	if err := types.Load(category.Name); err != nil {
		return fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
	}
	if types&enum.ElectronicDevice == 0 {
		return fmt.Errorf("[code %d] category not support specification", http.StatusBadRequest)
	}
	content, _ := json.Marshal(dtos.InvSpecification{
		RAM: specification.RAM,
		SSD: specification.SSD,
	})
	return s.Specs.Insert(ctx, entity.Specifications{
		InventoryID: specification.InventoryID,
		Content:     string(content),
	})
}

// ViewDataOf implements IProductService.
func (s *ProductService) ViewDataOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductView, error) {
	products, err := s.Products.GetByCategory(ctx, types, offset)
	if err != nil {
		return nil, err
	}
	var productView []dtos.ProductView
	for _, p := range products {
		_view := dtos.ProductView{
			ID:       p.ID,
			Price:    p.Price,
			Desc:     p.Description,
			Name:     p.Name,
			Image:    p.Image,
			Category: p.CategoryName,
		}
		if p.Specs != "" && types&enum.ElectronicDevice != 0 {
			var specs dtos.ProductSpecs
			if err := json.Unmarshal([]byte(p.Specs), &specs); err != nil {
				return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
			}
			_view.Specs = specs
		}
		productView = append(productView, _view)
	}
	return productView, nil
}

// GetInvByID implements IProductService.
func (s *ProductService) GetInvByID(ctx context.Context, inventoryID int64) (*dtos.Inventory, error) {
	stock, err := s.Inventory.GetByID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	product, err := s.Products.GetByID(ctx, stock.ProductID)
	if err != nil {
		return nil, err
	}
	var (
		result = dtos.Inventory{
			ID:           stock.ID,
			ProductName:  product.Name,
			ProductID:    strconv.Itoa(int(stock.ProductID)),
			Price:        stock.Price.String(),
			Available:    strconv.Itoa(int(stock.Available)),
			CurrencyCode: stock.CurrencyCode,
			Status:       stock.Status,
			Color:        stock.Color,
			ColorImg:     stock.ColorImg,
			Image:        strings.Split(stock.Image, ","),
			Specs:        nil,
		}
		specs []dtos.InvSpecification
	)
	specOfproduct, err := s.Specs.GetByInventoryID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	for _, spec := range specOfproduct {
		var _spec dtos.InvSpecification
		if err := json.Unmarshal([]byte(spec.Content), &_spec); err != nil {
			return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}
		_spec.ID = spec.ID
		specs = append(specs, _spec)
	}
	result.Specs = specs
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

	if err := json.Unmarshal([]byte(rawProduct.Specs), &productSpecs); err != nil {
		return nil, err
	}

	details.Name = rawProduct.Name
	details.Screen = productSpecs.Screen
	details.Display = productSpecs.Display
	details.Image = strings.Split(rawProduct.Image, ",")

	for _, stock := range rawStocks {
		var (
			inventory = dtos.Inventory{
				ID:           stock.ID,
				ProductName:  rawProduct.Name,
				ProductID:    strconv.Itoa(int(stock.ProductID)),
				Price:        stock.Price.String(),
				Available:    strconv.Itoa(int(stock.Available)),
				CurrencyCode: stock.CurrencyCode,
				Status:       stock.Status,
				Color:        stock.Color,
				ColorImg:     stock.ColorImg,
				Image:        strings.Split(stock.Image, ","),
				Specs:        nil,
			}
			specs []dtos.InvSpecification
		)
		specOfproduct, err := s.Specs.GetByInventoryID(ctx, stock.ID)
		if err != nil {
			return nil, err
		}
		for _, spec := range specOfproduct {
			var _spec dtos.InvSpecification
			if err := json.Unmarshal([]byte(spec.Content), &_spec); err != nil {
				return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
			}
			specs = append(specs, _spec)
		}
		inventory.Specs = specs
		stocks = append(stocks, inventory)
	}

	for _, stock := range stocks {
		var (
			detailColor = dtos.DetailColor{
				Name:    stock.Color,
				Img:     stock.ColorImg,
				Product: stock.Image,
				Specs:   nil,
			}
			detailSpec []dtos.DetailSpecs
		)
		for _, spec := range stock.Specs {
			detailSpec = append(detailSpec, dtos.DetailSpecs{
				RAM:   spec.RAM,
				SSD:   spec.SSD,
				Price: stock.Price,
			})
		}
		detailColor.Specs = append(detailColor.Specs, detailSpec...)
		details.Color = append(details.Color, detailColor)
	}

	return &details, nil
}

// UpdateInv implements IProductService.
func (s *ProductService) UpdateInv(ctx context.Context, inventory dtos.InvUpdate) error {
	pid, err := strconv.Atoi(inventory.ProductID)
	if err != nil {
		pid = -1
	}
	price, err := decimal.NewFromString(inventory.Price)
	if err != nil {
		price = decimal.NewFromInt(-1)
	}
	avai, err := strconv.ParseInt(inventory.Available, 10, 64)
	if err != nil {
		avai = -1
	}
	invID, _ := strconv.ParseInt(inventory.ID, 10, 64)
	return s.Inventory.Update(ctx, entity.Inventories{
		Price:        price,
		ID:           invID,
		Available:    avai,
		ProductID:    int64(pid),
		Status:       inventory.Status,
		CurrencyCode: inventory.CurrencyCode,
	})
}

// UploadInvStockImage implements IProductService.
func (s *ProductService) UploadInvStockImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
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

// DeleteInvByID implements IProductService.
func (s *ProductService) DeleteInvByID(ctx context.Context, inventoryID int64) error {
	return s.Inventory.DeleteByID(ctx, inventoryID)
}

// GetAllInvStock implements IProductService.
func (s *ProductService) GetAllInvStock(ctx context.Context, page int, limit int) (*dtos.InvStock, error) {
	inventories, err := s.Inventory.GetLimit(ctx, limit, page)
	if err != nil {
		return nil, errors.Service("get stock", err)
	}
	var stock dtos.InvStock

	for _, inv := range inventories {
		switch inv.Status {
		case "active":
			stock.Header.Active++
		case "draft":
			stock.Header.Draft++
		case "archived":
			stock.Header.Active++
		}
		var specs []dtos.InvSpecification
		specOfproduct, err := s.Specs.GetByInventoryID(ctx, inv.ID)
		if err != nil {
			return nil, err
		}
		for _, spec := range specOfproduct {
			var _spec dtos.InvSpecification
			if err := json.Unmarshal([]byte(spec.Content), &_spec); err != nil {
				return nil, err
			}
			specs = append(specs, _spec)
		}
		product, err := s.Products.GetByID(ctx, inv.ProductID)
		if err != nil {
			return nil, err
		}
		category, err := s.Category.GetByID(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}
		stock.Stock = append(stock.Stock, dtos.Inventory{
			Specs:        specs,
			ProductName:  product.Name,
			ProductID:    strconv.Itoa(int(inv.ProductID)),
			Image:        strings.Split(inv.Image, ","),
			Category:     category.Name,
			ID:           inv.ID,
			Price:        inv.Price.String(),
			Available:    strconv.Itoa(int(inv.Available)),
			CurrencyCode: inv.CurrencyCode,
			Status:       inv.Status,
			ColorImg:     inv.ColorImg,
			Color:        inv.Color,
		})
	}

	stock.Page = page
	stock.Limit = limit
	stock.Header.All = len(inventories)

	return &stock, nil
}

// GetInv implements IProductService.
func (s *ProductService) GetInv(ctx context.Context, productID int64) ([]entity.Inventories, error) {
	return s.Inventory.GetByProductID(ctx, productID)
}

// Search implements IProductService.
func (s *ProductService) Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error) {
	_products, err := s.Products.Search(ctx, keyword)
	if err != nil {
		return nil, errors.Service("keyword error", err)
	}
	var productSchema []dtos.ProductResponse
	for _, p := range _products {
		category, err := s.Category.GetByID(ctx, p.CategoryID)
		if err != nil {
			return nil, err
		}
		productSchema = append(productSchema, dtos.ProductResponse{
			ID:          p.ID,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Status:      p.Status,
			Image:       strings.Split(p.Image, ","),
			Created:     utils.HanoiTimezone(p.Created),
			Category:    category.Name,
		})
	}
	return productSchema, nil
}

// UpdateProductInfo implements IProductService.
func (s *ProductService) UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error {
	if product.CategoryID != 0 {
		_category, err := s.Category.GetByID(ctx, product.CategoryID)
		if err != nil {
			return fmt.Errorf("category not found %v", err)
		}
		var types enum.Category
		if err := types.Load(_category.Name); err != nil {
			return fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}
	}
	var (
		sampleSpec, _ = json.Marshal(dtos.ProductSpecs{})
		spec, _       = json.Marshal(product.Specs)
		sspec         = ""
	)
	if string(sampleSpec) != string(spec) {
		sspec = string(spec)
	}
	_product := entity.Products{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		SupplierID:  product.SupplierID,
		CategoryID:  product.CategoryID,
		Status:      product.Status,
		Specs:       sspec,
	}
	return s.Products.Update(ctx, _product)

}

// UploadProductImage implements IProductService.
func (s *ProductService) UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error {
	if fileHeader == nil {
		return fmt.Errorf("[code %d] missing file", http.StatusBadRequest)
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
func (s *ProductService) CreateProduct(ctx context.Context, products dtos.Product) (int64, error) {
	_category, err := s.Category.GetByID(ctx, products.CategoryID)
	if err != nil {
		return -1, fmt.Errorf("category not found %v", err)
	}

	var types enum.Category
	if err := types.Load(_category.Name); err != nil {
		return -1, fmt.Errorf("category invalid %v", err)
	}
	var prd = entity.Products{
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Status:      products.Status,
		Specs:       "{}",
	}
	if products.Specs != nil && types&enum.ElectronicDevice != 0 {
		var specs, ok = products.Specs.(dtos.ProductSpecs)
		if !ok {
			return -1, fmt.Errorf("[code: %d] invalid specifications", http.StatusBadRequest)
		}
		specsByte, _ := json.Marshal(specs)
		prd.Specs = string(specsByte)
	}
	return s.Products.Insert(ctx, prd)
}

// DeleteProductByID implements IProductService.
func (s *ProductService) DeleteProductByID(ctx context.Context, productID int64) error {
	return s.Products.DeleteByID(ctx, productID)
}

// InsertInv implements IProductService.
func (s *ProductService) InsertInv(ctx context.Context, product dtos.Inventory) error {
	var (
		pid, _    = strconv.Atoi(product.ProductID)
		price, _  = decimal.NewFromString(product.Price)
		avai, _   = strconv.Atoi(product.Available)
		inventory = entity.Inventories{
			Color:        product.Color,
			ColorImg:     product.ColorImg,
			Image:        strings.Join(product.Image, ","),
			ProductID:    int64(pid),
			Price:        price,
			Available:    int64(avai),
			CurrencyCode: product.CurrencyCode,
			Status:       "active",
		}
		types enum.Category
	)
	pID, _ := strconv.ParseInt(product.ProductID, 10, 64)
	p, err := s.Products.GetByID(ctx, pID)
	if err != nil {
		return err
	}
	category, err := s.Category.GetByID(ctx, p.CategoryID)
	if err != nil {
		return err
	}
	if err := types.Load(category.Name); err != nil {
		return fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
	}
	invID, err := s.Inventory.InsertProduct(ctx, inventory)
	if err != nil {
		return err
	}
	if types&enum.ElectronicDevice != 0 {
		for _, spec := range product.Specs {
			bSpec, _ := json.Marshal(spec)
			if err := s.Specs.Insert(ctx, entity.Specifications{
				InventoryID: invID,
				Content:     string(bSpec),
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetProductsLimit implements IProductService.
func (s *ProductService) GetProductsLimit(ctx context.Context, limit int) ([]dtos.ProductResponse, error) {
	products, err := s.Products.GetLimit(ctx, limit)
	if err != nil {
		return nil, err
	}
	var productResponse []dtos.ProductResponse
	for _, p := range products {
		var (
			product = dtos.ProductResponse{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     utils.HanoiTimezone(p.Created),
				Image:       strings.Split(p.Image, ","),
			}
			types enum.Category
		)
		category, err := s.Category.GetByID(ctx, p.CategoryID)
		if err != nil {
			return nil, err
		}
		if err := types.Load(category.Name); err != nil {
			return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}
		product.Category = category.Name
		productResponse = append(productResponse, product)
	}
	return productResponse, nil
}
