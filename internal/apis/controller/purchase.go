// Package controller purchase controller for handling purchase request.
package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/purchase"
	"swclabs/swix/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ IPurchase = (*Purchase)(nil)
var _ = app.Controller(NewPurchase)

// NewPurchase creates a new Purchase object
func NewPurchase(services purchase.IPurchase) IPurchase {
	return &Purchase{services: services}
}

// IPurchase interface for purchase controller
type IPurchase interface {
	AddToCarts(c echo.Context) error
	GetCarts(c echo.Context) error
	DeleteItem(c echo.Context) error
	CreateOrder(c echo.Context) error
	GetOrders(c echo.Context) error
	CreateDeliveryAddress(e echo.Context) error
	GetDeliveryAddress(e echo.Context) error
	CreateDelivery(e echo.Context) error
	GetDelivery(e echo.Context) error
}

// Purchase struct implementation of IPurchase
type Purchase struct {
	services purchase.IPurchase
}

// CreateDelivery .
// @Description create delivery info.
// @Tags delivery
// @Accept json
// @Produce json
// @Param addr body dtos.DeliveryBody true "delivery info request"
// @Success 200 {object} dtos.OK
// @Router /delivery [POST]
func (p *Purchase) CreateDelivery(e echo.Context) error {
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
// @Param uid query string true "user id"
// @Success 200 {object} dtos.OK
// @Router /delivery [GET]
func (p *Purchase) GetDelivery(e echo.Context) error {
	uid, err := strconv.Atoi(e.QueryParam("uid"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'uid' required",
		})
	}

	del, err := p.services.GetDelivery(e.Request().Context(), int64(uid))
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
// @Param uid query string true "user id"
// @Success 200 {object} []dtos.Address
// @Router /address [GET]
func (p *Purchase) GetDeliveryAddress(e echo.Context) error {
	uid, err := strconv.Atoi(e.QueryParam("uid"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'uid' required",
		})
	}

	addr, err := p.services.GetDeliveryAddress(e.Request().Context(), int64(uid))
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
func (p *Purchase) CreateDeliveryAddress(e echo.Context) error {
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
// @Param uid query string true "user id"
// @Param limit query string true "limit order"
// @Success 200 {object} []dtos.OrderSchema
// @Router /purchase/orders [GET]
func (p *Purchase) GetOrders(c echo.Context) error {
	sUserID := c.QueryParam("uid")
	if sUserID == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'uid' required",
		})
	}
	sLimit := c.QueryParam("limit")
	if sLimit == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'limit' required",
		})
	}

	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'uid' must be a positive integer",
		})
	}

	limit, err := strconv.ParseInt(sLimit, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'limit' must be a positive integer",
		})
	}

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
// @Param login body dtos.CreateOrderSchema true "order insert request"
// @Success 200 {object} dtos.OK
// @Router /purchase/orders [POST]
func (p *Purchase) CreateOrder(c echo.Context) error {
	var orderReq dtos.CreateOrderSchema
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
	msg, err := p.services.CreateOrders(c.Request().Context(), orderReq)
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
// @Param login body dtos.CartInsert true "cart insert request"
// @Success 200 {object} dtos.OK
// @Router /purchase/carts [POST]
func (p *Purchase) AddToCarts(c echo.Context) error {
	var cartReq dtos.CartInsert
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
	if err := p.services.AddToCart(c.Request().Context(), cartReq); err != nil {
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
// @Param uid query number true "user id"
// @Success 200 {object} dtos.CartSlices
// @Router /purchase/carts [GET]
func (p *Purchase) GetCarts(c echo.Context) error {
	sUserID := c.QueryParam("uid")
	if sUserID == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'uid' required",
		})
	}
	userID, err := strconv.ParseInt(sUserID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'uid' must be a positive integer",
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

// DeleteItem .
// @Description delete item from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Param id query int true "cart id"
// @Success 200 {object} dtos.OK
// @Router /purchase/carts [DELETE]
func (p *Purchase) DeleteItem(c echo.Context) error {
	cID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "'id' must be a positive integer",
		})
	}
	if err := p.
		services.DeleteItemFromCart(c.Request().Context(), cID); err != nil {
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
