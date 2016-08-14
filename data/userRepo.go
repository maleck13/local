package data

import (
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/errors"
	r "gopkg.in/dancannon/gorethink.v2"
)

type UserRepo struct {
	Config *config.Config
}

type User struct {
	*app.User
	ID string `gorethink:"id" json:"id"`
	LoginExpires int `json:"-"`
}

func NewUserFromRequest(u *app.User) *User {
	return &User{User: u}
}

func (ur *UserRepo) FindOneByFieldAndValue(field, val string) (*User, error) {
	sess, err := DbSession(ur.Config)
	if err != nil {
		return nil, errors.NewServiceError(err.Error(), 500)
	}
	res := &User{}
	q := map[string]interface{}{field: val}
	c, err := r.DB(DB_NAME).Table(USER_TABLE).Filter(q).Run(sess)
	if err != nil {
		return nil, errors.NewServiceError("failed to find a user "+err.Error(), 500)
	}
	if c.IsNil() {
		return nil, nil
	}
	c.One(res)
	return res, nil
}

func (ur *UserRepo) Save(u *User) error {
	sess, err := DbSession(ur.Config)
	if err != nil {
		return errors.NewServiceError(err.Error(), 500)
	}
	_, err = r.DB(DB_NAME).Table(USER_TABLE).Insert(u).RunWrite(sess)
	if err != nil {
		return errors.NewServiceError("failed to insert councillor "+err.Error(), 500)
	}
	return nil
}

func (ur *UserRepo) DeleteByFieldAndValue(field string, value interface{}) error {
	sess, err := DbSession(ur.Config)
	if err != nil {
		return errors.NewServiceError(err.Error(), 500)
	}
	q := map[string]interface{}{field: value}
	if _, err := r.DB(DB_NAME).Table(USER_TABLE).Filter(q).Delete().Run(sess); err != nil {
		return errors.NewServiceError("failed to delete users "+err.Error(), 500)
	}
	return nil
}
