package manager

import (
	"fmt"
	"net/http"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/manager"
	"swclabs/swix/pkg/lib/session"
	"swclabs/swix/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewController)

// NewController creates a new Manager object
func NewController(services manager.IManager) IController {
	return &Controller{
		service: services,
	}
}

// IController interface for manager
type IController interface {
	Login(c echo.Context) error
	SignUp(c echo.Context) error
	Logout(c echo.Context) error
	GetMe(c echo.Context) error
	Auth(c echo.Context) error
	UpdateUserImage(c echo.Context) error
	CheckLoginEmail(c echo.Context) error
	UpdateUserInfo(c echo.Context) error
}

// Controller struct implementation of IManager
type Controller struct {
	service manager.IManager
}

// Auth implements IManager.
func (manager *Controller) Auth(c echo.Context) error {
	var (
		email    = c.FormValue("email")
		password = c.FormValue("password")
	)
	accessToken, err := manager.service.Login(c.Request().Context(), dtos.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("error login: %v, %s, %s", err, email, password))
	}
	if err := session.Save(c, session.Base, "access_token", accessToken); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.Redirect(http.StatusSeeOther, "/docs/index.html")
}

// Login .
// @Description Login account.
// @Tags manager
// @Accept json
// @Produce json
// @Param login body dtos.LoginRequest true "Login"
// @Success 200 {object} dtos.LoginResponse
// @Router /auth/login [POST]
func (manager *Controller) Login(c echo.Context) error {
	var request dtos.LoginRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	// var account = service.New()
	accessToken, err := manager.service.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	// session := session.Default(c)
	// session.Set("access_token", accessToken)
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }

	if err := session.Save(c, session.Base, "access_token", accessToken); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtos.LoginResponse{
		Success: true,
		Token:   accessToken,
		Email:   request.Email,
	})
}

// SignUp .
// @Description Register account for admin.
// @Tags manager
// @Accept json
// @Produce json
// @Param sign_up body dtos.SignUpRequest true "Sign Up"
// @Success 200 {object} dtos.SignUpResponse
// @Router /auth/signup [POST]
func (manager *Controller) SignUp(c echo.Context) error {
	var request dtos.SignUpRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	if err := manager.service.SignUp(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: "user data invalid",
		})
	}
	return c.JSON(http.StatusCreated, dtos.SignUpResponse{
		Success: true,
		Msg:     "user has been created",
	})
}

// Logout .
// @Description logout user from the service
// @Tags manager
// @Accept json
// @Produce json
// @Success 200 {object} dtos.OK
// @Router /auth/logout [GET]
func (manager *Controller) Logout(c echo.Context) error {
	// session := sessions.Default(c)
	// session.Delete("access_token")
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }
	if err := session.Save(c, session.Base, "access_token", ""); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "user logged out",
	})
}

// GetMe .
// @Description get information for users.
// @Tags manager
// @Accept json
// @Produce json
// @Success 200 {object} model.Users
// @Router /users [GET]
func (manager *Controller) GetMe(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	email := session.Get(c, session.Base, "email")
	response, err := manager.service.UserInfo(c.Request().Context(), email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateUserInfo .
// @Description update information for users.
// @Tags manager
// @Accept json
// @Produce json
// @Param UserSchema body dtos.UserUpdate true "Update Users"
// @Success 200 {object} dtos.OK
// @Router /users [PUT]
func (manager *Controller) UpdateUserInfo(c echo.Context) error {
	var request dtos.UserUpdate
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if _valid := valid.Validate(&request); _valid != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: _valid.Error(),
		})
	}
	if err := manager.service.UpdateUserInfo(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "update user information successfully",
	})

}

// UpdateUserImage .
// @Description update information for users.
// @Tags manager
// @Accept json
// @Produce json
// @Param img formData file true "image of collections"
// @Success 200 {object} dtos.OK
// @Router /users/image [PUT]
func (manager *Controller) UpdateUserImage(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	email := session.Get(c, session.Base, "email")
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := manager.service.UploadAvatar(email, file); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "update user images successfully",
	})
}

// CheckLoginEmail .
// @Description check email address before login
// @Tags manager
// @Accept json
// @Produce json
// @Param email query string true "email address"
// @Success 200 {object} dtos.OK
// @Router /auth/email [GET]
func (manager *Controller) CheckLoginEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing query parameter: email",
		})
	}
	if err := manager.service.CheckLoginEmail(c.Request().Context(), email); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "email: " + email,
	})
}
