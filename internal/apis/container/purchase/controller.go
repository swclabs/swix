// Package purchase purchase controller for handling purchase request.
package purchase

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/x/ghn"
	"swclabs/swipex/internal/core/service/purchase"
	"swclabs/swipex/pkg/lib/crypto"
	"swclabs/swipex/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ IController = (*Controller)(nil)
var _ = app.Controller(NewController)

// NewController creates a new Purchase object
func NewController(services purchase.IPurchase) IController {
	return &Controller{services: services}
}

// IController interface for purchase controller
type IController interface {
	AddToCarts(c echo.Context) error
	GetCarts(c echo.Context) error
	DeleteCartItem(c echo.Context) error

	CreateOrder(c echo.Context) error
	CreateOrderForm(c echo.Context) error
	GetOrders(c echo.Context) error
	GetOrdersByCode(c echo.Context) error
	GetOrdersByAdmin(c echo.Context) error

	CreateDeliveryAddress(c echo.Context) error
	GetDeliveryAddress(c echo.Context) error
	CreateDelivery(c echo.Context) error
	GetDelivery(c echo.Context) error

	AddressProvince(c echo.Context) error
	AddressWard(c echo.Context) error
	AddressDistrict(c echo.Context) error

	CreateDeliveryOrder(c echo.Context) error
	DeliveryOrderInfo(c echo.Context) error

	GetCoupon(c echo.Context) error
	CreateCoupon(c echo.Context) error
	UseCoupon(c echo.Context) error
}

// Controller struct implementation of IPurchase
type Controller struct {
	services purchase.IPurchase
}

// GetOrdersByAdmin .
// @Description get list of orders.
// @Tags purchase
// @Accept json
// @Produce json
// @Param limit query string true "limit order"
// @Success 200 {object} []dtos.OrderInfo
// @Router /purchase/admin/orders [GET]
func (p *Controller) GetOrdersByAdmin(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	orderList, err := p.services.GetUsersByAdmin(c.Request().Context(), limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, orderList)
}

// CreateCoupon .
// @Description create coupon.
// @Tags purchase
// @Accept json
// @Produce json
// @Param coupon body dtos.CreateCoupon true "coupon request"
// @Success 200 {object} dtos.OrderInfo
// @Router /purchase/coupons [POST]
func (p *Controller) CreateCoupon(c echo.Context) error {
	var coupon dtos.CreateCoupon
	if err := c.Bind(&coupon); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _, err := p.services.CreateCoupon(c.Request().Context(), coupon); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: "your coupon has been created successfully",
	})
}

// GetCoupon .
// @Description get coupon.
// @Tags purchase
// @Accept json
// @Produce json
// @Success 200 {object} []dtos.Coupon
// @Router /purchase/coupons [GET]
func (p *Controller) GetCoupon(c echo.Context) error {
	coupons, err := p.services.GetCoupon(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, coupons)
}

// UseCoupon .
// @Description get coupon.
// @Tags purchase
// @Accept json
// @Produce json
// @Param code path string true "coupons code"
// @Success 200 {object} dtos.OK
// @Router /purchase/coupons/{code} [GET]
func (p *Controller) UseCoupon(c echo.Context) error {
	code := c.Param("code")
	userID, _, _ := crypto.Authenticate(c)
	if err := p.services.UseCoupon(c.Request().Context(), userID, code); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "your coupon has been used successfully",
	})
}

// GetOrdersByCode .
// @Description get order by code.
// @Tags delivery
// @Accept json
// @Produce json
// @Param code path string true "order code"
// @Success 200 {object} dtos.OrderInfo
// @Router /purchase/orders/{code} [GET]
func (p *Controller) GetOrdersByCode(c echo.Context) error {
	code := c.Param("code")
	// print(code)
	orders, err := p.services.GetOrderByCode(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, orders)
}

// CreateOrderForm .
// @Description create order.
// @Tags delivery
// @Accept json
// @Produce json
// @Param order body dtos.OrderForm true "order delivery body request"
// @Success 200 {object} dtos.OK
// @Router /purchase/admin/orders [POST]
func (p *Controller) CreateOrderForm(c echo.Context) error {
	var order dtos.OrderForm
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&order); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	msg, err := p.services.CreateOrderForm(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: fmt.Sprintf("your order %s has been created successfully", msg),
	})

}

// CreateDeliveryOrder .
// @Description create order delivery.
// @Tags delivery
// @Accept json
// @Produce json
// @Param order body ghn.CreateOrderDTO true "order delivery body request"
// @Success 200 {object} ghn.OrderDTO
// @Router /delivery/order [POST]
func (p *Controller) CreateDeliveryOrder(c echo.Context) error {
	var order ghn.CreateOrderDTO
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&order); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	resp, err := p.services.CreateDeliveryOrder(c.Request().Context(), order.ShopID, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}

// DeliveryOrderInfo .
// @Description get delivery order details by order code.
// @Tags delivery
// @Accept json
// @Produce json
// @Param code path string true "delivery order code"
// @Success 200 {object} ghn.OrderInfoDTO
// @Router /delivery/order/{code} [GET]
func (p *Controller) DeliveryOrderInfo(c echo.Context) error {
	orderCode := c.Param("code")
	orderInfo, err := p.services.DeliveryOrderInfo(c.Request().Context(), orderCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, orderInfo)
}

