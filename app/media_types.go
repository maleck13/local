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

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// An communication (default view)
//
// Identifier: application/vnd.goa.local.communication+json; view=default
type GoaLocalCommunication struct {
	Body         string  `form:"body" json:"body" xml:"body"`
	CouncillorID string  `form:"councillorID" json:"councillorID" xml:"councillorID"`
	From         *string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	// db id
	ID        string     `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	IsPrivate bool       `form:"isPrivate" json:"isPrivate" xml:"isPrivate"`
	Open      bool       `form:"open" json:"open" xml:"open"`
	Sent      *time.Time `form:"sent,omitempty" json:"sent,omitempty" xml:"sent,omitempty"`
	Subject   string     `form:"subject" json:"subject" xml:"subject"`
	To        *string    `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`
	UserID    *string    `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
}

// Validate validates the GoaLocalCommunication media type instance.
func (mt *GoaLocalCommunication) Validate() (err error) {
	if mt.CouncillorID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "councillorID"))
	}
	if mt.Subject == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subject"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}

	return
}

// GoaLocalCommunicationCollection is the media type for an array of GoaLocalCommunication (default view)
//
// Identifier: application/vnd.goa.local.communication+json; type=collection; view=default
type GoaLocalCommunicationCollection []*GoaLocalCommunication

// Validate validates the GoaLocalCommunicationCollection media type instance.
func (mt GoaLocalCommunicationCollection) Validate() (err error) {
	for _, e := range mt {
		if e.CouncillorID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "councillorID"))
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
	ID string `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
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

// A User of locals (default view)
//
// Identifier: application/vnd.goa.local.user+json; view=default
type GoaLocalUser struct {
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
