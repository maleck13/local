//************************************************************************//
// API "locals": Application Contexts
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
	"golang.org/x/net/context"
)

// CreateCouncillorAdminContext provides the admin createCouncillor action context.
type CreateCouncillorAdminContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewCreateCouncillorAdminContext parses the incoming request URL and body, performs validations and creates the
// context used by the admin controller createCouncillor action.
func NewCreateCouncillorAdminContext(ctx context.Context, service *goa.Service) (*CreateCouncillorAdminContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateCouncillorAdminContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateCouncillorAdminContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateCouncillorAdminContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// CreateUserContext provides the user create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *User
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller create action.
func NewCreateUserContext(ctx context.Context, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// DeleteUserContext provides the user delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller delete action.
func NewDeleteUserContext(ctx context.Context, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// Accepted sends a HTTP response with status code 202.
func (ctx *DeleteUserContext) Accepted() error {
	ctx.ResponseData.WriteHeader(202)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *DeleteUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListUserContext provides the user list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller list action.
func NewListUserContext(ctx context.Context, service *goa.Service) (*ListUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r GoaLocalUserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKFull sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKFull(r GoaLocalUserFullCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLogin sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKLogin(r GoaLocalUserLoginCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKPublic sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKPublic(r GoaLocalUserPublicCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// LoginUserContext provides the user login action context.
type LoginUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Login
}

// NewLoginUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller login action.
func NewLoginUserContext(ctx context.Context, service *goa.Service) (*LoginUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := LoginUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OKLogin sends a HTTP response with status code 200.
func (ctx *LoginUserContext) OKLogin(r *GoaLocalUserLogin) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ReadUserContext provides the user read action context.
type ReadUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewReadUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller read action.
func NewReadUserContext(ctx context.Context, service *goa.Service) (*ReadUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ReadUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ReadUserContext) OK(r *GoaLocalUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ReadUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ReadUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// SignupUserContext provides the user signup action context.
type SignupUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *User
}

// NewSignupUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller signup action.
func NewSignupUserContext(ctx context.Context, service *goa.Service) (*SignupUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SignupUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *SignupUserContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *SignupUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateUserContext provides the user update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *UpdateUser
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller update action.
func NewUpdateUserContext(ctx context.Context, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *GoaLocalUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdateUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
