package main

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/goadesign/goa"
	gJwt "github.com/goadesign/goa/middleware/security/jwt"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
func NewJWTMiddleware(conf *config.Config) goa.Middleware {
	key := conf.Jwt.Secret
	middleware := gJwt.New(key, JwtActorMiddeware(conf), app.NewJWTSecurity())
	return middleware
}

//JwtActorMiddeware uses the jwt token to set up an actor that can be used for Authorization purposes
func JwtActorMiddeware(conf *config.Config) goa.Middleware {
	actorMW := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			token := gJwt.ContextJWT(ctx)
			actor := domain.NewLocalActor(token)
			ctx = context.WithValue(ctx, "actor", actor)
			return h(ctx, rw, req)
		}
	}
	am, _ := goa.NewMiddleware(actorMW)
	return am
}
