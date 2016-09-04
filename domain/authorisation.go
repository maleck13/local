package domain

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maleck13/local/errors"
)

//Authorisor defines a Authorise api
type Authorisor interface {
	Authorise(entity AccessDefinor, action string, actor Actor) error
}

//Actor is who is looking to act on an entity
type Actor interface {
	Id() string
	Type() string
}

type AccessDefinor interface {
	AccessTypes() map[string][]string
	Owner() string
}

//NewAuthActor returns an AuthActor based on a jwtToken
func NewAuthActor(t *jwt.Token) AuthActor {
	return AuthActor{
		t,
	}
}

type AuthActor struct {
	*jwt.Token
}

type AdminActor struct {
}

func (aa AdminActor) Id() string {
	return "admin"
}
func (aa AdminActor) Type() string {
	return "admin"
}

type AdminAccessDefinor struct {
}

func (Aad AdminAccessDefinor) AccessTypes() map[string][]string {
	return map[string][]string{
		"admin": []string{"admin"},
	}
}

func (Aad AdminAccessDefinor) Owner() string {
	return "admin"
}

//Id return the userId of the acting user
func (aa AuthActor) Id() string {
	if nil == aa.Token {
		return ""
	}
	if claims, ok := aa.Token.Claims.(jwt.MapClaims); ok {
		return claims["id"].(string)
	}
	return ""
}

//Type return the type of the acting user
func (aa AuthActor) Type() string {
	if nil == aa.Token {
		return ""
	}
	if claims, ok := aa.Token.Claims.(jwt.MapClaims); ok {
		return claims["type"].(string)
	}
	return ""
}

type Access struct{}

//Authorise an action on an entity
func (auth Access) Authorise(entity AccessDefinor, action string, actor Actor) error {
	access := entity.AccessTypes()
	if actor.Id() == entity.Owner() {
		return nil
	}
	var who []string
	if action == "write" {
		who = access["write"]

	} else if action == "read" {
		who = access["read"]
	}
	for _, allowed := range who {
		if actor.Type() == allowed || "any" == allowed {
			return nil
		}
	}
	return errors.NewServiceError("unauthorised access", http.StatusUnauthorized)
}
