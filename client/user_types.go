//************************************************************************//
// API "locals": Application User Types
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

import "github.com/goadesign/goa"

// location user type.
type location struct {
	Lat *float64 `form:"Lat,omitempty" json:"Lat,omitempty" xml:"Lat,omitempty"`
	Lon *float64 `form:"Lon,omitempty" json:"Lon,omitempty" xml:"Lon,omitempty"`
}

// Publicize creates Location from location
func (ut *location) Publicize() *Location {
	var pub Location
	if ut.Lat != nil {
		pub.Lat = ut.Lat
	}
	if ut.Lon != nil {
		pub.Lon = ut.Lon
	}
	return &pub
}

// Location user type.
type Location struct {
	Lat *float64 `form:"Lat,omitempty" json:"Lat,omitempty" xml:"Lat,omitempty"`
	Lon *float64 `form:"Lon,omitempty" json:"Lon,omitempty" xml:"Lon,omitempty"`
}

// login user type.
type login struct {
	Email      *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	SignupType *string `form:"signupType,omitempty" json:"signupType,omitempty" xml:"signupType,omitempty"`
	Token      *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// Validate validates the login type instance.
func (ut *login) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Token == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	if ut.SignupType == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "signupType"))
	}

	return
}

// Publicize creates Login from login
func (ut *login) Publicize() *Login {
	var pub Login
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.SignupType != nil {
		pub.SignupType = *ut.SignupType
	}
	if ut.Token != nil {
		pub.Token = *ut.Token
	}
	return &pub
}

// Login user type.
type Login struct {
	Email      string `form:"email" json:"email" xml:"email"`
	SignupType string `form:"signupType" json:"signupType" xml:"signupType"`
	Token      string `form:"token" json:"token" xml:"token"`
}

// Validate validates the Login type instance.
func (ut *Login) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	if ut.SignupType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "signupType"))
	}

	return
}

// user user type.
type user struct {
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the user
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Unique user ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// an image url for the user
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// The area of the users local council
	Location *location `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Name of the user
	SecondName *string `form:"secondName,omitempty" json:"secondName,omitempty" xml:"secondName,omitempty"`
	// the signupType of user google local
	SignupType *string `form:"signupType,omitempty" json:"signupType,omitempty" xml:"signupType,omitempty"`
	// This can be an oauth token or a password
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	// the type of user admin local councillor
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Finalize sets the default values for user type instance.
func (ut *user) Finalize() {
	var defaultArea = ""
	if ut.Area == nil {
		ut.Area = &defaultArea
	}
	var defaultImage = ""
	if ut.Image == nil {
		ut.Image = &defaultImage
	}
	var defaultSignupType = ""
	if ut.SignupType == nil {
		ut.SignupType = &defaultSignupType
	}
	var defaultType = "local"
	if ut.Type == nil {
		ut.Type = &defaultType
	}
}

// Validate validates the user type instance.
func (ut *user) Validate() (err error) {
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if ut.SecondName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Token == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	return
}

// Publicize creates User from user
func (ut *user) Publicize() *User {
	var pub User
	if ut.Area != nil {
		pub.Area = *ut.Area
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.ID != nil {
		pub.ID = ut.ID
	}
	if ut.Image != nil {
		pub.Image = *ut.Image
	}
	if ut.Location != nil {
		pub.Location = ut.Location.Publicize()
	}
	if ut.SecondName != nil {
		pub.SecondName = *ut.SecondName
	}
	if ut.SignupType != nil {
		pub.SignupType = *ut.SignupType
	}
	if ut.Token != nil {
		pub.Token = *ut.Token
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	return &pub
}

// User user type.
type User struct {
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Unique user ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// an image url for the user
	Image string `form:"image" json:"image" xml:"image"`
	// The area of the users local council
	Location *Location `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Name of the user
	SecondName string `form:"secondName" json:"secondName" xml:"secondName"`
	// the signupType of user google local
	SignupType string `form:"signupType" json:"signupType" xml:"signupType"`
	// This can be an oauth token or a password
	Token string `form:"token" json:"token" xml:"token"`
	// the type of user admin local councillor
	Type string `form:"type" json:"type" xml:"type"`
}

// Validate validates the User type instance.
func (ut *User) Validate() (err error) {
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if ut.SecondName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}

	return
}
