// Package controller account management implementation
package controller

import (
	"net/http"
	"swclabs/swipecore/internal/core/domain/dto"
	"swclabs/swipecore/internal/core/service/accountmanagement"
	"swclabs/swipecore/pkg/lib/valid"

	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
)

// IAccountManagement interface for account management
type IAccountManagement interface {
	Login(c echo.Context) error
	SignUp(c echo.Context) error
	Logout(c echo.Context) error
	GetMe(c echo.Context) error
	UpdateUserImage(c echo.Context) error
	CheckLoginEmail(c echo.Context) error
	UpdateUserInfo(c echo.Context) error
}

// AccountManagement struct implementation of IAccountManagement
type AccountManagement struct {
	Service accountmanagement.IAccountManagement
}

// NewAccountManagement creates a new AccountManagement object
func NewAccountManagement(services accountmanagement.IAccountManagement) IAccountManagement {
	return &AccountManagement{
		Service: services,
	}
}

// Login .
// @Description Login account.
// @Tags account_management
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "Login"
// @Success 200 {object} dto.LoginResponse
// @Router /auth/login [POST]
func (account *AccountManagement) Login(c echo.Context) error {
	var request dto.LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: _valid.Error(),
		})
	}
	// var account = service.New()
	accessToken, err := account.Service.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	// session := session.Default(c)
	// session.Set("access_token", accessToken)
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }

	if err := utils.SaveSession(c, utils.BaseSessions, "access_token", accessToken); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.LoginResponse{
		Success: true,
		Token:   accessToken,
		Email:   request.Email,
	})
}

// SignUp .
// @Description Register account for admin.
// @Tags account_management
// @Accept json
// @Produce json
// @Param sign_up body dto.SignUpRequest true "Sign Up"
// @Success 200 {object} dto.SignUpResponse
// @Router /auth/signup [POST]
func (account *AccountManagement) SignUp(c echo.Context) error {
	var request dto.SignUpRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: _valid.Error(),
		})
	}
	if err := account.Service.SignUp(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: "user data invalid",
		})
	}
	return c.JSON(http.StatusCreated, dto.SignUpResponse{
		Success: true,
		Msg:     "user has been created",
	})
}

// Logout .
// @Description logout user from the service
// @Tags account_management
// @Accept json
// @Produce json
// @Success 200 {object} dto.OK
// @Router /auth/logout [GET]
func (account *AccountManagement) Logout(c echo.Context) error {
	// session := sessions.Default(c)
	// session.Delete("access_token")
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }
	if err := utils.SaveSession(c, utils.BaseSessions, "access_token", ""); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.OK{
		Msg: "user logged out",
	})
}

// GetMe .
// @Description get information for users.
// @Tags account_management
// @Accept json
// @Produce json
// @Success 200 {object} model.Users
// @Router /users [GET]
func (account *AccountManagement) GetMe(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	email := utils.Session(c, utils.BaseSessions, "email").(string)
	response, err := account.Service.UserInfo(c.Request().Context(), email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateUserInfo .
// @Description update information for users.
// @Tags account_management
// @Accept json
// @Produce json
// @Param UserSchema body dto.User true "Update Users"
// @Success 200 {object} dto.OK
// @Router /users [PUT]
func (account *AccountManagement) UpdateUserInfo(c echo.Context) error {
	var request dto.User
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: _valid.Error(),
		})
	}
	if err := account.Service.UpdateUserInfo(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.OK{
		Msg: "update user information successfully",
	})

}

// UpdateUserImage .
// @Description update information for users.
// @Tags account_management
// @Accept json
// @Produce json
// @Param img formData file true "image of collections"
// @Success 200 {object} dto.OK
// @Router /users/image [PUT]
func (account *AccountManagement) UpdateUserImage(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	email := utils.Session(c, utils.BaseSessions, "email").(string)
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: err.Error(),
		})
	}
	if err := account.Service.UploadAvatar(email, file); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.OK{
		Msg: "update user images successfully",
	})
}

// CheckLoginEmail .
// @Description check email address before login
// @Tags account_management
// @Accept json
// @Produce json
// @Param email query string true "email address"
// @Success 200 {object} dto.OK
// @Router /auth [GET]
func (account *AccountManagement) CheckLoginEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, dto.Error{
			Msg: "missing query parameter: email",
		})
	}
	if err := account.Service.CheckLoginEmail(c.Request().Context(), email); err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.OK{
		Msg: "email: " + email,
	})
}
