// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/xdto"
	"swclabs/swix/internal/core/repos/addresses"
	"swclabs/swix/internal/core/repos/carts"
	"swclabs/swix/internal/core/repos/categories"
	"swclabs/swix/internal/core/repos/deliveries"
	"swclabs/swix/internal/core/repos/inventories"
	"swclabs/swix/internal/core/repos/orders"
	"swclabs/swix/internal/core/repos/products"
	"swclabs/swix/internal/core/repos/users"
	"swclabs/swix/internal/core/x/ghnx"
	"swclabs/swix/pkg/infra/db"
	"swclabs/swix/pkg/utils"
	"time"

	"github.com/shopspring/decimal"
)

// New creates a new Purchase object
var New = app.Service(
	func(
		order orders.IOrders,
		cart carts.ICarts,
		user users.IUsers,
		inv inventories.IInventories,
		product products.IProducts,
		category categories.ICategories,
		address addresses.IAddress,
		delivery deliveries.IDeliveries,
		ghn ghnx.IGhnx,
	) IPurchase {
		return &Purchase{
			Cart:      cart,
			Order:     order,
			User:      user,
			Inventory: inv,
			Product:   product,
			Category:  category,
			Address:   address,
			Delivery:  delivery,
			Ghn:       ghn,
		}
	},
)

// Purchase struct for purchase service
type Purchase struct {
	Order     orders.IOrders
	Cart      carts.ICarts
	User      users.IUsers
	Category  categories.ICategories
	Product   products.IProducts
	Inventory inventories.IInventories
	Address   addresses.IAddress
	Delivery  deliveries.IDeliveries
	Ghn       ghnx.IGhnx
}

// CreateDeliveryOrder implements IPurchase.
func (p *Purchase) CreateDeliveryOrder(ctx context.Context, shopID int, order xdto.CreateOrderDTO) (*xdto.OrderDTO, error) {
	return p.Ghn.CreateOrder(ctx, shopID, order)
}

// DeliveryOrderInfo implements IPurchase.
func (p *Purchase) DeliveryOrderInfo(ctx context.Context, orderCode string) (*xdto.OrderInfoDTO, error) {
	return p.Ghn.OrderInfo(ctx, orderCode)
}

// AddressDistrict implements IPurchase.
func (p *Purchase) AddressDistrict(ctx context.Context, provinceID int) (*xdto.DistrictDTO, error) {
	return p.Ghn.District(ctx, provinceID)
}

// AddressProvince implements IPurchase.
func (p *Purchase) AddressProvince(ctx context.Context) (*xdto.ProvinceDTO, error) {
	return p.Ghn.Province(ctx)
}

// AddressWard implements IPurchase.
func (p *Purchase) AddressWard(ctx context.Context, districtID int) (*xdto.WardDTO, error) {
	return p.Ghn.Ward(ctx, districtID)
}

// CreateDelivery implements IPurchase.
func (p *Purchase) CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error {
	sendate, err := time.Parse(time.RFC3339, delivery.SentDate)
	if err != nil {
		sendate = time.Time{}
	}
	receivedate, err := time.Parse(time.RFC3339, delivery.ReceivedDate)
	if err != nil {
		receivedate = time.Time{}
	}
	return p.Delivery.Create(ctx, entity.Deliveries{
		UserID:       delivery.UserID,
		AddressID:    delivery.AddressID,
		Status:       delivery.Status,
		Method:       delivery.Method,
		Note:         delivery.Note,
		SentDate:     sendate,
		ReceivedDate: receivedate,
	})
}

// CreateDeliveryAddress implements IPurchase.
func (p *Purchase) CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error {
	return p.Address.Insert(ctx, entity.Addresses{
		UserID:   addr.UserID,
		Street:   addr.Street,
		City:     addr.City,
		Ward:     addr.Ward,
		District: addr.District,
	})
}

// GetDelivery implements IPurchase.
func (p *Purchase) GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error) {
	deliveries, err := p.Delivery.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var delivery = []dtos.Delivery{}
	for _, del := range deliveries {
		var (
			sentdate     string
			receiveddate string
		)
		if !del.SentDate.IsZero() {
			sentdate = del.SentDate.Format(time.RFC3339)
		}
		if !del.ReceivedDate.IsZero() {
			receiveddate = del.ReceivedDate.Format(time.RFC3339)
		}
		delivery = append(delivery, dtos.Delivery{
			ID:           del.ID,
			AddressID:    del.AddressID,
			UserID:       del.UserID,
			Status:       del.Status,
			Method:       del.Method,
			Note:         del.Note,
			SentDate:     sentdate,
			ReceivedDate: receiveddate,
		})
	}
	return delivery, nil
}

