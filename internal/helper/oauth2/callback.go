package oauth2

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/service"
	"github.com/swclabs/swipe-api/pkg/utils"
)

func (auth *Authenticator) OAuth2CallBack(ctx echo.Context) error {
	// session := sessions.Default(ctx)
	// if ctx.Query("state") != session.Get("state").(string) {
	// 	ctx.String(http.StatusBadRequest, "Invalid state parameter. %s", session.Get("state"))
	// 	return
	// }

	type Query struct {
		State string `query:"state"`
		Code  string `query:"code"`
	}
	var query Query

	if err := ctx.Bind(&query); err != nil {
		return ctx.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}

	state := utils.Session(ctx, utils.BaseSessions, "state").(string)
	if state != query.State {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Invalid state parameter. %s", state))
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(ctx.Request().Context(), query.Code)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
	}

	profile, err := auth.VerifyToken(token)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	account := service.NewAccountManagement()
	if err := account.OAuth2SaveUser(&domain.OAuth2SaveUser{
		Email:     profile.Email,
		FirstName: profile.GivenName,
		LastName:  profile.FamilyName,
		Image:     profile.Picture,
	}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	// session.Set("access_token", token.AccessToken)
	// // session.Set("profile", profile)
	// if err := session.Save(); err != nil {
	// 	return ctx.String(http.StatusInternalServerError, err.Error())
	// }
	if err := utils.SaveSession(ctx, utils.BaseSessions, "access_token", token.AccessToken); err != nil {
		return ctx.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	// Redirect to logged in page.
	return ctx.Redirect(http.StatusTemporaryRedirect, config.FeHomepage)
	// ctx.JSON(200, profile)
}
