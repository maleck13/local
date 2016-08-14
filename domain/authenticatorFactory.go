package domain

import (
	"net/http"

	e "github.com/maleck13/local/errors"
	"github.com/maleck13/local/config"
)

type Authenticator interface {
	Authenticate(token, email string) error
}

type AuthenticatorFactory struct {
	Config *config.Config
}

func NewAuthenticatorFactory(conf *config.Config) *AuthenticatorFactory {
	return &AuthenticatorFactory{
		Config:conf,
	}
}


func (af *AuthenticatorFactory)Factory(authType string) (Authenticator, error) {
	if authType == "google" {
		return NewGoogleApi(af.Config), nil
	}
	return nil, e.NewServiceError("unknown authentication type "+authType, http.StatusUnauthorized)
}
