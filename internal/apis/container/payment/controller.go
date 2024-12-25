package payment

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	pm "github.com/swclabs/swipex/internal/core/service/payment"
	"github.com/swclabs/swipex/pkg/gen/payment"
)

var _ = app.Controller(NewController)

// NewController creates a new Article object
func NewController(service *pm.Payment) IController {
	return &Controller{
		Services: service,
	}
}

// IController interface for article controller
type IController interface {
	Status(c echo.Context) error
	Payment(c echo.Context) error
	PaymentReturn(c echo.Context) error
}

// Controller struct implementation of IArticle
type Controller struct {
	Services *pm.Payment
}

// Status .
// @Description check payment service status
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body payment.PaymentRequest true "payment request"
// @Success 200 {object} payment.PaymentResponse
// @Router /payment [POST]
func (pmc *Controller) Payment(c echo.Context) error {
	var req payment.PaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}

	req.IpAddress = c.RealIP()

	resp, err := pmc.Services.ProcessPayment(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

// PaymentReturn implements IController.
func (pmc *Controller) PaymentReturn(c echo.Context) error {
	panic("unimplemented")
}

// Status .
// @Description check payment service status
// @Tags payment
// @Accept json
// @Produce json
// @Success 200 {object} payment.StatusResponse
// @Router /payment/status [GET]
func (pmc *Controller) Status(c echo.Context) error {
	resp, err := pmc.Services.CheckStatus(c.Request().Context(), &payment.StatusRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}
