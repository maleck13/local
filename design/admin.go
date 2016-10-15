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
		Response(Created)
		Response(Unauthorized) // of HTTP responses.
	})
})
