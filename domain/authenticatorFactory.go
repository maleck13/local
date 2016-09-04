package domain

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/maleck13/local/config"
	e "github.com/maleck13/local/errors"
)

var errValidationFailed = goa.NewErrorClass("validation_failed", 401)

//Authenticator defines a user authenticator interface
type Authenticator interface {
	Authenticate(token, email string) (*User, error)
}

//AuthenticatorFactory decides which type of authenticator to return
type AuthenticatorFactory struct {
	Config *config.Config
}

//Factory is the method to be called to get your Authenticator
func (af *AuthenticatorFactory) Factory(authType string) (Authenticator, error) {
	if authType == "google" {
		return NewGoogleAPI(af.Config), nil
	}
	return nil, e.NewServiceError("unknown authentication type "+authType, http.StatusUnauthorized)
}
