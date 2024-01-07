package controller

import (
	"net/http"
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/internal/service"
	"swclabs/swiftcart/pkg/oauth2"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HealthCheck .
// @Description health check app server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/common/healthcheck [GET]
func HealthCheck(c *gin.Context) {
	c.JSON(200, service.HealthCheck())
}

// Auth0Login .
// @Description Auth0 Login form.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/oauth2/login [GET]
func Auth0Login(c *gin.Context) {
	auth := oauth2.New()
	url := auth.AuthCodeURL(auth.State)
	session := sessions.Default(c)
	session.Set("state", auth.State)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func Auth0Callback(c *gin.Context) {
	auth := oauth2.New()
	auth.OAuth2CallBack(c)
}

func WorkerCheck(c *gin.Context) {
	if err := service.WorkerCheck(); err != nil {
		c.JSON(400, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(200, service.HealthCheck())
}

func Foo(ctx *gin.Context) {
	// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
	// in the middleware before, it will be sent through as well
	panic("y tho")
}
