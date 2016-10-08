package main

import (
	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/domain/local"
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
	var (
		user *app.User
		err  error
		// userRepo is configured with AdminActor to allow anonymous write
		userRepo     = domain.NewUserRepo(config.Conf, domain.NewAdminActor(), local.AuthorisationService{})
		localService = local.NewService(config.Conf, userRepo)
	)
	user, err = localService.AddProviderData(ctx.Payload)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	if _, err := localService.Register(user); err != nil {
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
	//setup an admin user repo to allow write acces to anonymous user
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), local.AuthorisationService{})
	localService := local.NewService(config.Conf, userRepo)
	token, err := localService.Authenticate(ctx.Payload.SignupType, ctx.Payload.Token, ctx.Payload.Email)
	if err != nil {
		return err
	}
	user, err := userRepo.FindOneByFieldAndValue("Email", ctx.Payload.Email)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Header().Add("Bearer", token)
	userLogin := &app.GoaLocalUserLogin{Token: token, ID: &user.ID}
	return ctx.OKLogin(userLogin)
}

// Read runs the read action.
func (c *UserController) Read(ctx *app.ReadUserContext) error {
	// UserController_Read: start_implement
	var access = local.AuthorisationService{}
	actor := ctx.Value("actor").(domain.Actor)
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
		Type:       user.Type(),
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement
	var access = local.AuthorisationService{}
	actor := ctx.Value("actor").(domain.Actor)
	userRepo := domain.UserRepo{Config: config.Conf, Actor: actor, Authorisor: access}
	localService := local.NewService(config.Conf, userRepo)
	user, err := localService.Update(ctx.Payload)
	if err != nil {
		return err
	}
	res := &app.GoaLocalUser{
		Area:       &user.Area,
		Email:      user.Email,
		ID:         &user.ID,
		SignupType: &user.SignupType,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Type:       user.Type(),
	}
	return ctx.OK(res)
}
