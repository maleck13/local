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

// CloseCommunicationsContext provides the communications close action context.
type CloseCommunicationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewCloseCommunicationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the communications controller close action.
func NewCloseCommunicationsContext(ctx context.Context, service *goa.Service) (*CloseCommunicationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CloseCommunicationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CloseCommunicationsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CloseCommunicationsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CloseCommunicationsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCommunicationsContext provides the communications list action context.
type ListCommunicationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Cid string
}

// NewListCommunicationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the communications controller list action.
func NewListCommunicationsContext(ctx context.Context, service *goa.Service) (*ListCommunicationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListCommunicationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCid := req.Params["cid"]
	if len(paramCid) > 0 {
		rawCid := paramCid[0]
		rctx.Cid = rawCid
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCommunicationsContext) OK(r GoaLocalCommunicationCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.communication+json; type=collection")
	if r == nil {
		r = GoaLocalCommunicationCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListCommunicationsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCommunicationsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// RecieveEmailCommunicationsContext provides the communications recieveEmail action context.
type RecieveEmailCommunicationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewRecieveEmailCommunicationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the communications controller recieveEmail action.
func NewRecieveEmailCommunicationsContext(ctx context.Context, service *goa.Service) (*RecieveEmailCommunicationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RecieveEmailCommunicationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RecieveEmailCommunicationsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *RecieveEmailCommunicationsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *RecieveEmailCommunicationsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// SendCommunicationsContext provides the communications send action context.
type SendCommunicationsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Communication
}

// NewSendCommunicationsContext parses the incoming request URL and body, performs validations and creates the
// context used by the communications controller send action.
func NewSendCommunicationsContext(ctx context.Context, service *goa.Service) (*SendCommunicationsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := SendCommunicationsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SendCommunicationsContext) OK(r *GoaLocalCommunication) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.communication+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SendCommunicationsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *SendCommunicationsContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListForCountyAndAreaCouncillorsContext provides the councillors listForCountyAndArea action context.
type ListForCountyAndAreaCouncillorsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Area   string
	County string
}

// NewListForCountyAndAreaCouncillorsContext parses the incoming request URL and body, performs validations and creates the
// context used by the councillors controller listForCountyAndArea action.
func NewListForCountyAndAreaCouncillorsContext(ctx context.Context, service *goa.Service) (*ListForCountyAndAreaCouncillorsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListForCountyAndAreaCouncillorsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramArea := req.Params["area"]
	if len(paramArea) > 0 {
		rawArea := paramArea[0]
		rctx.Area = rawArea
	}
	paramCounty := req.Params["county"]
	if len(paramCounty) > 0 {
		rawCounty := paramCounty[0]
		rctx.County = rawCounty
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListForCountyAndAreaCouncillorsContext) OK(r GoaLocalCouncillorCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.councillor+json; type=collection")
	if r == nil {
		r = GoaLocalCouncillorCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListForCountyAndAreaCouncillorsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// ReadByIDCouncillorsContext provides the councillors readById action context.
type ReadByIDCouncillorsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewReadByIDCouncillorsContext parses the incoming request URL and body, performs validations and creates the
// context used by the councillors controller readById action.
func NewReadByIDCouncillorsContext(ctx context.Context, service *goa.Service) (*ReadByIDCouncillorsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ReadByIDCouncillorsContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ReadByIDCouncillorsContext) OK(r *GoaLocalCouncillor) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.councillor+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ReadByIDCouncillorsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
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
	if r == nil {
		r = GoaLocalUserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLogin sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKLogin(r GoaLocalUserLoginCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json; type=collection")
	if r == nil {
		r = GoaLocalUserLoginCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKPublic sends a HTTP response with status code 200.
func (ctx *ListUserContext) OKPublic(r GoaLocalUserPublicCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.local.user+json; type=collection")
	if r == nil {
		r = GoaLocalUserPublicCollection{}
	}
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
