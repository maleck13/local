package domain

import (
	"crypto/md5"
	"fmt"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/pkg/errors"
	r "gopkg.in/dancannon/gorethink.v2"
)

// Councillor is data representation of a local council member
type Councillor struct {
	*app.Councillor
	ID string `gorethink:"id,omitempty"`
}

// AccessTypes defines the non owner access
func (c *Councillor) AccessTypes() map[string][]string {
	access := map[string][]string{
		"read":  []string{"admin"},
		"write": []string{"admin"},
	}
	return access
}

//Owner implements AccessDefinor
func (c *Councillor) Owner() string {
	return c.ID
}

// GenerateID a repeatable id
func (c *Councillor) GenerateID() string {
	data := []byte(c.Email)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// NewCouncillor returns a new councillor model
func NewCouncillor(c *app.Councillor) *Councillor {
	return &Councillor{
		Councillor: c,
	}
}

// CouncillorSaver is the interface for persisting councillors to the db
type CouncillorSaver interface {
	SaveUpdate(councillor *Councillor) error
}

// CouncillorRepo is a repository for handling data interactions for councillors
type CouncillorRepo struct {
	Config     *config.Config
	Actor      Actor
	Authorisor Authorisor
}

// SaveUpdate will save or update a councillor
func (cr CouncillorRepo) SaveUpdate(councillor *Councillor) error {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return errors.Wrap(err, "unexpected error getting database session")
	}
	if err := cr.Authorisor.Authorise(councillor, "write", cr.Actor); err != nil {
		return err
	}
	var q = r.DB(data.DB_NAME).Table(data.COUNCILLORS_TABLE)
	if councillor.ID == "" {
		councillor.ID = councillor.GenerateID()
		q = q.Insert(councillor)
	} else {
		q = q.Get(councillor.ID).Update(councillor)
	}
	_, err = q.RunWrite(sess)
	if err != nil {
		return errors.Wrap(err, "unexpected error writing data to db")
	}
	return nil
}

// DeleteAllByKeyValue will remove all entities that match. Can only be called by an admin user
func (cr CouncillorRepo) DeleteAllByKeyValue(key string, value interface{}) error {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return errors.Wrap(err, "unexpected error getting database session")
	}
	if cr.Actor.Type() != "admin" {
		return errors.New("only admins can delete all")
	}
	filter := map[string]interface{}{key: value}
	_, err = r.DB(data.DB_NAME).Table(data.COUNCILLORS_TABLE).Filter(filter).Delete().RunWrite(sess)
	if err != nil {
		return errors.Wrap(err, "failed to delete councillors")
	}
	return nil
}

// FindOneByKeyValue will find the first row that matches key and value
func (cr CouncillorRepo) FindOneByKeyValue(key string, value interface{}) (*Councillor, error) {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error getting database session")
	}
	q := map[string]interface{}{key: value}
	c, err := r.DB(data.DB_NAME).Table(data.COUNCILLORS_TABLE).Filter(q).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read councillor")
	}
	if c.IsNil() {
		return nil, nil
	}
	councillor := &Councillor{}
	if err := c.One(councillor); err != nil {
		return nil, errors.Wrap(err, "failed to marshal result into a councillor")
	}
	return councillor, nil
}

// NewCouncillorRepo returns a configured CouncillorRepo
func NewCouncillorRepo(conf *config.Config, actor Actor, authorisor Authorisor) CouncillorRepo {
	return CouncillorRepo{
		Config:     conf,
		Actor:      actor,
		Authorisor: authorisor,
	}
}
