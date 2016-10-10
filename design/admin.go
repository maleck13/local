package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

//User handler definition host:3001/user
var _ = Resource("admin", func() {

	BasePath("/admin")     // together. They map to REST resources for REST
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("admin:access") // Enforce presence of "api" scope in JWT claims.
	})
	Action("createCouncillor", func() { // Actions define a single API endpoint together
		Description("admin api to add a councillor") // with its path, parameters (both path
		Routing(POST("/councillor"))                 // parameters and querystring values) and payload
		Response(Created)                            // Responses define the shape and status code
		Response(Unauthorized)                       // of HTTP responses.
	})
})

var adminCouncillorPayload = Type("Councillor", func() {
	Attribute("firstName", String, "Name of the user")
	Attribute("secondName", String, "Name of the user")
	Attribute("area", String, "The area of the users local council", func() {
		Default("")
	})
	Attribute("image", String, "an image url for the user", func() {
		Default("")
	})
	Attribute("email", String, "email for the councillor", func() {
		Default("")
	})
	Attribute("phone", String, "a phone contact for the user", func() {
		Default("")
	})
	Attribute("web", String, "a web link for the user", func() {
		Default("")
	})
	Attribute("party", String, "the councillors party", func() {
		Default("")
	})
	Attribute("twitter", String, "twitter handle for the user")
	Attribute("facebook", String, "facebook handle for the user")

	Required("firstName", "secondName", "area", "image", "phone", "email", "party")
})
