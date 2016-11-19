package domain

import (
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	r "gopkg.in/dancannon/gorethink.v2"
)

// Communication represent a piece of Communication between two parties
type Communication struct {
	*app.Communication
	UserID string `json:"userID"`
}

// AccessTypes defines the non owner access
func (c *Communication) AccessTypes() map[string][]string {
	access := map[string][]string{
		"read":  []string{"admin", "councillor"},
		"write": []string{"admin"},
	}
	return access
}

//Owner implements AccessDefinor
func (c *Communication) Owner() string {
	return c.UserID
}

func NewCommunication(c *app.Communication) *Communication {
	return &Communication{
		Communication: c,
	}
}

// CommunicationSaver is the interface for how Communications should be updated and saved
type CommunicationSaver interface {
	SaveUpdate(*Communication) error
}

type CommunicationFinder interface {
	FindAllByKeyValue(key string, value interface{}) ([]*Communication, error)
	FindAllByKeyValues(kv map[string]interface{}) ([]*Communication, error)
}

type CommunicationRepo struct {
	Config     *config.Config
	Authorisor Authorisor
	Actor      Actor
}

func NewCommunicationRepo(conf *config.Config, actor Actor, authorisor Authorisor) CommunicationRepo {
	return CommunicationRepo{
		Config:     conf,
		Authorisor: authorisor,
		Actor:      actor,
	}
}

func (cr CommunicationRepo) SaveUpdate(c *Communication) error {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return errors.Wrap(err, "unexpected error getting database session")
	}
	if err := cr.Authorisor.Authorise(c, "write", cr.Actor); err != nil {
		return err
	}
	var q = r.DB(data.DB_NAME).Table(data.COMMUNICATIONS_TABLE)
	if c.ID != "" {
		q = q.Update(c)
	} else {
		if c.CommID == nil {
			cid := uuid.NewV4().String()
			c.CommID = &cid
		}
		q = q.Insert(c)
	}
	if _, err := q.RunWrite(sess); err != nil {
		return errors.Wrap(err, "failed to SaveUpdate Communication")
	}
	return nil

}

func (cr CommunicationRepo) FindAllByKeyValues(kv map[string]interface{}) ([]*Communication, error) {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error getting database session")
	}
	c, err := r.DB(data.DB_NAME).Table(data.COMMUNICATIONS_TABLE).Filter(kv).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error listing Communications")
	}
	var ret = []*Communication{}
	if c.IsNil() {
		return ret, nil
	}
	if err := c.All(&ret); err != nil {
		return nil, errors.Wrap(err, "unexpected error retreiving result set")
	}
	return ret, nil
}

// FindAllByCouncillerIDAndUserID returns a list of Communications between a councillor and a user
func (cr CommunicationRepo) FindAllByRecepientIDAndUserID(cid, uid string, open bool) ([]*Communication, error) {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error getting database session")
	}
	f := map[string]interface{}{"UserID": uid, "RecepientID": cid, "Open": open}
	c, err := r.DB(data.DB_NAME).Table(data.COMMUNICATIONS_TABLE).Filter(f).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error listing Communications")
	}
	var ret = []*Communication{}
	if c.IsNil() {
		return ret, nil
	}
	if err := c.All(&ret); err != nil {
		return nil, errors.Wrap(err, "unexpected error retreiving result set")
	}
	return ret, nil
}

// FindAllByKeyValue finds based on passed key and value. Case sensitive. Applies a filter to the result based on the actors access
func (cr CommunicationRepo) FindAllByKeyValue(key string, value interface{}) ([]*Communication, error) {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error getting database session")
	}
	f := map[string]interface{}{key: value}
	c, err := r.DB(data.DB_NAME).Table(data.COMMUNICATIONS_TABLE).Filter(f).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error listing Communications")
	}
	var ret = []*Communication{}
	if c.IsNil() {
		return ret, nil
	}
	if err := c.All(&ret); err != nil {
		return nil, errors.Wrap(err, "unexpected error retreiving result set")
	}
	return cr.filterEntities(ret, "read"), nil
}

func (cr CommunicationRepo) FindFirstOpenBySenderAndReceipent(sender, receipient string) (*Communication, error) {
	sess, err := data.DbSession(cr.Config)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error getting database session")
	}
	f := map[string]interface{}{"RecepientID": receipient, "UserID": sender, "Open": true}
	c, err := r.DB(data.DB_NAME).Table(data.COMMUNICATIONS_TABLE).Filter(f).Limit(1).Run(sess)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error finding Communications")
	}
	if c.IsNil() {
		return nil, nil
	}
	comm := &Communication{}
	if err := c.One(comm); err != nil {
		return nil, errors.Wrap(err, "failed to find one communication")
	}
	return comm, nil
}

func (cr CommunicationRepo) filterEntities(entities []*Communication, action string) []*Communication {
	filtered := []*Communication{}
	for _, e := range entities {
		if err := cr.Authorisor.Authorise(e, action, cr.Actor); err == nil {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
