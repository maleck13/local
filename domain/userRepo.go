package domain

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/maleck13/local/errors"
	r "gopkg.in/dancannon/gorethink.v2"
)

//Authorisor defines a Authorise api Service is an implementor
type Authorisor interface {
	Authorise(entity AccessDefinor, action string, actor Actor) error
}

//Actor is who is looking to act on an entity
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

type UserRepo struct {
	Config     *config.Config
	Actor      Actor
	Authorisor Authorisor
}

//NewUserRepo returns a configured UserRepo with it dependencies. This should be used when acting on behalf of a user
func NewUserRepo(conf *config.Config, actor Actor, authorisor Authorisor) UserRepo {
	return UserRepo{
		Config:     conf,
		Actor:      actor,
		Authorisor: authorisor,
	}
}

type User struct {
	*app.User
	ID           string `gorethink:"id,omitempty"`
	LoginExpires int    `json:"-"`
}

//AccessTypes implements AccessDefinor
func (u *User) AccessTypes() map[string][]string {
	access := map[string][]string{
		"read":  []string{"admin"},
		"write": []string{"admin"},
	}
	return access
}

//Owner implements AccessDefinor
func (u *User) Owner() string {
	return u.ID
}

//Id return the id of the user implements the Actor interface
func (u *User) Id() string {
	return u.ID
}

//Type return the type of the user implements the Actor interface
func (u *User) Type() string {
	return u.User.Type
}

func NewUserFromRequest(u *app.User) *User {
	return &User{User: u}
}

func (ur UserRepo) FindOneByFieldAndValue(field, val string) (*User, error) {
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return nil, errors.NewServiceError(err.Error(), 500)
	}
	res := &User{}
	q := map[string]interface{}{field: val}
	c, err := r.DB(data.DB_NAME).Table(data.USER_TABLE).Filter(q).Run(sess)
	if err != nil {
		return nil, errors.NewServiceError("failed to find a user "+err.Error(), 500)
	}
	if c.IsNil() {
		return nil, nil
	}
	c.One(res)
	if err := ur.Authorisor.Authorise(res, "read", ur.Actor); err != nil {
		return nil, err
	}
	return res, nil
}

func (ur UserRepo) SaveUpdate(u *User) error {
	if err := ur.Authorisor.Authorise(u, "write", ur.Actor); err != nil {
		return err
	}
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return errors.NewServiceError(err.Error(), 500)
	}
	var q = r.DB(data.DB_NAME).Table(data.USER_TABLE)
	fmt.Println("updating with id ", u.ID)
	if u.ID == "" {
		q = q.Insert(u)
	} else {
		q = q.Get(u.ID).Update(u)
	}
	_, err = q.RunWrite(sess)
	if err != nil {
		return errors.NewServiceError("failed to insert user "+err.Error(), 500)
	}
	return nil
}

func (ur UserRepo) DeleteByFieldAndValue(field string, value interface{}) error {
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return errors.NewServiceError(err.Error(), 500)
	}
	q := map[string]interface{}{field: value}
	if _, err := r.DB(data.DB_NAME).Table(data.USER_TABLE).Filter(q).Delete().Run(sess); err != nil {
		return errors.NewServiceError("failed to delete users "+err.Error(), 500)
	}
	return nil
}
