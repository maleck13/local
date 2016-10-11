package domain

import (
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/pkg/errors"
	r "gopkg.in/dancannon/gorethink.v2"
)

// User represents a user of locals in the database
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

// NewUserFromRequest conversts an app.User to domain.User
func NewUserFromRequest(u *app.User) *User {
	return &User{User: u}
}

// UserSaver defines something that saves users
type UserSaver interface {
	SaveUpdate(u *User) error
}

// UserDeleter defines how a user should be deleted
type UserDeleter interface {
	DeleteByFieldAndValue(field string, value interface{}) error
}

// UserFinder defines how users should be found
type UserFinder interface {
	FindOneByFieldAndValue(field, val string) (*User, error)
	FindAllByTypeAndArea(uType, area string) ([]*User, error)
}

// UserFinderDeleterSaver composite interface
type UserFinderDeleterSaver interface {
	UserSaver
	UserDeleter
	UserFinder
}

// UserRepo access users in the data layer
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

// FindOneByFieldAndValue returns a user based on a field and value if there is no user found both error and user will be nil
func (ur UserRepo) FindOneByFieldAndValue(field, val string) (*User, error) {
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error FindOneByFieldAndValue")
	}
	res := &User{}
	q := map[string]interface{}{field: val}
	c, err := r.DB(data.DB_NAME).Table(data.USER_TABLE).Filter(q).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error FindOneByFieldAndValue")
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

// FindAllByTypeAndArea finds all users by type and by area
func (ur UserRepo) FindAllByTypeAndArea(uType, area string) ([]*User, error) {
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error FindAllByTypeAndArea")
	}
	res := []*User{}
	q := map[string]string{"Type": uType, "Area": area}
	c, err := r.DB(data.DB_NAME).Table(data.USER_TABLE).Filter(q).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error FindAllByTypeAndArea")
	}
	if c.IsNil() {
		return nil, nil
	}
	if err := c.All(&res); err != nil {
		return nil, errors.Wrap(err, "unexpected error FindAllByTypeAndArea")
	}
	for _, u := range res {
		if err := ur.Authorisor.Authorise(u, "read", ur.Actor); err != nil {
			return nil, err
		}
	}
	return res, nil
}

// SaveUpdate will save or update a User
func (ur UserRepo) SaveUpdate(u *User) error {
	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return errors.Wrap(err, "unexpected error SaveUpdate")
	}
	if err := ur.Authorisor.Authorise(u, "write", ur.Actor); err != nil {
		return err
	}
	var q = r.DB(data.DB_NAME).Table(data.USER_TABLE)
	if u.ID == "" {
		q = q.Insert(u)
	} else {
		q = q.Get(u.ID).Update(u)
	}
	_, err = q.RunWrite(sess)
	if err != nil {
		return errors.Wrap(err, "unexpected error SaveUpdate")
	}
	return nil
}

// DeleteByFieldAndValue remove a user based of the value of the field
func (ur UserRepo) DeleteByFieldAndValue(field string, value interface{}) error {

	sess, err := data.DbSession(ur.Config)
	if err != nil {
		return errors.Wrap(err, "unexpected error DeleteByFieldAndValue")
	}
	q := map[string]interface{}{field: value}
	if _, err := r.DB(data.DB_NAME).Table(data.USER_TABLE).Filter(q).Delete().Run(sess); err != nil {
		return errors.Wrap(err, "unexpected error DeleteByFieldAndValue")
	}
	return nil
}
