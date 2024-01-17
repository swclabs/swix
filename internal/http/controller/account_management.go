package controller

import (
	"net/http"

	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/service"
	"github.com/swclabs/swipe-api/pkg/validator"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AccountManagement struct {
	service domain.IAccountManagementService
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		service: service.NewAccountManagement(),
	}
}

// Login
// @Description Login account.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "Login"
// @Success 200 {object} domain.LoginResponse
// @Router /v1/auth/login [POST]
func (account *AccountManagement) Login(c *gin.Context) {
	var request domain.LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	if _valid := validator.Validate(request); _valid != "" {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
		return
	}
	// var account = service.NewAccountManagement()
	accessToken, err := account.service.Login(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	session := sessions.Default(c)
	session.Set("access_token", accessToken)
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.LoginResponse{
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
// @Param sign_up body domain.SignUpRequest true "Sign Up"
// @Success 200 {object} domain.SignUpResponse
// @Router /v1/auth/signup [POST]
func (account *AccountManagement) SignUp(c *gin.Context) {
	var request domain.SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	if _valid := validator.Validate(request); _valid != "" {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
		return
	}
	if err := account.service.SignUp(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "user data invalid",
		})
		return
	}
	c.JSON(http.StatusCreated, domain.SignUpResponse{
		Success: true,
		Msg:     "user has been created",
	})
}

// Logout
// @Description logout user from the service
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} domain.OK
// @Router /v1/auth/logout [GET]
func (account *AccountManagement) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("access_token")
	if err := session.Save(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.OK{
		Msg: "user logged out",
	})
}

// GetMe
// @Description get information for users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} domain.UserInfo
// @Router /v1/users [GET]
func (account *AccountManagement) GetMe(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("email").(string)
	response, err := account.service.UserInfo(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
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
// @Param UserInfo body domain.UserUpdate true "Update User"
// @Success 200 {object} domain.OK
// @Router /v1/users [PUT]
func (account *AccountManagement) UpdateUserInfo(c *gin.Context) {
	var request domain.UserUpdate
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	if _valid := validator.Validate(request); _valid != "" {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
		return
	}
	if err := account.service.UpdateUserInfo(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.OK{
		Msg: "update user information successfully",
	})

}

// UpdateUserImage
// @Description update information for users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} domain.OK
// @Router /v1/users/image [PUT]
func (account *AccountManagement) UpdateUserImage(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("email").(string)
	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	if err := account.service.UploadAvatar(email, file); err != nil {
		c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.OK{
		Msg: "update user images successfully",
	})
}
