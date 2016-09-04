package main

import (
	"github.com/goadesign/goa"
	gJwt "github.com/goadesign/goa/middleware/security/jwt"
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
	authFactory := &domain.AuthenticatorFactory{Config: config.Conf}

	authenticator, err := authFactory.Factory(ctx.Payload.SignupType)
	if err != nil {
		return errors.LogAndReturnError(err)

	}
	user, err := authenticator.Authenticate(ctx.Payload.Token, ctx.Payload.Email)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	jsonWT := domain.JWT{Config: config.Conf}
	token, err := jsonWT.CreateToken(user)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	ctx.ResponseWriter.Header().Add("Bearer", token)
	userLogin := &app.GoaLocalUserLogin{Token: token, ID: &user.ID}
	return ctx.OKLogin(userLogin)
}

// Read runs the read action.
func (c *UserController) Read(ctx *app.ReadUserContext) error {
	// UserController_Read: start_implement
	var access = domain.Access{}
	actor := ctx.Value("actor").(domain.AuthActor)
	userRepo := domain.UserRepo{Config: config.Conf, Actor: actor, Authorisor: access}
	user, err := userRepo.FindOneByFieldAndValue("id", ctx.Params.Get("id"))
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	if user == nil {
		return ctx.NotFound()
	}
	if err := access.Authorise(user, "read", actor); err != nil {
		return err
	}
	// UserController_Read: end_implement
	res := &app.GoaLocalUser{
		Area:       &user.Area,
		Email:      user.Email,
		ID:         &user.ID,
		SignupType: &user.SignupType,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	token := gJwt.ContextJWT(ctx)
	if nil == token {
		return ctx.Unauthorized()
	}

	res := &app.GoaLocalUser{}
	return ctx.OK(res)
}
