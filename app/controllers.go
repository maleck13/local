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

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Create(*CreateUserContext) error
	Delete(*DeleteUserContext) error
	List(*ListUserContext) error
	Login(*LoginUserContext) error
	Read(*ReadUserContext) error
	Update(*UpdateUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/user/signup", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/:id", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/user/login", ctrl.MuxHandler("preflight", handleUserOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateUserContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*User)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleUserOrigin(h)
	service.Mux.Handle("POST", "/user/signup", ctrl.MuxHandler("Create", h, unmarshalCreateUserPayload))
	service.LogInfo("mount", "ctrl", "User", "action", "Create", "route", "POST /user/signup")

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

// unmarshalCreateUserPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateUserPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
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
