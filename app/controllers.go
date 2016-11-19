//************************************************************************//
// API "locals": Application Controllers
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

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AdminController is the controller interface for the Admin actions.
type AdminController interface {
	goa.Muxer
	CreateCouncillor(*CreateCouncillorAdminContext) error
}

// MountAdminController "mounts" a Admin resource controller on the given service.
func MountAdminController(service *goa.Service, ctrl AdminController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/admin/councillor", ctrl.MuxHandler("preflight", handleAdminOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateCouncillorAdminContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.CreateCouncillor(rctx)
	}
	h = handleAdminOrigin(h)
	h = handleSecurity("jwt", h, "admin:access")
	service.Mux.Handle("POST", "/admin/councillor", ctrl.MuxHandler("CreateCouncillor", h, nil))
	service.LogInfo("mount", "ctrl", "Admin", "action", "CreateCouncillor", "route", "POST /admin/councillor", "security", "jwt")
}

// handleAdminOrigin applies the CORS response headers corresponding to the origin.
func handleAdminOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// CommunicationsController is the controller interface for the Communications actions.
type CommunicationsController interface {
	goa.Muxer
	Close(*CloseCommunicationsContext) error
	List(*ListCommunicationsContext) error
	RecieveEmail(*RecieveEmailCommunicationsContext) error
	Send(*SendCommunicationsContext) error
}

// MountCommunicationsController "mounts" a Communications resource controller on the given service.
func MountCommunicationsController(service *goa.Service, ctrl CommunicationsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/communications/close/:id", ctrl.MuxHandler("preflight", handleCommunicationsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/communications/councillor/:rid", ctrl.MuxHandler("preflight", handleCommunicationsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/communications/email/recieve", ctrl.MuxHandler("preflight", handleCommunicationsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/communications/send", ctrl.MuxHandler("preflight", handleCommunicationsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCloseCommunicationsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Close(rctx)
	}
	h = handleCommunicationsOrigin(h)
	service.Mux.Handle("DELETE", "/communications/close/:id", ctrl.MuxHandler("Close", h, nil))
	service.LogInfo("mount", "ctrl", "Communications", "action", "Close", "route", "DELETE /communications/close/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCommunicationsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleCommunicationsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("GET", "/communications/councillor/:rid", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Communications", "action", "List", "route", "GET /communications/councillor/:rid", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRecieveEmailCommunicationsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.RecieveEmail(rctx)
	}
	h = handleCommunicationsOrigin(h)
	service.Mux.Handle("POST", "/communications/email/recieve", ctrl.MuxHandler("RecieveEmail", h, nil))
	service.LogInfo("mount", "ctrl", "Communications", "action", "RecieveEmail", "route", "POST /communications/email/recieve")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSendCommunicationsContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Communication)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Send(rctx)
	}
	h = handleCommunicationsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("POST", "/communications/send", ctrl.MuxHandler("Send", h, unmarshalSendCommunicationsPayload))
	service.LogInfo("mount", "ctrl", "Communications", "action", "Send", "route", "POST /communications/send", "security", "jwt")
}

