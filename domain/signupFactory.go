package domain

import (
	"net/http"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	e "github.com/maleck13/local/errors"
)

//Registerer defines methods for registration
type Registerer interface {
	Register(*app.User) (*User, error)
}

//NewSignUpFactory returns a instance of SignUpFactory
func NewSignUpFactory(conf *config.Config) *SignUpFactory {
	return &SignUpFactory{
		Config: conf,
	}
}

type SignUpFactory struct {
	Config *config.Config
}

//Factory depending on the signup type it will return the correct Registerer
func (sf *SignUpFactory) Factory(signUpType string) (Registerer, error) {
	if signUpType == "google" {
		return NewGoogleAPI(sf.Config), nil
	}
	return nil, e.NewServiceError("unknown sign up type ", http.StatusBadRequest)
}
