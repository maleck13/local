package design

import (
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

var _ = Resource("swagger", func() {
	Description("The swagger definition")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})
