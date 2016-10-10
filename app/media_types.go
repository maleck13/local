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

import "github.com/goadesign/goa"

// A User of locals (default view)
//
// Identifier: application/vnd.goa.local.user+json; view=default
type GoaLocalUser struct {
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
