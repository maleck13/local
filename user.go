package main

//TODO refactor this test to be more like the councillors test

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/domain/communication"
	"github.com/maleck13/local/domain/councillor"
	"github.com/maleck13/local/domain/local"
	"github.com/maleck13/local/errors"
)

func domainUserToLocalUser(user *domain.User) *app.GoaLocalUser {
	return &app.GoaLocalUser{
		Area:       &user.Area,
		Email:      user.Email,
		ID:         &user.ID,
		SignupType: &user.SignupType,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Type:       user.Type(),
		County:     user.County,
	}
}

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
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
	var (
		query    = ctx.Request.URL.Query()
		userType = query.Get("type")
		area     = query.Get("area")
		access   = domain.AuthorisationService{}
		actor    = ctx.Value("actor").(domain.Actor)
	)
	if userType == "" || area == "" {
		return goa.ErrBadRequest("expected a user type and area", 400)
	}
	userRepo := domain.UserRepo{Config: config.Conf, Actor: actor, Authorisor: access}
	domainUsers, err := userRepo.FindAllByTypeAndArea(userType, area)
	if err != nil {
		return err
	}
	// UserController_List: end_implement
	res := app.GoaLocalUserCollection{}
	for _, du := range domainUsers {
		res = append(res, domainUserToLocalUser(du))
	}
	return ctx.OK(res)
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	//setup an admin user repo to allow write acces to anonymous user
	var (
		userRepo    = domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
		authService = domain.NewAuthenticateService(ctx.Payload.SignupType, config.Conf, userRepo)
	)

	if err := authService.Authenticate(ctx.Payload.Token, ctx.Payload.Email); err != nil {
		return err
	}
	user, err := userRepo.FindOneByFieldAndValue("Email", ctx.Payload.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return goa.ErrNotFound("no such user")
	}
	token, err := authService.CreateToken(user.ID, user.Email, user.Type(), nil)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Header().Add("Bearer", token)
	userLogin := &app.GoaLocalUserLogin{Token: token, ID: &user.ID, Type: user.Type(), Status: true, County: user.County, Area: &user.Area}
	return ctx.OKLogin(userLogin)
}

// Read runs the read action.
func (c *UserController) Read(ctx *app.ReadUserContext) error {
	// UserController_Read: start_implement
	var access = domain.AuthorisationService{}
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
	res := domainUserToLocalUser(user)
	return ctx.OK(res)
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement
	var access = domain.AuthorisationService{}
	actor := ctx.Value("actor").(domain.Actor)
	userRepo := domain.UserRepo{Config: config.Conf, Actor: actor, Authorisor: access}
	localService := local.NewService(config.Conf, userRepo)
	user, err := localService.Update(ctx.Payload)
	if err != nil {
		return err
	}
	res := domainUserToLocalUser(user)
	return ctx.OK(res)
}

// Signup adds a new user from a signup
func (c *UserController) Signup(ctx *app.SignupUserContext) error {
	var (
		user *app.User
		err  error
		// userRepo is configured with AdminActor to allow anonymous write
		userRepo     = domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
		localService = local.NewService(config.Conf, userRepo)
	)
	user, err = localService.AddProviderData(ctx.Payload)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	if _, err := localService.Register(user); err != nil {
		return errors.LogAndReturnError(err)
	}
	return ctx.Created()
}

// SignUpCouncillor checks if a councillor exists with the given email and then sends out a verification email to that address
func (c *UserController) SignUpCouncillor(ctx *app.SignUpCouncillorUserContext) error {
	var (
		actor        = domain.NewAdminActor()
		authorisor   = domain.AuthorisationService{}
		userRepo     = domain.NewUserRepo(config.Conf, actor, authorisor)
		commsRepo    = domain.NewCommunicationRepo(config.Conf, actor, authorisor)
		localService = local.NewService(config.Conf, userRepo)
		commsService = communication.NewService(config.Conf, commsRepo)
		authService  = domain.NewAuthenticateService("local", config.Conf, userRepo)
	)

	exists, err := localService.CheckCouncillorExists(ctx.Payload.Email)
	if err != nil {
		return err
	}
	if nil == exists {
		return ctx.NotFound()
	}

	//create temp token
	token, err := authService.CreateToken(exists.ID, exists.Email, "councillor", func(c jwt.MapClaims) jwt.Claims {
		c["exp"] = time.Now().Add(time.Minute * 15).Unix()
		c["scopes"] = "password:reset"
		return c
	})
	if err != nil {
		return err
	}
	sender := communication.NewEmailer("locals", "noreply@locals.ie")
	reciever := communication.NewEmailer("Councillor", ctx.Payload.Email)
	message := communication.NewEmailMessage("Locals.ie verification", config.Conf.SiteHost+"user/signup/verify?key="+token+"&uid="+exists.ID+"", "", "verification")
	if err := commsService.Send(communication.Email, sender, reciever, message); err != nil {
		return err
	}
	return ctx.OK(nil)

}

// VerifySignup uses a link with a key in it sent from a mail to verify the signup is valid and actives the associated user.
func (c *UserController) VerifySignup(ctx *app.VerifySignupUserContext) error {
	var (
		actor             = domain.NewAdminActor()
		authorisor        = domain.AuthorisationService{}
		userRepo          = domain.NewUserRepo(config.Conf, actor, authorisor)
		councillorRepo    = domain.NewCouncillorRepo(config.Conf, actor, authorisor)
		authService       = domain.NewAuthenticateService("local", config.Conf, userRepo)
		localService      = local.NewService(config.Conf, userRepo)
		councillorService = councillor.NewService(councillorRepo)
	)
	key := *ctx.Key
	uid := *ctx.UID
	if "" == key || "" == uid {
		return goa.ErrUnauthorized("missing key")
	}

	if err := authService.VerifyVerificationToken(key, uid, "password:reset"); err != nil {
		return err
	}

	if err := localService.ActivateUser(uid); err != nil {
		return err
	}

	if actor.Type() == "councillor" {
		user, err := userRepo.FindOneByFieldAndValue("id", uid)
		if err != nil {
			return err
		}
		if nil == user {
			return goa.ErrNotFound("failed to find user")
		}
		cllr, err := councillorRepo.FindOneByKeyValue("Email", user.Email)
		if err != nil {
			return err
		}
		if nil == user {
			return goa.ErrNotFound("failed to find cllr")
		}
		if err := councillorService.SetCouncillorUID(cllr.ID, uid); err != nil {
			return err
		}
	}

	http.Redirect(ctx.ResponseWriter, ctx.Request, config.Conf.SiteHost+"passwordreset?key="+key+"&uid="+uid, http.StatusSeeOther)
	return nil
}

// Resetpassword resets a users password. If we get here the user has a valid JWT token with the reset:password scope. So we can proceed using the actor and reset the password
func (c *UserController) Resetpassword(ctx *app.ResetpasswordUserContext) error {
	var (
		actor       = ctx.Value("actor").(domain.Actor)
		authorisor  = domain.AuthorisationService{}
		userRepo    = domain.NewUserRepo(config.Conf, actor, authorisor)
		authService = domain.NewAuthenticateService("local", config.Conf, userRepo)
	)
	if err := authService.ResetPassword(actor.Id(), ctx.Payload.Newpassword); err != nil {
		return err
	}
	return nil
}
