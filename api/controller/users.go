package controller

import (
	"example/swiftcart/internal/schema"
	"example/swiftcart/internal/service"
	"example/swiftcart/pkg/lib/validator"
	"example/swiftcart/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login
// @Description Login account.
// @Tags auth
// @Accept json
// @Produce json
// @Param sign_up body schema.LoginRequest true "Login"
// @Success 200 {object} schema.LoginResponse
// @Router /v1/auth/login [POST]
func Login(c *gin.Context) {
	var users = service.NewUsers()
	var request schema.LoginRequest
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
	accessToken, err := users.Login(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schema.LoginResponse{
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
// @Router /v1/auth/signup [POST]
func SignUp(c *gin.Context) {
	var users = service.NewUsers()
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
	if err := users.SignUp(&request); err != nil {
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

// GetMe
// @Description get information for users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} schema.InforResponse
// @Router /v1/users [GET]
func GetMe(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	email, err := utils.ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: "Invalid authorization header",
		})
		return
	}
	var users = service.NewUsers()
	response, err := users.Infor(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
