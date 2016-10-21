package main

import (
	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/pkg/errors"
)

// AdminController implements the admin resource.
type AdminController struct {
	*goa.Controller
}

// NewAdminController creates a admin controller.
func NewAdminController(service *goa.Service) *AdminController {
	return &AdminController{Controller: service.NewController("AdminController")}
}

// CreateCouncillor runs the createCouncillor action.
func (c *AdminController) CreateCouncillor(ctx *app.CreateCouncillorAdminContext) error {
	var (
		actor      = ctx.Value("actor").(domain.Actor)
		cRepo      = domain.NewCouncillorRepo(config.Conf, actor, domain.AuthorisationService{})
		uRepo      = domain.NewUserRepo(config.Conf, actor, domain.AuthorisationService{})
		councillor = &app.GoaLocalCouncillor{}
		user       = &app.User{}
	)
	uploaded, f, err := ctx.FormFile("file")
	if err != nil {
		return errors.Wrap(err, "failed to parse file")
	}
	uploadService := domain.UploadService{
		Config: config.Conf,
	}
	imagePath, err := uploadService.Upload(f.Filename, uploaded)
	if err != nil {
		return err
	}
	area := ctx.FormValue("area")
	firstName := ctx.FormValue("firstName")
	secondName := ctx.FormValue("seconName")
	email := ctx.FormValue("email")
	party := ctx.FormValue("party")
	phone := ctx.FormValue("phone")
	web := ctx.FormValue("party")
	address := ctx.FormValue("address")
	county := ctx.FormValue("county")
	twitterHandler := ctx.FormValue("twitter")
	facebookName := ctx.FormValue("facebook")
	user.Area = area
	user.County = county
	user.Type = "councillor"
	user.Email = email
	user.Active = false
	user.FirstName = firstName
	user.SecondName = secondName
	user.SignupType = "local"
	du := domain.NewUser(user)
	if err := uRepo.SaveUpdate(du); err != nil {
		return err
	}

	councillor.Area = area
	councillor.FirstName = firstName
	councillor.SecondName = secondName
	councillor.Image = imagePath
	councillor.Email = email
	councillor.Address = address
	councillor.Party = party
	councillor.Phone = phone
	councillor.Web = web
	councillor.Twitter = &twitterHandler
	councillor.Facebook = &facebookName
	councillor.UserID = du.ID

	if err := councillor.Validate(); err != nil {
		return err
	}
	if err := cRepo.SaveUpdate(domain.NewCouncillor(councillor)); err != nil {
		return err
	}

	return ctx.Created()
}
