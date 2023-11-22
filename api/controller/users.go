package controller

import (
	"net/http"

	"swclabs/swiftcart/internal/schema"
	"swclabs/swiftcart/internal/service"
	"swclabs/swiftcart/pkg/x/jwt"
	"swclabs/swiftcart/pkg/x/validator"

	"github.com/gin-gonic/gin"
)

// Login
// @Description Login account.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body schema.LoginRequest true "Login"
// @Success 200 {object} schema.LoginResponse
// @Router /v1/auth/login [POST]
func Login(c *gin.Context) {
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
	var account = service.NewAccountManagement()
	accessToken, err := account.Login(&request)
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
	var account = service.NewAccountManagement()
	if err := account.SignUp(&request); err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: "user data invalid",
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
// @Success 200 {object} schema.UserInfo
// @Router /v1/users [GET]
func GetMe(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	email, err := jwt.ParseToken(authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: "Invalid authorization header",
		})
		return
	}
	var account = service.NewAccountManagement()
	response, err := account.UserInfo(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

// UpdateUserInfo
// @Description update information for users.
// @Tags users
// @Accept json
// @Produce json
// @Param UserInfo body schema.UserUpdate true "Update User"
// @Success 200 {object} schema.OK
// @Router /v1/users [PUT]
func UpdateUserInfo(c *gin.Context) {
	var request schema.UserUpdate
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
	var user = service.NewAccountManagement()
	if err := user.UpdateUserInfo(&request); err != nil {
		c.JSON(http.StatusBadRequest, schema.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schema.OK{
		Msg: "update user information successfully",
	})

}
