// Package oauth2 implements the OAuth2
package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/users"
	"swclabs/swipecore/internal/core/service/manager"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/lib/crypto"

	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
)

// OAuth2CallBack handles the OAuth2 callback
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
		return ctx.JSON(http.StatusBadRequest, dtos.Error{
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

	var (
		dbpool = db.GetPool()
		blob   = blob.Connection()
	)
	account := manager.New(
		blob,
		users.New(dbpool),
		accounts.New(dbpool),
		addresses.New(dbpool),
	)
	if err := account.OAuth2SaveUser(
		context.TODO(),
		dtos.OAuth2SaveUser{
			Email:     profile.Email,
			FirstName: profile.GivenName,
			LastName:  profile.FamilyName,
			Image:     profile.Picture,
		}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	accessToken, err := crypto.GenerateToken(profile.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	if err := utils.SaveSession(ctx, utils.BaseSessions, "access_token", accessToken); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	// Redirect to logged in page.
	// return ctx.Redirect(http.StatusTemporaryRedirect, config.FeHomepage)
	return ctx.JSON(200, profile)
}
