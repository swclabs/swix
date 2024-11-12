// Package controller implements the controller interface
package healthcheck

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/service/healthcheck"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewController)

// NewController creates a new Base object
func NewController(services healthcheck.IService) IController {
	return &Controller{
		service: services,
	}
}

// IController interface for base controller
type IController interface {
	HealthCheck(c echo.Context) error
	WorkerCheck(c echo.Context) error
}

// Controller struct implementation of IBase
type Controller struct {
	service healthcheck.IService
}

// HealthCheck .
// @Description health check api server.
// @Tags base
// @Accept json
// @Produce json
// @Success 200
// @Router /status [GET]
func (b *Controller) HealthCheck(c echo.Context) error {
	return c.JSON(200, b.service.HealthCheck(c.Request().Context()))
}

// WorkerCheck .
// @Description health check worker consume server.
// @Tags base
// @Accept json
// @Produce json
// @Success 200
// @Router /worker [GET]
func (b *Controller) WorkerCheck(c echo.Context) error {
	results, err := healthcheck.UseTask(b.service).WorkerCheckResult(c.Request().Context(), 10)
	if err != nil {
		return c.JSON(400, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(200, dtos.OK{
		Msg: results,
	})
}

// Auth0Login .
// @Description Auth0 Login form.
// @Tags base
// @Accept json
// @Produce json
// @Success 200
// @Router /oauth2/login [GET]
// func Auth0Login(c echo.Context) error {
// 	auth := oauth2.New()
// 	url := auth.AuthCodeURL(auth.State)
// 	if err := session.Save(c, session.Base, "state", auth.State); err != nil {
// 		return c.JSON(http.StatusInternalServerError, dtos.Error{
// 			Msg: err.Error(),
// 		})
// 	}
// 	return c.Redirect(http.StatusTemporaryRedirect, url)
// }

// Auth0Callback .
// func Auth0Callback(c echo.Context) error {
// 	auth := oauth2.New()
// 	return auth.OAuth2CallBack(c)
// }
