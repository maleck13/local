package main

import (
	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/errors"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	signUpService := domain.NewSignUpFactory(config.Conf)
	signup, err := signUpService.Factory(ctx.Payload.SignupType)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	_, err = signup.Register(ctx.Payload)
	if err != nil {
		return errors.LogAndReturnError(err)
	}

	ctx.Created()
	// SignupController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {

	// UserController_Delete: start_implement

	// Put your logic here

	// UserController_Delete: end_implement

	return ctx.Accepted()
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here

	// UserController_List: end_implement
	res := app.GoaLocalUserCollection{}
	return ctx.OK(res)
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	authFactory := domain.NewAuthenticatorFactory(config.Conf)

	authenticator, err := authFactory.Factory(ctx.Payload.SignupType)
	if err != nil {
		return errors.LogAndReturnError(err)

	}
	if err := authenticator.Authenticate(ctx.Payload.Token, ctx.Payload.Email); err != nil {
		return errors.LogAndReturnError(err)
	}

	return ctx.NoContent()
}

// Read runs the read action.
func (c *UserController) Read(ctx *app.ReadUserContext) error {
	// UserController_Read: start_implement

	// Put your logic here

	// UserController_Read: end_implement
	res := &app.GoaLocalUser{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here

	// UserController_Update: end_implement
	res := &app.GoaLocalUser{}
	return ctx.OK(res)
}
