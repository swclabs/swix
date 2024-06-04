package controller

import (
	"net/http"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/core/utils/oauth2"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
)

// HealthCheck .
// @Description health check api server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Register /common/healthcheck [GET]
func HealthCheck(c echo.Context) error {
	common := common.New()
	return c.JSON(200, common.HealthCheck(c.Request().Context()))
}

// Auth0Login .
// @Description Auth0 Login form.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Register /oauth2/login [GET]
func Auth0Login(c echo.Context) error {
	auth := oauth2.New()
	url := auth.AuthCodeURL(auth.State)
	if err := utils.SaveSession(c, utils.BaseSessions, "state", auth.State); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Auth0Callback(c echo.Context) error {
	auth := oauth2.New()
	return auth.OAuth2CallBack(c)
}

// WorkerCheck .
// @Description health check worker consume server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Register /common/worker [GET]
func WorkerCheck(c echo.Context) error {
	common := common.New()
	results, err := common.DelayWorkerCheckResult(c.Request().Context())
	if err != nil {
		return c.JSON(400, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(200, domain.OK{
		Msg: results,
	})
}

func Foo(ctx echo.Context) error {
	// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
	// in the middleware before, it will be sent through as well
	panic("y tho")
}
