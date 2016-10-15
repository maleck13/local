package main

import (
	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
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
