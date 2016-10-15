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
)

// A Councillor (default view)
//
// Identifier: application/vnd.goa.local.councillor+json; view=default
type GoaLocalCouncillor struct {
	// a phone contact for the user
	Address string `form:"address" json:"address" xml:"address"`
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
	// The area of the users local council
	County string `form:"county" json:"county" xml:"county"`
	// email for the councillor
	Email string `form:"email" json:"email" xml:"email"`
	// facebook handle for the user
	Facebook *string `form:"facebook,omitempty" json:"facebook,omitempty" xml:"facebook,omitempty"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// db id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
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
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The county the user lives in
	County *string `form:"county,omitempty" json:"county,omitempty" xml:"county,omitempty"`
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

// A User of locals (full view)
//
// Identifier: application/vnd.goa.local.user+json; view=full
type GoaLocalUserFull struct {
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// API href for making requests on the bottle
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// Unique bottle ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The area of the users local council
	Location *Location `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// The area of the users local council
	LoginExpires *int `form:"loginExpires,omitempty" json:"loginExpires,omitempty" xml:"loginExpires,omitempty"`
	// Name of the user
	SecondName string `form:"secondName" json:"secondName" xml:"secondName"`
	// the signupType of user google local
	SignupType *string `form:"signupType,omitempty" json:"signupType,omitempty" xml:"signupType,omitempty"`
	// This can be an oauth token or a password
	Token string `form:"token" json:"token" xml:"token"`
	// the type of user admin local councillor
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the GoaLocalUserFull media type instance.
func (mt *GoaLocalUserFull) Validate() (err error) {
	if mt.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if mt.SecondName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	return
}

// A User of locals (login view)
//
// Identifier: application/vnd.goa.local.user+json; view=login
type GoaLocalUserLogin struct {
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

// DecodeGoaLocalUserFull decodes the GoaLocalUserFull instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserFull(resp *http.Response) (*GoaLocalUserFull, error) {
	var decoded GoaLocalUserFull
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

// GoaLocalUserCollection is the media type for an array of GoaLocalUser (full view)
//
// Identifier: application/vnd.goa.local.user+json; type=collection; view=full
type GoaLocalUserFullCollection []*GoaLocalUserFull

// Validate validates the GoaLocalUserFullCollection media type instance.
func (mt GoaLocalUserFullCollection) Validate() (err error) {
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
		if e.Token == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "token"))
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

// DecodeGoaLocalUserFullCollection decodes the GoaLocalUserFullCollection instance encoded in resp body.
func (c *Client) DecodeGoaLocalUserFullCollection(resp *http.Response) (GoaLocalUserFullCollection, error) {
	var decoded GoaLocalUserFullCollection
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
