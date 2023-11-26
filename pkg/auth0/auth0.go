package auth0

import (
	"context"
	"errors"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/utils"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type IAuth0 interface {
	VerifyIDToken(context.Context, *oauth2.Token) (*oidc.IDToken, error)
}

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
	State string
}

// New instantiates the *Authenticator.
func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+config.Auth0Domain+"/",
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     config.Auth0ClientID,
		ClientSecret: config.Auth0ClientSecret,
		RedirectURL:  config.Auth0CallbackUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		State:    utils.RandomString(10),
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (auth *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: auth.ClientID,
	}

	return auth.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
