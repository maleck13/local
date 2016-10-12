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

// councillor user type.
type councillor struct {
	// a phone contact for the user
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
	// email for the councillor
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// facebook handle for the user
	Facebook *string `form:"facebook,omitempty" json:"facebook,omitempty" xml:"facebook,omitempty"`
	// Name of the user
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// an image url for the user
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// the councillors party
	Party *string `form:"party,omitempty" json:"party,omitempty" xml:"party,omitempty"`
	// a phone contact for the user
	Phone *string `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
	// Name of the user
	SecondName *string `form:"secondName,omitempty" json:"secondName,omitempty" xml:"secondName,omitempty"`
	// twitter handle for the user
	Twitter *string `form:"twitter,omitempty" json:"twitter,omitempty" xml:"twitter,omitempty"`
	// a web link for the user
	Web *string `form:"web,omitempty" json:"web,omitempty" xml:"web,omitempty"`
}

// Finalize sets the default values for councillor type instance.
func (ut *councillor) Finalize() {
	var defaultAddress = ""
	if ut.Address == nil {
		ut.Address = &defaultAddress
	}
	var defaultArea = ""
	if ut.Area == nil {
		ut.Area = &defaultArea
	}
	var defaultEmail = ""
	if ut.Email == nil {
		ut.Email = &defaultEmail
	}
	var defaultImage = ""
	if ut.Image == nil {
		ut.Image = &defaultImage
	}
	var defaultParty = ""
	if ut.Party == nil {
		ut.Party = &defaultParty
	}
	var defaultPhone = ""
	if ut.Phone == nil {
		ut.Phone = &defaultPhone
	}
	var defaultWeb = ""
	if ut.Web == nil {
		ut.Web = &defaultWeb
	}
}

// Validate validates the councillor type instance.
func (ut *councillor) Validate() (err error) {
	if ut.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if ut.SecondName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if ut.Area == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "area"))
	}
	if ut.Image == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	if ut.Phone == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "phone"))
	}
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Party == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "party"))
	}
	if ut.Address == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	return
}

// Publicize creates Councillor from councillor
func (ut *councillor) Publicize() *Councillor {
	var pub Councillor
	if ut.Address != nil {
		pub.Address = *ut.Address
	}
	if ut.Area != nil {
		pub.Area = *ut.Area
	}
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.Facebook != nil {
		pub.Facebook = ut.Facebook
	}
	if ut.FirstName != nil {
		pub.FirstName = *ut.FirstName
	}
	if ut.Image != nil {
		pub.Image = *ut.Image
	}
	if ut.Party != nil {
		pub.Party = *ut.Party
	}
	if ut.Phone != nil {
		pub.Phone = *ut.Phone
	}
	if ut.SecondName != nil {
		pub.SecondName = *ut.SecondName
	}
	if ut.Twitter != nil {
		pub.Twitter = ut.Twitter
	}
	if ut.Web != nil {
		pub.Web = *ut.Web
	}
	return &pub
}

// Councillor user type.
type Councillor struct {
	// a phone contact for the user
	Address string `form:"address" json:"address" xml:"address"`
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
	// email for the councillor
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
	// a web link for the user
	Web string `form:"web" json:"web" xml:"web"`
}

// Validate validates the Councillor type instance.
func (ut *Councillor) Validate() (err error) {
	if ut.FirstName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "firstName"))
	}
	if ut.SecondName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secondName"))
	}
	if ut.Area == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "area"))
	}
	if ut.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	if ut.Phone == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "phone"))
	}
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if ut.Party == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "party"))
	}
	if ut.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
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
	// The area of the users local council
	Area *string `form:"area,omitempty" json:"area,omitempty" xml:"area,omitempty"`
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
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Area != nil {
		pub.Area = *ut.Area
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
	// The area of the users local council
	Area string `form:"area" json:"area" xml:"area"`
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

	return
}
