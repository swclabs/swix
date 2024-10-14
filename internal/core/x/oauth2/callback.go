// Package oauth2 implements the OAuth2
package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/repos/accounts"
	"swclabs/swix/internal/core/repos/addresses"
	"swclabs/swix/internal/core/repos/users"
	"swclabs/swix/internal/core/service/manager"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/db"
	"swclabs/swix/pkg/lib/crypto"
	"swclabs/swix/pkg/lib/session"

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

	state := session.Get(ctx, session.Base, "state")
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
	manager := manager.New(
		blob,
		users.New(dbpool),
		accounts.New(dbpool),
		addresses.New(dbpool),
	)
	ID, err := manager.OAuth2SaveUser(context.TODO(),
		dtos.OAuth2SaveUser{
			Email:     profile.Email,
			FirstName: profile.GivenName,
			LastName:  profile.FamilyName,
			Image:     profile.Picture,
		},
	)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	accessToken, err := crypto.GenerateToken(ID, profile.Email, "Customer")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	if err := session.Save(ctx, session.Base, "access_token", accessToken); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}

	// Redirect to logged in page.
	// return ctx.Redirect(http.StatusTemporaryRedirect, config.FeHomepage)
	return ctx.JSON(200, profile)
}
