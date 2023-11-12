package controller

import (
	"example/komposervice/internal/schema"
	"example/komposervice/internal/service"
	"example/komposervice/pkg/lib/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

var authService = service.NewAuth()

// SignIn
// @Description Sign IN account.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_up body schema.SignInRequest true "Sign In"
// @Success 200 {object} schema.SignInResponse
// @Router /v1/auth/sign-in [POST]
func SignIn(c *gin.Context) {
	var request schema.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	if _valid := validator.Validate(request); _valid != "" {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: _valid,
		})
		return
	}
	accessToken, err := authService.SignIn(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schema.SignInResponse{
		Success: true,
		Token:   accessToken,
		Email:   request.Email,
	})
}

// SignUp
// @Description Register account for admin.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_up body schema.SignUpRequest true "Sign Up"
// @Success 200 {object} schema.SignUpResponse
// @Router /v1/auth/sign-up [POST]
func SignUp(c *gin.Context) {
	var request schema.SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	if _valid := validator.Validate(request); _valid != "" {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: _valid,
		})
		return
	}
	if err := authService.SignUp(request); err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, schema.SignUpResponse{
		Success: true,
		Msg:     "user has been created",
	})
}
