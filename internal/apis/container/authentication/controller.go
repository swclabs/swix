package authentication

import (
	"fmt"
	"net/http"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/service/authentication"
	"swclabs/swipex/internal/core/x/oauth2"
	"swclabs/swipex/pkg/lib/crypto"
	"swclabs/swipex/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewController)

// NewController creates a new Authentication object
func NewController(services authentication.IAuthentication) IController {
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
	OAuth2(c echo.Context) error
}

// Controller struct implementation of IAuthentication
type Controller struct {
	service authentication.IAuthentication
}

// Auth implements IAuthentication.
func (auth *Controller) Auth(c echo.Context) error {
	var (
		email    = c.FormValue("email")
		password = c.FormValue("password")
	)
	_, err := auth.service.Login(c.Request().Context(), dtos.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("error login: %v, %s, %s", err, email, password))
	}
	// if err := session.Save(c, session.Base, "access_token", accessToken); err != nil {
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }
	return c.Redirect(http.StatusSeeOther, "/docs/index.html")
}

// Login .
// @Description Login account.
// @Tags authentication
// @Accept json
// @Produce json
// @Param login body dtos.LoginRequest true "Login"
// @Success 200 {object} dtos.LoginResponse
// @Router /auth/login [POST]
func (auth *Controller) Login(c echo.Context) error {
	_, email, _ := crypto.Authenticate(c)
	fmt.Println(email)
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
	accessToken, err := auth.service.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	// session := session.Default(c)
	// session.Set("access_token", accessToken)
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }

	// if err := session.Save(c, session.Base, "access_token", accessToken); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dtos.Error{
	// 		Msg: err.Error(),
	// 	})
	// }

	return c.JSON(http.StatusOK, dtos.LoginResponse{
		Success: true,
		Token:   accessToken,
		Email:   request.Email,
	})
}

// SignUp .
// @Description Register account for admin.
// @Tags authentication
// @Accept json
// @Produce json
// @Param sign_up body dtos.SignUpRequest true "Sign Up"
// @Success 200 {object} dtos.SignUpResponse
// @Router /auth/signup [POST]
func (auth *Controller) SignUp(c echo.Context) error {
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
	if err := auth.service.SignUp(c.Request().Context(), request); err != nil {
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
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {object} dtos.OK
// @Router /auth/logout [GET]
func (auth *Controller) Logout(c echo.Context) error {
	// session := sessions.Default(c)
	// session.Delete("access_token")
	// if err := session.Save(); err != nil {
	// 	return c.String(http.StatusInternalServerError, err.Error())
	// }
	// if err := session.Save(c, session.Base, "access_token", ""); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, dtos.Error{
	// 		Msg: err.Error(),
	// 	})
	// }
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "user logged out",
	})
}

// GetMe .
// @Description get information for users.
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {object} model.Users
// @Router /users [GET]
func (auth *Controller) GetMe(c echo.Context) error {
	_, email, _ := crypto.Authenticate(c)
	response, err := auth.service.UserInfo(c.Request().Context(), email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateUserInfo .
// @Description update information for users.
// @Tags authentication
// @Accept json
// @Produce json
// @Param UserSchema body dtos.UserUpdate true "Update Users"
// @Success 200 {object} dtos.OK
// @Router /users [PUT]
func (auth *Controller) UpdateUserInfo(c echo.Context) error {
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
	if err := auth.service.UpdateUserInfo(c.Request().Context(), request); err != nil {
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
// @Tags authentication
// @Accept json
// @Produce json
// @Param img formData file true "image of collections"
// @Success 200 {object} dtos.OK
// @Router /users/image [PUT]
func (auth *Controller) UpdateUserImage(c echo.Context) error {
	// session := sessions.Default(c)
	// email := session.Get("email").(string)
	_, email, _ := crypto.Authenticate(c)
	file, err := c.FormFile("img")
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := auth.service.UploadAvatar(email, file); err != nil {
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
// @Tags authentication
// @Accept json
// @Produce json
// @Param email query string true "email address"
// @Success 200 {object} dtos.OK
// @Router /auth/email [GET]
func (auth *Controller) CheckLoginEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing query parameter: email",
		})
	}
	if err := auth.service.CheckLoginEmail(c.Request().Context(), email); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.OK{
			Msg: "email invalid " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.OK{
		Msg: "email valid ",
	})
}

// OAuth2 .
// @Description Auth0 verify token.
// @Tags authentication
// @Accept json
// @Produce json
// @Param access_token query string true "google access token"
// @Success 200
// @Router /oauth2/google [GET]
func (auth *Controller) OAuth2(c echo.Context) error {
	oauth2 := oauth2.New()
	accessToken := c.QueryParam("access_token")
	profile, err := oauth2.VerifyAccessToken(accessToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	userID, err := auth.service.OAuth2SaveUser(c.Request().Context(), dtos.OAuth2SaveUser{
		Email:     profile.Email,
		FirstName: profile.Name,
		LastName:  profile.FamilyName,
		Image:     profile.Picture,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	accessToken, err = crypto.GenerateToken(userID, profile.Email, "customer")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dtos.LoginResponse{
		Success: true,
		Token:   accessToken,
		Email:   profile.Email,
	})
}
