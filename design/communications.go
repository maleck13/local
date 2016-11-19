package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
) // User defines the media type used to render users. The mediaType gives all fields and the views break it down to specific responses

var _ = Resource("communications", func() {
	BasePath("/communications") // together. They map to REST resources for REST
	Action("recieveEmail", func() {
		Description("recieve an email")
		Routing(POST("/email/recieve"))
		Response(OK)
		Response(Unauthorized)
		Response(InternalServerError)
	})
	Action("send", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Description("send and email ")
		Routing(POST("/send"))
		Payload(CommunicationPayload)
		Response(OK, func() {
			Media(Communication)
		})
		Response(Unauthorized)
		Response(InternalServerError)
	})
	Action("list", func() {
		Security(JWT, func() { // Use JWT to auth requests to this endpoint
			Scope("api:access") // Enforce presence of "api" scope in JWT claims.
		})
		Description("read communications ")
		Routing(GET("/councillor/:rid"))
		Params(func() { // (shape of the request body).
			Param("rid", String, "recepientID")
			Param("commID", String, "communication id")
		})
		Response(OK, func() {
			Media(CollectionOf(Communication))
		})
		Response(Unauthorized)
		Response(InternalServerError)
	})
	Action("close", func() {
		Description("recieve an email")
		Routing(DELETE("/close/:id"))
		Params(func() { // (shape of the request body).
			Param("id", String)
		})
		Response(OK)
		Response(Unauthorized)
		Response(InternalServerError)
	})
})

var Communication = MediaType("application/vnd.goa.local.communication+json", func() {
	Description("An communication")
	Reference(CommunicationPayload)
	Attributes(func() {
		Attribute("isPrivate")
		Attribute("recepientID")
		Attribute("subject")
		Attribute("body")
		Attribute("from")
		Attribute("to")
		Attribute("open")
		Attribute("type")
		Attribute("id")
		Attribute("userID", String)
		Attribute("error")
		Attribute("sent")
		Attribute("commID", String)
	})
	View("default", func() {
		Attribute("isPrivate")
		Attribute("recepientID")
		Attribute("userID")
		Attribute("subject")
		Attribute("body")
		Attribute("from")
		Attribute("to")
		Attribute("type")
		Attribute("open")
		Attribute("id")
		Attribute("sent")
		Attribute("commID")
	})
	Required("recepientID", "subject", "body", "isPrivate")
})

var CommunicationPayload = Type("Communication", func() {
	Attribute("id", String, "db id", func() {
		Metadata("struct:tag:gorethink", "id,omitempty")
		Metadata("struct:tag:json", "id,omitempty")
		Metadata("struct:tag:form", "id,omitempty")
		Default("")
	})
	Attribute("isPrivate", Boolean, func() {
		Default(true)
	})
	Attribute("recepientID", String)
	Attribute("subject", String)
	Attribute("body", String)
	Attribute("from", String)
	Attribute("to", String)
	Attribute("open", Boolean, func() {
		Default(true)
	})
	Attribute("type", String)
	Attribute("error", String, func() {
		Default("")
	})
	Attribute("sent", DateTime)
	Attribute("commID", String)
	Required("recepientID", "subject", "body", "isPrivate", "type")
})
