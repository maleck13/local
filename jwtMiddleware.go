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
// and validates its content.
func NewJWTMiddleware(conf *config.Config) goa.Middleware {
	key := conf.Jwt.Secret
	securityContext := app.NewJWTSecurity()
	securityContext.Scopes["admin:access"] = "admin access"
	middleware := gJwt.New(key, JwtActorMiddeware(conf), securityContext)
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
