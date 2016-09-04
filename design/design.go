package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access") // Define "api:access" scope
})

var _ = API("locals", func() { // API defines the microservice endpoint and
	Title("You and local government")                                       // other global properties. There should be one
	Description("A platform for interacting with you and local government") // and exactly one API definition appearing in
	Scheme("http")                                                          // the design.
	Host("localhost:3000")
	Origin("*", func() {
		Methods("GET", "POST", "DELETE", "PUT")
		Headers("x-auth")
	})
	Produces("application/json")
})

// User defines the media type used to render users. The mediaType gives all fields and the views break it down to specific responses
var User = MediaType("application/vnd.goa.local.user+json", func() {
	Description("A User of locals")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", String, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("firstName", String, "Name of the user")
		Attribute("secondName", String, "Name of the user")
		Attribute("signupType", String, "the signupType of user google local")
		Attribute("token", String, "This can be an oauth token or a password")
		Attribute("email", String, "The email of the user")
		Attribute("area", String, "The area of the users local council")
		Attribute("loginExpires", Integer, "The area of the users local council")
		Attribute("location", Location, "The area of the users local council")
		Attribute("type", String, func() {
			Description("the type of user admin local councillor")
			Default("local")
		})
		Attribute("image", String, "an image url for the user")
		Required("firstName", "secondName", "email", "token")
	})
	View("full", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("firstName")
		Attribute("secondName")
		Attribute("signupType")
		Attribute("email")
		Attribute("area")
		Attribute("location", Location)
		Attribute("loginExpires")
		Attribute("type")
		Attribute("token")
	})

	View("login", func() {
		Attribute("token")
		Attribute("id") //users id
		Attribute("loginExpires")
	})

	View("default", func() { // View defines a rendering of the media type. this is used to return to the client
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("firstName")
		Attribute("secondName")
		Attribute("signupType")
		Attribute("email")
		Attribute("area")
		Attribute("type")
	})
	//visible to everyone
	View("public", func() {
		Attribute("id")
		Attribute("firstName")
		Attribute("area")
	})
})

//User handler definition host:3001/user
var _ = Resource("user", func() {
	BasePath("/user") // together. They map to REST resources for REST
	DefaultMedia(User, "default")
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:access") // Enforce presence of "api" scope in JWT claims.
	})

	Action("create", func() { // Actions define a single API endpoint together
		Description("Signup a user") // with its path, parameters (both path
		Routing(POST("/signup"))     // parameters and querystring values) and payload
		Payload(UserPayload)
		Response(Created)  // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
		NoSecurity()
	})
	Action("list", func() {
		Description("get a list user")
		Routing(GET("/"))
		Response(OK, func() {
			Media(CollectionOf(User))
		})
		Response(Unauthorized)
		Response(NotFound)
	})
	Action("read", func() {
		Description("get a user")
		Routing(GET("/:id"))
		Params(func() { // (shape of the request body).
			Param("id", String, "user ID")
		})
		Response(OK)
		Response(Unauthorized)
		Response(NotFound)
	})
	Action("delete", func() {
		Description("delete a user")
		Routing(DELETE("/:id"))
		Params(func() { // (shape of the request body).
			Param("id", String, "user ID")
		})
		Response(Accepted)
		Response(Unauthorized)
		Response(NotFound)
	})
	Action("update", func() {
		Description("update a user")
		Routing(POST("/:id"))
		Payload(UserPayload)
		Params(func() { // (shape of the request body).
			Param("id", String, "user ID")
		})
		Response(OK)
		Response(Unauthorized)
		Response(NotFound)
	})

	Action("login", func() {
		Description("login user")
		Routing(POST("login"))
		Payload(Login)
		Response(OK, func() {
			Media(User, "login")
		})
		NoSecurity()
	})

})

// Below are custom types to be used in the api. They can be referenced in the Action as Payload type see login
var Location = Type("Location", func() {
	Attribute("Lon", Number)
	Attribute("Lat", Number)
})

var Login = Type("Login", func() {
	Attribute("email", String)
	Attribute("token", String)
	Attribute("signupType", String)
	Required("email", "token", "signupType")
})

var UserPayload = Type("User", func() {
	Attribute("id", String, "Unique user ID")
	Attribute("firstName", String, "Name of the user")
	Attribute("secondName", String, "Name of the user")
	Attribute("token", String, "This can be an oauth token or a password")
	Attribute("email", String, "The email of the user")
	Attribute("signupType", String, "the signupType of user google local", func() {
		Default("")
	})
	Attribute("area", String, "The area of the users local council", func() {
		Default("")
	})
	Attribute("location", Location, "The area of the users local council")
	Attribute("type", String, func() {
		Description("the type of user admin local councillor")
		Default("local")
	})
	Attribute("image", String, "an image url for the user", func() {
		Default("")
	})
	Required("firstName", "secondName", "email", "token")
})
