package domain

import (
	"net/http"

	"github.com/maleck13/local/config"
	e "github.com/maleck13/local/errors"
	"github.com/dgrijalva/jwt-go"
)



type Authenticator interface {
	Authenticate(token, email string) (*jwt.Token,error)
}

type AuthenticatorFactory struct {
	Config *config.Config
}

type JwTTokenHandler interface {
	Validate(token string)error
	Refresh()
	CreateToken()(*jwt.Token,error)
}

type Jwt struct {
	Config *config.Config
}

func (j *Jwt)Validate(token string)error  {
	jht.Parse()
}

func (j *Jwt)CreateToken()(*jwt.Token,error)  {
	t := jwt.New(jwt.SigningMethodHS256)

}

func (j *Jwt)Refresh()  {

}

func NewAuthenticatorFactory(conf *config.Config) *AuthenticatorFactory {
	return &AuthenticatorFactory{
		Config: conf,
	}
}

func (af *AuthenticatorFactory) Factory(authType string) (Authenticator, error) {
	if authType == "google" {
		return NewGoogleApi(af.Config), nil
	}
	return nil, e.NewServiceError("unknown authentication type "+authType, http.StatusUnauthorized)
}