// GetDeliveryAddress implements IPurchase.
func (p *Purchase) GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error) {
	addrs, err := p.Address.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var addresses = []dtos.Address{}
	for _, addr := range addrs {
		addresses = append(addresses, dtos.Address{
			ID:       addr.ID,
			Street:   addr.Street,
			City:     addr.City,
			Ward:     addr.Ward,
			District: addr.District,
		})
	}
	return addresses, nil
}

// GetOrdersByUserID implements IPurchaseService.
func (p *Purchase) GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderSchema, error) {
	// Get orders by user ID
	orders, err := p.Order.Get(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	// Get user by user ID
	user, err := p.User.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var orderSchema = []dtos.OrderSchema{}

	for _, order := range orders {
		// Get products by order code
		_orders, err := p.Order.GetOrderItemByCode(ctx, order.UUID)
		if err != nil {
			return nil, fmt.Errorf("error getting order by UUID: %w", err)
		}
		// Merge product and order schema
		orderSchema = append(orderSchema, dtos.OrderSchema{
			Items:     _orders,
			ID:        order.ID,
			UserID:    user.ID,
			Time:      order.Time.Format(time.RFC3339),
			UUID:      order.UUID,
			Status:    order.Status,
			UserEmail: user.Email,
			Username:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		})
	}
	return orderSchema, nil
}

// DeleteItemFromCart implements IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, cartID int64) error {
	return p.Cart.RemoveItem(ctx, cartID)
}

// AddToCart implements IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart dtos.CartInsertDTO) error {
	user, err := p.User.GetByEmail(ctx, cart.Email)
	if err != nil {
		return fmt.Errorf("error getting user by email: %v", err)
	}
	return p.Cart.Insert(ctx, entity.Carts{
		UserID:      user.ID,
		InventoryID: cart.InventoryID,
		Quantity:    cart.Quantity,
	})
}

// GetCart implements IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userID int64, limit int) (*dtos.Carts, error) {
	carts, err := p.Cart.GetCartInfo(ctx, userID)
	if err != nil {
		return nil, err
	}
	var cartResp = dtos.Carts{
		UserID:   userID,
		Products: []dtos.Cart{},
	}
	for _, cart := range carts {
		var specs dtos.Specs
		if err := json.Unmarshal([]byte(cart.InventorySpecs), &specs); err != nil {
			return nil, fmt.Errorf("error unmarshal inventory specs: %v", err)
		}
		cartResp.Products = append(cartResp.Products, dtos.Cart{
			Name:           cart.Name,
			CartID:         cart.CartID,
			InventoryID:    cart.InventoryID,
			ProductID:      cart.ProductID,
			Quantity:       cart.Quantity,
			Color:          cart.Color,
			InventoryPrice: cart.InventoryPrice,
			CurrencyCode:   cart.CurrencyCode,
			InventoryImage: cart.InventoryImage,
			CategoryName:   cart.CategoryName,
			InventorySpecs: specs,
		})
	}
	return &cartResp, nil
}

// CreateOrders implements IPurchaseService.
func (p *Purchase) CreateOrders(ctx context.Context, createOrder dtos.CreateOrderDTO) (string, error) {
	tx, err := db.NewTransaction(ctx)
	if err != nil {
		return "", err
	}
	var (
		_uuid              = utils.GenerateOrderCode(16)
		inventoryRepo      = inventories.New(tx)
		orderRepo          = orders.New(tx)
		totalAmount        decimal.Decimal
		productTotalAmount []decimal.Decimal
	)
	for {
		_order, err := orderRepo.GetByUUID(ctx, _uuid)
		if err != nil || _order.UUID == "" {
			break
		}
		_uuid = utils.GenerateOrderCode(16)
	}
	for _, product := range createOrder.Products {
		inven, err := inventoryRepo.GetByID(ctx, product.InventoryID)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
		totalAmount = totalAmount.Add(inven.Price)
		productTotalAmount = append(
			productTotalAmount,
			inven.Price.Mul(decimal.NewFromInt32(int32(product.Quantity))))
	}
	user, err := p.User.GetByEmail(ctx, createOrder.Email)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return "", fmt.Errorf("error getting user by email: %v", err)
	}
	orderID, err := orderRepo.Create(ctx, entity.Orders{
		UUID:        _uuid,
		DeliveryID:  createOrder.DeleveryID,
		UserID:      user.ID,
		Status:      "pending",
		TotalAmount: totalAmount,
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return "", err
	}
	for idx, product := range createOrder.Products {
		if err := orderRepo.InsertProduct(ctx, entity.ProductInOrder{
			OrderID:     orderID,
			InventoryID: product.InventoryID,
			Quantity:    product.Quantity,
			TotalAmount: productTotalAmount[idx],
			SpecsID:     product.SpecsID,
		}); err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
	}
	return _uuid, tx.Commit(ctx)
}
