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

import (
	"github.com/goadesign/goa"
	"time"
)

// communication user type.
type communication struct {
	Body   *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	CommID *string `form:"commID,omitempty" json:"commID,omitempty" xml:"commID,omitempty"`
	Error  *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	From   *string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	// db id
	ID          *string    `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	IsPrivate   *bool      `form:"isPrivate,omitempty" json:"isPrivate,omitempty" xml:"isPrivate,omitempty"`
	Open        *bool      `form:"open,omitempty" json:"open,omitempty" xml:"open,omitempty"`
	RecepientID *string    `form:"recepientID,omitempty" json:"recepientID,omitempty" xml:"recepientID,omitempty"`
	Sent        *time.Time `form:"sent,omitempty" json:"sent,omitempty" xml:"sent,omitempty"`
	Subject     *string    `form:"subject,omitempty" json:"subject,omitempty" xml:"subject,omitempty"`
	To          *string    `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`
	Type        *string    `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// Finalize sets the default values for communication type instance.
func (ut *communication) Finalize() {
	var defaultError = ""
	if ut.Error == nil {
		ut.Error = &defaultError
	}
	var defaultID = ""
	if ut.ID == nil {
		ut.ID = &defaultID
	}
	var defaultIsPrivate = true
	if ut.IsPrivate == nil {
		ut.IsPrivate = &defaultIsPrivate
	}
	var defaultOpen = true
	if ut.Open == nil {
		ut.Open = &defaultOpen
	}
}

// Validate validates the communication type instance.
func (ut *communication) Validate() (err error) {
	if ut.RecepientID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "recepientID"))
	}
	if ut.Subject == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subject"))
	}
	if ut.Body == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if ut.IsPrivate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "isPrivate"))
	}
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// Publicize creates Communication from communication
func (ut *communication) Publicize() *Communication {
	var pub Communication
	if ut.Body != nil {
		pub.Body = *ut.Body
	}
	if ut.CommID != nil {
		pub.CommID = ut.CommID
	}
	if ut.Error != nil {
		pub.Error = *ut.Error
	}
	if ut.From != nil {
		pub.From = ut.From
	}
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.IsPrivate != nil {
		pub.IsPrivate = *ut.IsPrivate
	}
	if ut.Open != nil {
		pub.Open = *ut.Open
	}
	if ut.RecepientID != nil {
		pub.RecepientID = *ut.RecepientID
	}
	if ut.Sent != nil {
		pub.Sent = ut.Sent
	}
	if ut.Subject != nil {
		pub.Subject = *ut.Subject
	}
	if ut.To != nil {
		pub.To = ut.To
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	return &pub
}

// Communication user type.
type Communication struct {
	Body   string  `form:"body" json:"body" xml:"body"`
	CommID *string `form:"commID,omitempty" json:"commID,omitempty" xml:"commID,omitempty"`
	Error  string  `form:"error" json:"error" xml:"error"`
	From   *string `form:"from,omitempty" json:"from,omitempty" xml:"from,omitempty"`
	// db id
	ID          string     `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	IsPrivate   bool       `form:"isPrivate" json:"isPrivate" xml:"isPrivate"`
	Open        bool       `form:"open" json:"open" xml:"open"`
	RecepientID string     `form:"recepientID" json:"recepientID" xml:"recepientID"`
	Sent        *time.Time `form:"sent,omitempty" json:"sent,omitempty" xml:"sent,omitempty"`
	Subject     string     `form:"subject" json:"subject" xml:"subject"`
	To          *string    `form:"to,omitempty" json:"to,omitempty" xml:"to,omitempty"`
	Type        string     `form:"type" json:"type" xml:"type"`
}

// Validate validates the Communication type instance.
func (ut *Communication) Validate() (err error) {
	if ut.RecepientID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "recepientID"))
	}
	if ut.Subject == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "subject"))
	}
	if ut.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

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

// updateUser user type.
type updateUser struct {
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The area of the users local council
	County *string `form:"county,omitempty" json:"county,omitempty" xml:"county,omitempty"`
	// The email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the user
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Unique user ID
	ID *string `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	// an image url for the user
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// Name of the user
	SecondName *string `form:"secondName,omitempty" json:"secondName,omitempty" xml:"secondName,omitempty"`
}

// Finalize sets the default values for updateUser type instance.
func (ut *updateUser) Finalize() {
	var defaultArea = ""
	if ut.Area == nil {
		ut.Area = &defaultArea
	}
	var defaultCounty = ""
	if ut.County == nil {
		ut.County = &defaultCounty
	}
	var defaultImage = ""
	if ut.Image == nil {
		ut.Image = &defaultImage
	}
}

// Validate validates the updateUser type instance.
func (ut *updateUser) Validate() (err error) {
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if ut.SecondName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ut.Area == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "area"))
	}

	return
}

// Publicize creates UpdateUser from updateUser
func (ut *updateUser) Publicize() *UpdateUser {
	var pub UpdateUser
	if ut.Area != nil {
		pub.Area = *ut.Area
	}
	if ut.County != nil {
		pub.County = *ut.County
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Image != nil {
		pub.Image = *ut.Image
	}
	if ut.SecondName != nil {
		pub.SecondName = *ut.SecondName
	}
	return &pub
}

// UpdateUser user type.
type UpdateUser struct {
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
	// The area of the users local council
	County string `form:"county" json:"county" xml:"county"`
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Unique user ID
	ID string `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	// an image url for the user
	Image string `form:"image" json:"image" xml:"image"`
	// Name of the user
	SecondName string `form:"secondName" json:"secondName" xml:"secondName"`
}

// Validate validates the UpdateUser type instance.
func (ut *UpdateUser) Validate() (err error) {
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if ut.SecondName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ut.Area == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "area"))
	}

	return
}

// user user type.
type user struct {
	// Unique user ID
	ID *string `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	// whether the user is active or not
	Active *bool `form:"active,omitempty" json:"active,omitempty" xml:"active,omitempty"`
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// The area of the users local council
	County *string `form:"county,omitempty" json:"county,omitempty" xml:"county,omitempty"`
	// The email of the user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Name of the user
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
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
	var defaultID = ""
	if ut.ID == nil {
		ut.ID = &defaultID
	}
	var defaultActive = true
	if ut.Active == nil {
		ut.Active = &defaultActive
	}
	var defaultArea = ""
	if ut.Area == nil {
		ut.Area = &defaultArea
	}
	var defaultCounty = ""
	if ut.County == nil {
		ut.County = &defaultCounty
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
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}

// Publicize creates User from user
func (ut *user) Publicize() *User {
	var pub User
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Active != nil {
		pub.Active = *ut.Active
	}
	if ut.Area != nil {
		pub.Area = *ut.Area
	}
	if ut.County != nil {
		pub.County = *ut.County
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
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
	// Unique user ID
	ID string `form:"id,omitempty" gorethink:"id,omitempty" json:"id,omitempty"`
	// whether the user is active or not
	Active bool `form:"active" json:"active" xml:"active"`
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
	// The area of the users local council
	County string `form:"county" json:"county" xml:"county"`
	// The email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
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
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	return
}
