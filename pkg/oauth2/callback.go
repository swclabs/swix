package oauth2

import (
	"net/http"

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (auth *Authenticator) OAuth2CallBack(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if ctx.Query("state") != session.Get("state").(string) {
		ctx.String(http.StatusBadRequest, "Invalid state parameter. %s", session.Get("state"))
		return
	}

	// Exchange an authorization code for a token.
	token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
		return
	}

	profile, err := auth.VerifyToken(token)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	account := service.NewAccountManagement()
	if err := account.OAuth2SaveUser(&domain.OAuth2SaveUser{
		Email:     profile.Email,
		FirstName: profile.GivenName,
		LastName:  profile.FamilyName,
		Image:     profile.Picture,
	}); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	session.Set("access_token", token.AccessToken)
	// session.Set("profile", profile)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Redirect to logged in page.
	ctx.Redirect(http.StatusTemporaryRedirect, config.FeHomepage)
	// ctx.JSON(200, profile)
}
