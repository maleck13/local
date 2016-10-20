package domain

// holds the common types for repos

import jwt "github.com/dgrijalva/jwt-go"

//Authorisor defines a Authorise api Service is an implementor
type Authorisor interface {
	Authorise(entity AccessDefinor, action string, actor Actor) error
}

//Actor represents what / who is looking to act on an entity
type Actor interface {
	Id() string
	Type() string
}

// AccessDefinor defines the inteface for determining the access an actor has
type AccessDefinor interface {
	AccessTypes() map[string][]string
	Owner() string
}

//NewLocalActor returns an Local Actor based on a jwtToken
func NewLocalActor(t *jwt.Token) Actor {
	return localActor{
		t,
	}
}

type localActor struct {
	*jwt.Token
}

//Id return the userId of the acting user
func (aa localActor) Id() string {
	if nil == aa.Token {
		return ""
	}
	if claims, ok := aa.Token.Claims.(jwt.MapClaims); ok {
		return claims["id"].(string)
	}
	return ""
}

//Type return the type of the acting user
func (aa localActor) Type() string {
	if nil == aa.Token {
		return ""
	}
	if claims, ok := aa.Token.Claims.(jwt.MapClaims); ok {
		return claims["type"].(string)
	}
	return ""
}

type adminActor struct {
}

func (aa adminActor) Id() string {
	return "admin"
}
func (aa adminActor) Type() string {
	return "admin"
}

type AdminAccessDefinor struct {
}

func (AdminAccessDefinor) AccessTypes() map[string][]string {
	return map[string][]string{
		"admin": []string{"admin"},
	}
}

func (AdminAccessDefinor) Owner() string {
	return "admin"
}

func NewAdminActor() Actor {
	return adminActor{}
}
