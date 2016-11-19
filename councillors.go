package main

import (
	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/domain/constituents"
	"github.com/maleck13/local/domain/councillor"
	"github.com/pkg/errors"
)

// CouncillorsController implements the councillors resource.
type CouncillorsController struct {
	*goa.Controller
}

// NewCouncillorsController creates a councillors controller.
func NewCouncillorsController(service *goa.Service) *CouncillorsController {
	return &CouncillorsController{Controller: service.NewController("CouncillorsController")}
}

// ListForCountyAndArea runs the ListForCountyAndArea action.
func (c *CouncillorsController) ListForCountyAndArea(ctx *app.ListForCountyAndAreaCouncillorsContext) error {
	// CouncillorsController_ListForUser: start_implement
	var (
		actor          = ctx.Value("actor").(domain.Actor)
		authorisor     = domain.AuthorisationService{}
		councillorRepo = domain.NewCouncillorRepo(config.Conf, actor, authorisor)
		county         = ctx.County
		area           = ctx.Area
	)
	// Put your logic here
	cs, err := councillorRepo.FindByCountyAndArea(county, area)
	if err != nil {
		return err
	}
	// CouncillorsController_ListForUser: end_implement
	res := app.GoaLocalCouncillorCollection{}
	for _, c := range cs {
		res = append(res, c.GoaLocalCouncillor)
	}
	return ctx.OK(res)
}

// ReadByID returns the councillor for the corrosponding ID
func (c *CouncillorsController) ReadByID(ctx *app.ReadByIDCouncillorsContext) error {
	var (
		actor          = ctx.Value("actor").(domain.Actor)
		authorisor     = domain.AuthorisationService{}
		councillorRepo = domain.NewCouncillorRepo(config.Conf, actor, authorisor)
		ID             = ctx.ID
	)
	counciller, err := councillorRepo.FindOneByKeyValue("id", ID)
	if err != nil {
		return err
	}
	return ctx.OK(counciller.GoaLocalCouncillor)

}

// Update updates a councillor model
func (c *CouncillorsController) Update(ctx *app.UpdateCouncillorsContext) error {
	var (
		actor             = ctx.Value("actor").(domain.Actor)
		authorisor        = domain.AuthorisationService{}
		councillorRepo    = domain.NewCouncillorRepo(config.Conf, actor, authorisor)
		councillorService = councillor.NewService(councillorRepo)
		ID                = ctx.ID
	)
	updated, err := councillorService.Update(ID, ctx.Payload)
	if err != nil {
		return err
	}
	return ctx.OK(updated.GoaLocalCouncillor)
}

// UploadProfilePic adds an image for a councillor
func (c *CouncillorsController) UploadProfilePic(ctx *app.UploadProfilePicCouncillorsContext) error {
	var (
		actor = ctx.Value("actor").(domain.Actor)
		cRepo = domain.NewCouncillorRepo(config.Conf, actor, domain.AuthorisationService{})
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
	councillor, err := cRepo.FindOneByKeyValue("id", ctx.ID)
	if err != nil {
		return err
	}
	councillor.Image = imagePath
	if err := cRepo.SaveUpdate(councillor); err != nil {
		return err
	}
	return ctx.OK(councillor.GoaLocalCouncillor)
}

// ListConstituents will return a list of constituents based on a counciller id
func (c *CouncillorsController) ListConstituents(ctx *app.ListConstituentsCouncillorsContext) error {
	var (
		actor        = ctx.Value("actor").(domain.Actor)
		cRepo        = domain.NewCouncillorRepo(config.Conf, actor, domain.AuthorisationService{})
		lRepo        = domain.NewUserRepo(config.Conf, actor, domain.AuthorisationService{})
		commsRepo    = domain.NewCommunicationRepo(config.Conf, actor, domain.AuthorisationService{})
		constService = constituents.Service{ConstituentFinder: lRepo, CouncillorFinder: cRepo, CommsFinder: commsRepo}
	)

	constituents, err := constService.ConsistuentsForCouncillor(ctx.ID)
	if err != nil {
		return err
	}
	return ctx.OKNocomms(constituents)
}

func parseCouncillorFormPost(ctx *app.UpdateCouncillorsContext) *app.GoaLocalCouncillor {
	area := ctx.FormValue("area")
	firstName := ctx.FormValue("firstName")
	secondName := ctx.FormValue("secondName")
	email := ctx.FormValue("email")
	party := ctx.FormValue("party")
	phone := ctx.FormValue("phone")
	web := ctx.FormValue("party")
	address := ctx.FormValue("address")
	county := ctx.FormValue("county")
	twitterHandler := ctx.FormValue("twitter")
	facebookName := ctx.FormValue("facebook")
	councillor := &app.GoaLocalCouncillor{}
	councillor.Area = area
	councillor.FirstName = firstName
	councillor.SecondName = secondName
	councillor.Email = email
	councillor.Address = address
	councillor.Party = party
	councillor.Phone = phone
	councillor.Web = web
	councillor.Twitter = &twitterHandler
	councillor.Facebook = &facebookName
	councillor.County = county
	return councillor
}
