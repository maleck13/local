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

// NewCouncillorRepo returns a configured CouncillorRepo
func NewCouncillorRepo(conf *config.Config, actor Actor, authorisor Authorisor) CouncillorRepo {
	return CouncillorRepo{
		Config:     conf,
		Actor:      actor,
		Authorisor: authorisor,
	}
}
