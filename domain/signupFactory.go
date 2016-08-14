package domain

import (
	"github.com/maleck13/local/app"
	e "github.com/maleck13/local/errors"
	"net/http"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
)

type Registerer interface {
	Register(*app.User) (*data.User, error)
}

func NewSignUpFactory(conf *config.Config)*SignUpFactory{
	return &SignUpFactory{
		Config :conf,
	}
}

type SignUpFactory struct {
	Config *config.Config
}

func (sf *SignUpFactory) Factory(signUpType string) (Registerer, error) {
	if signUpType == "google" {
		return NewGoogleApi(sf.Config), nil
	}
	return nil, e.NewServiceError("unknown sign up type ", http.StatusBadRequest)
}