// AddressDistrict .
// @Description get district by province ID.
// @Tags address
// @Accept json
// @Produce json
// @Param province_id query number true "province id"
// @Success 200 {object} []entity.District
// @Router /address/district [GET]
func (p *Controller) AddressDistrict(c echo.Context) error {
	provinceID := c.QueryParam("province_id")
	resp, err := p.services.AddressDistrict(c.Request().Context(), provinceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}

// AddressProvince .
// @Description get province
// @Tags address
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Province
// @Router /address/province [GET]
func (p *Controller) AddressProvince(c echo.Context) error {
	resp, err := p.services.AddressProvince(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}

// AddressWard .
// @Description get ward by district ID.
// @Tags address
// @Accept json
// @Produce json
// @Param district_id query number true "district id"
// @Success 200 {object} []entity.Commune
// @Router /address/ward [GET]
func (p *Controller) AddressWard(c echo.Context) error {
	districtID := c.QueryParam("district_id")
	resp, err := p.services.AddressWard(c.Request().Context(), districtID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, resp)
}

// CreateDelivery .
// @Description create delivery info.
// @Tags delivery
// @Accept json
// @Produce json
// @Param addr body dtos.DeliveryBody true "delivery info request"
// @Success 200 {object} dtos.OK
// @Router /delivery [POST]
func (p *Controller) CreateDelivery(e echo.Context) error {
	var body dtos.DeliveryBody
	if err := e.Bind(&body); err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := p.services.CreateDelivery(e.Request().Context(), body); err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, dtos.OK{
		Msg: "your delivery info has been saved",
	})
}

// GetDelivery .
// @Description get delivery info by user id.
// @Tags delivery
// @Accept json
// @Produce json
// @Success 200 {object} dtos.OK
// @Router /delivery [GET]
func (p *Controller) GetDelivery(e echo.Context) error {
	userID, _, _ := crypto.Authenticate(e)
	del, err := p.services.GetDelivery(e.Request().Context(), userID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, del)
}

// GetDeliveryAddress .
// @Description get address delivery.
// @Tags address
// @Accept json
// @Produce json
// @Success 200 {object} []dtos.Address
// @Router /address [GET]
func (p *Controller) GetDeliveryAddress(e echo.Context) error {
	userID, _, _ := crypto.Authenticate(e)
	addr, err := p.services.GetDeliveryAddress(e.Request().Context(), userID)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, addr)
}

// CreateDeliveryAddress .
// @Description create address delivery.
// @Tags address
// @Accept json
// @Produce json
// @Param addr body dtos.DeliveryAddress true "address request"
// @Success 200 {object} []dtos.ProductResponse
// @Router /address [POST]
func (p *Controller) CreateDeliveryAddress(e echo.Context) error {
	var addr dtos.DeliveryAddress
	if err := e.Bind(&addr); err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := p.services.CreateDeliveryAddress(e.Request().Context(), addr); err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, dtos.OK{
		Msg: "your address has been saved",
	})
}

// GetOrders .
// @Description get list of orders.
// @Tags purchase
// @Accept json
// @Produce json
// @Param limit query string true "limit order"
// @Success 200 {object} []dtos.OrderInfo
// @Router /purchase/orders [GET]
func (p *Controller) GetOrders(c echo.Context) error {
	sLimit := c.QueryParam("limit")
	if sLimit == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'limit' required",
		})
	}
	limit, err := strconv.ParseInt(sLimit, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'limit' must be a positive integer",
		})
	}

	userID, _, _ := crypto.Authenticate(c)
	orders, err := p.services.GetOrdersByUserID(c.Request().Context(), userID, int(limit))
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
	return c.JSON(http.StatusOK, orders)
}

// CreateOrder .
// @Description create order.
// @Tags purchase
// @Accept json
// @Produce json
// @Param login body dtos.Order true "order insert request"
// @Success 200 {object} dtos.OK
// @Router /purchase/orders [POST]
func (p *Controller) CreateOrder(c echo.Context) error {
	var orderReq dtos.Order
	if err := c.Bind(&orderReq); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&orderReq); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	userID, email, _ := crypto.Authenticate(c)
	if orderReq.Customer.Email != email {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "email must be the same as the login user",
		})
	}
	msg, err := p.services.CreateOrders(c.Request().Context(), userID, orderReq)
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
	return c.JSON(http.StatusCreated, dtos.OK{
		Msg: fmt.Sprintf("your order %s has been created successfully", msg),
	})
}

// AddToCarts .
// @Description add item to carts.
// @Tags purchase
// @Accept json
// @Produce json
// @Param login body dtos.CartDTO true "cart insert request"
// @Success 200 {object} dtos.OK
// @Router /purchase/carts [POST]
func (p *Controller) AddToCarts(c echo.Context) error {
	var cartReq dtos.CartDTO
	if err := c.Bind(&cartReq); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&cartReq); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	_, email, _ := crypto.Authenticate(c)
	if err := p.services.AddToCart(
		c.Request().Context(), dtos.CartInsertDTO{CartDTO: cartReq, Email: email}); err != nil {
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
		Msg: "your item has been added to cart successfully",
	})
}

// GetCarts .
// @Description get list of items from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Success 200 {object} dtos.Carts
// @Router /purchase/carts [GET]
func (p *Controller) GetCarts(c echo.Context) error {
	userID, _, _ := crypto.Authenticate(c)
	if userID == -1 {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "session expired",
		})
	}
	items, err := p.services.GetCart(c.Request().Context(), userID, 10)
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
	return c.JSON(http.StatusOK, *items)
}

// DeleteCartItem .
// @Description delete item from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Param id path string true "inventory id"
// @Success 200 {object} dtos.OK
// @Router /purchase/carts/{id} [DELETE]
func (p *Controller) DeleteCartItem(c echo.Context) error {
	itemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'id' must be a positive integer",
		})
	}
	userID, _, _ := crypto.Authenticate(c)
	if err := p.services.DeleteItemFromCart(c.Request().Context(), itemID, userID); err != nil {
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
		Msg: "your item has been deleted successfully",
	})
}
