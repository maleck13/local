//************************************************************************//
// API "locals": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/maleck13/local/design
// --out=$(GOPATH)/src/github.com/maleck13/local
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// An communication (default view)
//
// Identifier: application/vnd.goa.local.communication+json; view=default
type GoaLocalCommunication struct {
	Body   string  `form:"body" json:"body" xml:"body"`
	CommID *string `form:"commID,omitempty" json:"commID,omitempty" xml:"commID,omitempty"`
	From   *string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	// db id
	ID          string     `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	IsPrivate   bool       `form:"isPrivate" json:"isPrivate" xml:"isPrivate"`
	Open        bool       `form:"open" json:"open" xml:"open"`
	RecepientID string     `form:"recepientID" json:"recepientID" xml:"recepientID"`
	Sent        *time.Time `form:"sent,omitempty" json:"sent,omitempty" xml:"sent,omitempty"`
	Subject     string     `form:"subject" json:"subject" xml:"subject"`
	To          *string    `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`
	Type        *string    `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	UserID      *string    `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
}

// Validate validates the GoaLocalCommunication media type instance.
func (mt *GoaLocalCommunication) Validate() (err error) {
	if mt.RecepientID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "recepientID"))
	}
	if mt.Subject == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subject"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}

	return
}

// DecodeGoaLocalCommunication decodes the GoaLocalCommunication instance encoded in resp body.
func (c *Client) DecodeGoaLocalCommunication(resp *http.Response) (*GoaLocalCommunication, error) {
	var decoded GoaLocalCommunication
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaLocalCommunicationCollection is the media type for an array of GoaLocalCommunication (default view)
//
// Identifier: application/vnd.goa.local.communication+json; type=collection; view=default
type GoaLocalCommunicationCollection []*GoaLocalCommunication

// Validate validates the GoaLocalCommunicationCollection media type instance.
func (mt GoaLocalCommunicationCollection) Validate() (err error) {
	for _, e := range mt {
		if e.RecepientID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "recepientID"))
		}
		if e.Subject == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "subject"))
		}
		if e.Body == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "body"))
		}

	}
	return
}

// DecodeGoaLocalCommunicationCollection decodes the GoaLocalCommunicationCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalCommunicationCollection(resp *http.Response) (GoaLocalCommunicationCollection, error) {
	var decoded GoaLocalCommunicationCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// GoaLocalConsituents media type (default view)
//
// Identifier: application/vnd.goa.local.consituents+json; view=default
type GoaLocalConsituents struct {
	ID         *string                  `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	FirstName  *string                  `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	OpenComms  []*GoaLocalCommunication `form:"openComms,omitempty" json:"openComms,omitempty" xml:"openComms,omitempty"`
	SecondName *string                  `form:"secondName,omitempty" json:"secondName,omitempty" xml:"secondName,omitempty"`
}

// Validate validates the GoaLocalConsituents media type instance.
func (mt *GoaLocalConsituents) Validate() (err error) {
	for _, e := range mt.OpenComms {
		if e.RecepientID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.openComms[*]`, "recepientID"))
		}
		if e.Subject == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.openComms[*]`, "subject"))
		}
		if e.Body == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response.openComms[*]`, "body"))
		}

	}
	return
}

// GoaLocalConsituents media type (nocomms view)
//
// Identifier: application/vnd.goa.local.consituents+json; view=nocomms
type GoaLocalConsituentsNocomms struct {
	ID           *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	FirstName    *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	HasOpenComms *bool   `form:"hasOpenComms,omitempty" json:"hasOpenComms,omitempty" xml:"hasOpenComms,omitempty"`
	SecondName   *string `form:"secondName,omitempty" json:"secondName,omitempty" xml:"secondName,omitempty"`
}

