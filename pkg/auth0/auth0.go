package auth0

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type IAuth0 interface {
	VerifyIDToken(context.Context, *oauth2.Token) (*oidc.IDToken, error)
}
