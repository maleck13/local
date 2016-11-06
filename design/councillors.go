package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
) // User defines the media type used to render users. The mediaType gives all fields and the views break it down to specific responses

var _ = Resource("councillors", func() {
	BasePath("/councillors") // together. They map to REST resources for REST
	DefaultMedia(Councillor, "default")
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})
	Action("listForCountyAndArea", func() {
		Description("list councillors based on a users details") // with its path, parameters (both path
		Routing(GET("/"))
		Params(func() { // (shape of the request body).
			Param("area", String)
			Param("county", String)
		})
		Response(OK, func() {
			Media(CollectionOf(Councillor))
		})
		Response(Unauthorized) // of HTTP responses.
	})
	Action("readById", func() {
		Description("read a councillor based on an id")
		Routing(GET("/:id"))
		Params(func() { // (shape of the request body).
			Param("id", String)
		})
		Response(OK, func() {
			Media(Councillor)
		})
		Response(Unauthorized) // of HTTP responses.
	})
	Action("update", func() {
		Description("update a councillors details")
		Routing(POST("/:id"))
		Payload(CouncillorUpdate)
		Params(func() { // (shape of the request body).
			Param("id", String)
		})
		Response(OK, func() {
			Media(Councillor)
		})
		Response(Unauthorized) // of HTTP responses.
	})
	Action("uploadProfilePic", func() {
		Description("update a councillors details")
		Routing(POST("/:id/image"))
		Params(func() { // (shape of the request body).
			Param("id", String)
		})
		Response(OK, func() {
			Media(Councillor)
		})
		Response(Unauthorized) // of HTTP responses.
	})
})

var Councillor = MediaType("application/vnd.goa.local.councillor+json", func() {
	Reference(UserPayload)
	Attributes(func() { // Attributes define the media type shape.

		Attribute("ID")
		Attribute("firstName")
		Attribute("secondName")
		Attribute("area")
		Attribute("county")
		Attribute("email")
		Attribute("inOffice", Boolean, "whether the councillor is still in office", func() {
			Default(false)
		})
		Attribute("userID", String, "reference to the user associated with this councillor")
		Attribute("image", String, "an image url for the user", func() {
			Default("")
		})

		Attribute("phone", String, "a phone contact for the user", func() {
			Default("")
		})
		Attribute("address", String, "a phone contact for the user", func() {
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
		Required("firstName", "secondName", "area", "image", "phone", "email", "party", "address", "county", "userID")
	})
	View("default", func() {
		Attribute("firstName", String, "Name of the user")
		Attribute("userID", String, "reference to the user associated with this councillor")
		Attribute("secondName", String, "Name of the user")
		Attribute("area", String, "The area of the users local council", func() {
			Default("")
		})
		Attribute("county", String, "The area of the users local council", func() {
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
		Attribute("address", String, "a phone contact for the user", func() {
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
		Attribute("ID", String, "db id")
	})
})

var CouncillorUpdate = Type("CouncillorUpdate", func() {
	Attribute("firstName", String)
	Attribute("secondName", String)
	Attribute("area", String)
	Attribute("county", String)
	Attribute("email", String)
	Attribute("inOffice", Boolean, "whether the councillor is still in office", func() {
		Default(false)
	})
	Attribute("phone", String, "a phone contact for the user", func() {
		Default("")
	})
	Attribute("address", String, "a phone contact for the user", func() {
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
	Required("firstName", "secondName", "area", "phone", "email", "party", "address", "county")
})
