package domain

import (
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/maleck13/local/errors"
	r "gopkg.in/dancannon/gorethink.v2"
)

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

func (ur UserRepo) Save(u *User) error {
	if err := ur.Authorisor.Authorise(u, "write", ur.Actor); err != nil {
		return err
	}
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return errors.NewServiceError(err.Error(), 500)
	}
	_, err = r.DB(data.DB_NAME).Table(data.USER_TABLE).Insert(u).RunWrite(sess)
	if err != nil {
		return errors.NewServiceError("failed to insert councillor "+err.Error(), 500)
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
