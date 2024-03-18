package oauth2

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/utils"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type I0Auth2 interface {
	VerifyToken(*oauth2.Token) (*GoogleOAuth2, error)
	VerifyTokenByte(token *oauth2.Token) ([]byte, error)
	OAuth2CallBack(ctx *gin.Context)
}

type GoogleOAuth2 struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
}

// Authenticator is used to authenticate our users.
type Authenticator struct {
	oauth2.Config
	State string
}

// New instantiates the *Authenticator.
func New() *Authenticator {
	conf := oauth2.Config{
		ClientID:     config.Auth0ClientID,
		ClientSecret: config.Auth0ClientSecret,
		RedirectURL:  config.Auth0CallbackUrl,
		Endpoint:     google.Endpoint,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return &Authenticator{
		Config: conf,
		State:  utils.RandomString(10),
	}
}

func (auth *Authenticator) VerifyToken(token *oauth2.Token) (*GoogleOAuth2, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var profile GoogleOAuth2
	if err := json.Unmarshal(body, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (auth *Authenticator) VerifyTokenByte(token *oauth2.Token) ([]byte, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