// handleCommunicationsOrigin applies the CORS response headers corresponding to the origin.
func handleCommunicationsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalSendCommunicationsPayload unmarshals the request body into the context request data Payload field.
func unmarshalSendCommunicationsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &communication{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// CouncillorsController is the controller interface for the Councillors actions.
type CouncillorsController interface {
	goa.Muxer
	ListConstituents(*ListConstituentsCouncillorsContext) error
	ListForCountyAndArea(*ListForCountyAndAreaCouncillorsContext) error
	ReadByID(*ReadByIDCouncillorsContext) error
	Update(*UpdateCouncillorsContext) error
	UploadProfilePic(*UploadProfilePicCouncillorsContext) error
}

// MountCouncillorsController "mounts" a Councillors resource controller on the given service.
func MountCouncillorsController(service *goa.Service, ctrl CouncillorsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/councillors/:id/consituents", ctrl.MuxHandler("preflight", handleCouncillorsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/councillors", ctrl.MuxHandler("preflight", handleCouncillorsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/councillors/:id", ctrl.MuxHandler("preflight", handleCouncillorsOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/councillors/:id/image", ctrl.MuxHandler("preflight", handleCouncillorsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListConstituentsCouncillorsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.ListConstituents(rctx)
	}
	h = handleCouncillorsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("GET", "/councillors/:id/consituents", ctrl.MuxHandler("ListConstituents", h, nil))
	service.LogInfo("mount", "ctrl", "Councillors", "action", "ListConstituents", "route", "GET /councillors/:id/consituents", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListForCountyAndAreaCouncillorsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.ListForCountyAndArea(rctx)
	}
	h = handleCouncillorsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("GET", "/councillors", ctrl.MuxHandler("ListForCountyAndArea", h, nil))
	service.LogInfo("mount", "ctrl", "Councillors", "action", "ListForCountyAndArea", "route", "GET /councillors", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReadByIDCouncillorsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.ReadByID(rctx)
	}
	h = handleCouncillorsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("GET", "/councillors/:id", ctrl.MuxHandler("ReadByID", h, nil))
	service.LogInfo("mount", "ctrl", "Councillors", "action", "ReadByID", "route", "GET /councillors/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateCouncillorsContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CouncillorUpdate)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleCouncillorsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("POST", "/councillors/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateCouncillorsPayload))
	service.LogInfo("mount", "ctrl", "Councillors", "action", "Update", "route", "POST /councillors/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadProfilePicCouncillorsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.UploadProfilePic(rctx)
	}
	h = handleCouncillorsOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("POST", "/councillors/:id/image", ctrl.MuxHandler("UploadProfilePic", h, nil))
	service.LogInfo("mount", "ctrl", "Councillors", "action", "UploadProfilePic", "route", "POST /councillors/:id/image", "security", "jwt")
}

// handleCouncillorsOrigin applies the CORS response headers corresponding to the origin.
func handleCouncillorsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalUpdateCouncillorsPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateCouncillorsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &councillorUpdate{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger-ui/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/index.html", "route", "GET /swagger-ui/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Login(*LoginUserContext) error
	Read(*ReadUserContext) error
	Resetpassword(*ResetpasswordUserContext) error
	SignUpCouncillor(*SignUpCouncillorUserContext) error
	Signup(*SignupUserContext) error
	Update(*UpdateUserContext) error
	VerifySignup(*VerifySignupUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/user/:id", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/login", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/resetpassword", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/councillor/signup", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/signup", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/signup/verify", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("DELETE", "/user/:id", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Delete", "route", "DELETE /user/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("GET", "/user", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "List", "route", "GET /user", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Login)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Login(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/user/login", ctrl.MuxHandler("Login", h, unmarshalLoginUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Login", "route", "POST /user/login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReadUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Read(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("GET", "/user/:id", ctrl.MuxHandler("Read", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Read", "route", "GET /user/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewResetpasswordUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ResetpasswordUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Resetpassword(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h, "password:reset")
	service.Mux.Handle("POST", "/user/resetpassword", ctrl.MuxHandler("Resetpassword", h, unmarshalResetpasswordUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Resetpassword", "route", "POST /user/resetpassword", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSignUpCouncillorUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SignUpCouncillorUserPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.SignUpCouncillor(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/user/councillor/signup", ctrl.MuxHandler("SignUpCouncillor", h, unmarshalSignUpCouncillorUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "SignUpCouncillor", "route", "POST /user/councillor/signup")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSignupUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*User)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Signup(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/user/signup", ctrl.MuxHandler("Signup", h, unmarshalSignupUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Signup", "route", "POST /user/signup")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUser)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleUserOrigin(h)
	h = handleSecurity("jwt", h, "api:access")
	service.Mux.Handle("POST", "/user/:id", ctrl.MuxHandler("Update", h, unmarshalUpdateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Update", "route", "POST /user/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVerifySignupUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.VerifySignup(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("GET", "/user/signup/verify", ctrl.MuxHandler("VerifySignup", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "VerifySignup", "route", "GET /user/signup/verify")
}

// handleUserOrigin applies the CORS response headers corresponding to the origin.
func handleUserOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalLoginUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalLoginUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &login{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalResetpasswordUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalResetpasswordUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &resetpasswordUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalSignUpCouncillorUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalSignUpCouncillorUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &signUpCouncillorUserPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalSignupUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalSignupUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &user{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUser{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
