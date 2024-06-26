package oauth2

import (
	"context"
	"fmt"
	"net/http"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/users"
	"swclabs/swipecore/internal/core/service/accountmanagement"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/lib/jwt"
	"swclabs/swipecore/pkg/lib/worker"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
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

	dbpool := db.GetPool()
	account := accountmanagement.New(
		users.New(dbpool),
		accounts.New(dbpool),
		addresses.New(dbpool),
		worker.NewClient(config.LoadEnv()),
	)
	if err := account.OAuth2SaveUser(
		context.TODO(),
		domain.OAuth2SaveUser{
			Email:     profile.Email,
			FirstName: profile.GivenName,
			LastName:  profile.FamilyName,
			Image:     profile.Picture,
		}); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	accessToken, err := jwt.GenerateToken(profile.Email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	if err := utils.SaveSession(ctx, utils.BaseSessions, "access_token", accessToken); err != nil {
		return ctx.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	// Redirect to logged in page.
	// return ctx.Redirect(http.StatusTemporaryRedirect, config.FeHomepage)
	return ctx.JSON(200, profile)
}
