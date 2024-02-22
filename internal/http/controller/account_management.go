package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/service"
	"github.com/swclabs/swipe-api/pkg/tools"
	"github.com/swclabs/swipe-api/pkg/utils"
)

type AccountManagement struct {
	service domain.IAccountManagementService
}

func NewAccountManagement() IAccountManagement {
	return &AccountManagement{
		service: service.NewAccountManagement(),
	}
}

type IAccountManagement interface {
	Login(c echo.Context) error
	SignUp(c echo.Context) error
	Logout(c echo.Context) error
	GetMe(c echo.Context) error
	UpdateUserImage(c echo.Context) error
	CheckLoginEmail(c echo.Context) error
	UpdateUserInfo(c echo.Context) error
}

// Login
// @Description Login account.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "Login"
// @Success 200 {object} domain.LoginResponse
// @Router /auth/login [POST]
func (account *AccountManagement) Login(c echo.Context) error {
	var request domain.LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if _valid := tools.Validate(request); _valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
	}
	// var account = service.NewAccountManagement()
	accessToken, err := account.service.Login(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	// session := session.Default(c)
	// session.Set("access_token", accessToken)
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }

	if err := utils.SaveSession(c, utils.BaseSessions, "access_token", accessToken); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.LoginResponse{
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
// @Router /auth/signup [POST]
func (account *AccountManagement) SignUp(c echo.Context) error {
	var request domain.SignUpRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if _valid := tools.Validate(request); _valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
	}
	if err := account.service.SignUp(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "user data invalid",
		})
	}
	return c.JSON(http.StatusCreated, domain.SignUpResponse{
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
// @Router /auth/logout [GET]
func (account *AccountManagement) Logout(c echo.Context) error {
	// session := sessions.Default(c)
	// session.Delete("access_token")
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }
	if err := utils.SaveSession(c, utils.BaseSessions, "access_token", ""); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.OK{
		Msg: "user logged out",
	})
}

// GetMe
// @Description get information for users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} domain.UserInfo
// @Router /users [GET]
func (account *AccountManagement) GetMe(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	email := utils.Session(c, utils.BaseSessions, "email").(string)
	response, err := account.service.UserInfo(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateUserInfo
// @Description update information for users.
// @Tags users
// @Accept json
// @Produce json
// @Param UserInfo body domain.UserUpdate true "Update User"
// @Success 200 {object} domain.OK
// @Router /users [PUT]
func (account *AccountManagement) UpdateUserInfo(c echo.Context) error {
	var request domain.UserUpdate
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if _valid := tools.Validate(request); _valid != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: _valid,
		})
	}
	if err := account.service.UpdateUserInfo(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.OK{
		Msg: "update user information successfully",
	})

}

// UpdateUserImage
// @Description update information for users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} domain.OK
// @Router /users/image [PUT]
func (account *AccountManagement) UpdateUserImage(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	email := utils.Session(c, utils.BaseSessions, "email").(string)
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if err := account.service.UploadAvatar(email, file); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.OK{
		Msg: "update user images successfully",
	})
}

// CheckLoginEmail
// @Description check email address before login
// @Tags auth
// @Accept json
// @Produce json
// @Param email query string true "email address"
// @Success 200 {object} domain.OK
// @Router /auth [GET]
func (account *AccountManagement) CheckLoginEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing query parameter: email",
		})
	}
	if err := account.service.CheckLoginEmail(email); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, domain.OK{
		Msg: "email: " + email,
	})
}