// DecodeGoaLocalConsituents decodes the GoaLocalConsituents instance encoded in resp body.
func (c *Client) DecodeGoaLocalConsituents(resp *http.Response) (*GoaLocalConsituents, error) {
	var decoded GoaLocalConsituents
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeGoaLocalConsituentsNocomms decodes the GoaLocalConsituentsNocomms instance encoded in resp body.
func (c *Client) DecodeGoaLocalConsituentsNocomms(resp *http.Response) (*GoaLocalConsituentsNocomms, error) {
	var decoded GoaLocalConsituentsNocomms
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaLocalConsituentsCollection is the media type for an array of GoaLocalConsituents (default view)
//
// Identifier: application/vnd.goa.local.consituents+json; type=collection; view=default
type GoaLocalConsituentsCollection []*GoaLocalConsituents

// Validate validates the GoaLocalConsituentsCollection media type instance.
func (mt GoaLocalConsituentsCollection) Validate() (err error) {
	for _, e := range mt {
		for _, e := range e.OpenComms {
			if e.RecepientID == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].openComms[*]`, "recepientID"))
			}
			if e.Subject == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].openComms[*]`, "subject"))
			}
			if e.Body == "" {
				err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*].openComms[*]`, "body"))
			}

		}
	}
	return
}

// GoaLocalConsituentsCollection is the media type for an array of GoaLocalConsituents (nocomms view)
//
// Identifier: application/vnd.goa.local.consituents+json; type=collection; view=nocomms
type GoaLocalConsituentsNocommsCollection []*GoaLocalConsituentsNocomms

// DecodeGoaLocalConsituentsCollection decodes the GoaLocalConsituentsCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalConsituentsCollection(resp *http.Response) (GoaLocalConsituentsCollection, error) {
	var decoded GoaLocalConsituentsCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeGoaLocalConsituentsNocommsCollection decodes the GoaLocalConsituentsNocommsCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalConsituentsNocommsCollection(resp *http.Response) (GoaLocalConsituentsNocommsCollection, error) {
	var decoded GoaLocalConsituentsNocommsCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// GoaLocalCouncillor media type (default view)
//
// Identifier: application/vnd.goa.local.councillor+json; view=default
type GoaLocalCouncillor struct {
	// Unique user ID
	ID string `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	// a phone contact for the user
	Address string `form:"address" json:"address" xml:"address"`
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
	// The area of the users local council
	County string `form:"county" json:"county" xml:"county"`
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// facebook handle for the user
	Facebook *string `form:"facebook,omitempty" json:"facebook,omitempty" xml:"facebook,omitempty"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// an image url for the user
	Image string `form:"image" json:"image" xml:"image"`
	// the councillors party
	Party string `form:"party" json:"party" xml:"party"`
	// a phone contact for the user
	Phone string `form:"phone" json:"phone" xml:"phone"`
	// Name of the user
	SecondName string `form:"secondName" json:"secondName" xml:"secondName"`
	// twitter handle for the user
	Twitter *string `form:"twitter,omitempty" json:"twitter,omitempty" xml:"twitter,omitempty"`
	// reference to the user associated with this councillor
	UserID string `form:"userID" json:"userID" xml:"userID"`
	// a web link for the user
	Web string `form:"web" json:"web" xml:"web"`
}

// Validate validates the GoaLocalCouncillor media type instance.
func (mt *GoaLocalCouncillor) Validate() (err error) {
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if mt.SecondName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if mt.Area == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "area"))
	}
	if mt.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	if mt.Phone == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "phone"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Party == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "party"))
	}
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}
	if mt.County == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "county"))
	}
	if mt.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userID"))
	}

	return
}

// DecodeGoaLocalCouncillor decodes the GoaLocalCouncillor instance encoded in resp body.
func (c *Client) DecodeGoaLocalCouncillor(resp *http.Response) (*GoaLocalCouncillor, error) {
	var decoded GoaLocalCouncillor
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaLocalCouncillorCollection is the media type for an array of GoaLocalCouncillor (default view)
//
// Identifier: application/vnd.goa.local.councillor+json; type=collection; view=default
type GoaLocalCouncillorCollection []*GoaLocalCouncillor

// Validate validates the GoaLocalCouncillorCollection media type instance.
func (mt GoaLocalCouncillorCollection) Validate() (err error) {
	for _, e := range mt {
		if e.FirstName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "firstName"))
		}
		if e.SecondName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "secondName"))
		}
		if e.Area == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "area"))
		}
		if e.Image == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "image"))
		}
		if e.Phone == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "phone"))
		}
		if e.Email == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "email"))
		}
		if e.Party == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "party"))
		}
		if e.Address == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "address"))
		}
		if e.County == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "county"))
		}
		if e.UserID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "userID"))
		}

	}
	return
}

// DecodeGoaLocalCouncillorCollection decodes the GoaLocalCouncillorCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalCouncillorCollection(resp *http.Response) (GoaLocalCouncillorCollection, error) {
	var decoded GoaLocalCouncillorCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A User of locals (default view)
//
// Identifier: application/vnd.goa.local.user+json; view=default
type GoaLocalUser struct {
	// whether the user is activated or not
	Active bool `form:"active" json:"active" xml:"active"`
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The county the user lives in
	County string `form:"county" json:"county" xml:"county"`
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// API href for making requests on the bottle
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// Unique bottle ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the user
	SecondName string `form:"secondName" json:"secondName" xml:"secondName"`
	// the signupType of user google local
	SignupType *string `form:"signupType,omitempty" json:"signupType,omitempty" xml:"signupType,omitempty"`
	// the type of user admin local councillor
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the GoaLocalUser media type instance.
func (mt *GoaLocalUser) Validate() (err error) {
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if mt.SecondName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}

	return
}

// A User of locals (login view)
//
// Identifier: application/vnd.goa.local.user+json; view=login
type GoaLocalUserLogin struct {
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The county the user lives in
	County string `form:"county" json:"county" xml:"county"`
	// Unique bottle ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The area of the users local council
	LoginExpires *int `form:"loginExpires,omitempty" json:"loginExpires,omitempty" xml:"loginExpires,omitempty"`
	// user action status
	Status bool `form:"status" json:"status" xml:"status"`
	// This can be an oauth token or a password
	Token string `form:"token" json:"token" xml:"token"`
	// the type of user admin local councillor
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the GoaLocalUserLogin media type instance.
func (mt *GoaLocalUserLogin) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	return
}

// A User of locals (public view)
//
// Identifier: application/vnd.goa.local.user+json; view=public
type GoaLocalUserPublic struct {
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Unique bottle ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// Validate validates the GoaLocalUserPublic media type instance.
func (mt *GoaLocalUserPublic) Validate() (err error) {
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}

	return
}

// DecodeGoaLocalUser decodes the GoaLocalUser instance encoded in resp body.
func (c *Client) DecodeGoaLocalUser(resp *http.Response) (*GoaLocalUser, error) {
	var decoded GoaLocalUser
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeGoaLocalUserLogin decodes the GoaLocalUserLogin instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserLogin(resp *http.Response) (*GoaLocalUserLogin, error) {
	var decoded GoaLocalUserLogin
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeGoaLocalUserPublic decodes the GoaLocalUserPublic instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserPublic(resp *http.Response) (*GoaLocalUserPublic, error) {
	var decoded GoaLocalUserPublic
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// GoaLocalUserCollection is the media type for an array of GoaLocalUser (default view)
//
// Identifier: application/vnd.goa.local.user+json; type=collection; view=default
type GoaLocalUserCollection []*GoaLocalUser

// Validate validates the GoaLocalUserCollection media type instance.
func (mt GoaLocalUserCollection) Validate() (err error) {
	for _, e := range mt {
		if e.FirstName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "firstName"))
		}
		if e.SecondName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "secondName"))
		}
		if e.Email == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "email"))
		}

	}
	return
}

// GoaLocalUserCollection is the media type for an array of GoaLocalUser (login view)
//
// Identifier: application/vnd.goa.local.user+json; type=collection; view=login
type GoaLocalUserLoginCollection []*GoaLocalUserLogin

// Validate validates the GoaLocalUserLoginCollection media type instance.
func (mt GoaLocalUserLoginCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Token == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "token"))
		}

	}
	return
}

// GoaLocalUserCollection is the media type for an array of GoaLocalUser (public view)
//
// Identifier: application/vnd.goa.local.user+json; type=collection; view=public
type GoaLocalUserPublicCollection []*GoaLocalUserPublic

// Validate validates the GoaLocalUserPublicCollection media type instance.
func (mt GoaLocalUserPublicCollection) Validate() (err error) {
	for _, e := range mt {
		if e.FirstName == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "firstName"))
		}

	}
	return
}

// DecodeGoaLocalUserCollection decodes the GoaLocalUserCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserCollection(resp *http.Response) (GoaLocalUserCollection, error) {
	var decoded GoaLocalUserCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeGoaLocalUserLoginCollection decodes the GoaLocalUserLoginCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserLoginCollection(resp *http.Response) (GoaLocalUserLoginCollection, error) {
	var decoded GoaLocalUserLoginCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeGoaLocalUserPublicCollection decodes the GoaLocalUserPublicCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserPublicCollection(resp *http.Response) (GoaLocalUserPublicCollection, error) {
	var decoded GoaLocalUserPublicCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
