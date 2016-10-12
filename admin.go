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
		councillor = &app.Councillor{}
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
	councillor.Area = ctx.FormValue("area")
	councillor.FirstName = ctx.FormValue("firstName")
	councillor.SecondName = ctx.FormValue("secondName")
	councillor.Image = imagePath
	councillor.Email = ctx.FormValue("email")
	councillor.Party = ctx.FormValue("party")
	councillor.Phone = ctx.FormValue("phone")
	councillor.Web = ctx.FormValue("web")
	twitterHandler := ctx.FormValue("twitter")
	councillor.Twitter = &twitterHandler
	facebookName := ctx.FormValue("facebook")
	councillor.Facebook = &facebookName
	if err := councillor.Validate(); err != nil {
		return err
	}
	if err := cRepo.SaveUpdate(domain.NewCouncillor(councillor)); err != nil {
		return err
	}

	return ctx.Created()
}
