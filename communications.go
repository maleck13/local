package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/domain/communication"
	"github.com/pkg/errors"
)

// CommunicationsController implements the communications resource.
type CommunicationsController struct {
	*goa.Controller
}

func communicationToView(dc *domain.Communication) *app.GoaLocalCommunication {
	return &app.GoaLocalCommunication{
		Body:        dc.Body,
		Subject:     dc.Subject,
		Sent:        dc.Sent,
		ID:          dc.ID,
		IsPrivate:   dc.IsPrivate,
		Open:        dc.Open,
		From:        dc.From,
		To:          dc.To,
		CommID:      dc.CommID,
		RecepientID: dc.RecepientID,
	}
}

// NewCommunicationsController creates a communications controller.
func NewCommunicationsController(service *goa.Service) *CommunicationsController {
	return &CommunicationsController{Controller: service.NewController("CommunicationsController")}
}

// RecieveEmail runs the recieveEmail action.
func (c *CommunicationsController) RecieveEmail(ctx *app.RecieveEmailCommunicationsContext) error {
	// CommunicationsController_RecieveEmail: start_implement
	fmt.Println("content type ", ctx.Request.Header.Get("content-type"))
	body := ctx.Request.Body
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return errors.Wrap(err, "error reading body")
	}

	fmt.Println("recieved ", string(data))
	// Put your logic here

	// CommunicationsController_RecieveEmail: end_implement
	return ctx.OK(nil)
}

// Send runs the send action.
func (c *CommunicationsController) Send(ctx *app.SendCommunicationsContext) error {
	actor := ctx.Value("actor").(domain.Actor)
	authorisor := domain.NewAuthorisationService()
	localsRepo := domain.NewUserRepo(config.Conf, actor, authorisor)
	commsRepo := domain.NewCommunicationRepo(config.Conf, actor, authorisor)
	if !("local" == actor.Type() || "councillor" == actor.Type()) {
		return goa.ErrBadRequest("invalid actor type ")
	}
	user, err := localsRepo.FindOneByFieldAndValue("id", actor.Id())
	if err != nil {
		return err
	}
	if user == nil {
		return goa.ErrNotFound("user not found")
	}
	recipient, err := localsRepo.FindOneByFieldAndValue("id", ctx.Payload.RecepientID)
	if err != nil {
		return err
	}
	if nil == recipient {
		return goa.ErrNotFound("councillor not found")
	}
	sender := communication.NewEmailer(user.FullName(), user.Email)
	reciever := communication.NewEmailer(recipient.FullName(), recipient.Email)
	if actor.Type() == "councillor" {
		sender = communication.NewEmailer(recipient.FullName(), recipient.Email)
		reciever = communication.NewEmailer(user.FullName(), user.Email)
	}
	message := communication.NewEmailMessage(ctx.Payload.Body, ctx.Payload.Subject, "", "communication")
	comm := domain.NewCommunication(ctx.Payload)
	commsService := communication.NewService(config.Conf, commsRepo)
	if err := commsService.Send(communication.Email, sender, reciever, message); err != nil {
		return err
	}
	now := time.Now()
	comm.Sent = &now
	senderMail := sender.Address()
	comm.From = &senderMail
	recieptMail := reciever.Address()
	comm.To = &recieptMail
	comm.UserID = user.ID
	if err := commsRepo.SaveUpdate(comm); err != nil {
		return err
	}
	return ctx.OK(communicationToView(comm))
}

// Close closes the communication
func (c *CommunicationsController) Close(ctx *app.CloseCommunicationsContext) error {
	return ctx.OK(nil)
}

// List lists the communications based on the requesting user for the given councillor
func (c *CommunicationsController) List(ctx *app.ListCommunicationsContext) error {
	actor := ctx.Value("actor").(domain.Actor)
	authorisor := domain.AuthorisationService{}
	commsRepo := domain.NewCommunicationRepo(config.Conf, actor, authorisor)
	var comms = []*domain.Communication{}
	var err error
	res := app.GoaLocalCommunicationCollection{}
	//TODO need to group these communications together to form replies etc
	if ctx.CommID == nil {
		comms, err = commsRepo.FindAllByRecepientIDAndUserID(ctx.Rid, actor.Id(), true)
		if err != nil {
			return err
		}
		for _, c := range comms {
			res = append(res, communicationToView(c))
		}
		return ctx.OK(res)
	}
	q := map[string]interface{}{"UserID": actor.Id(), "RecipientID": ctx.Rid, "CommID": ctx.CommID}
	comms, err = commsRepo.FindAllByKeyValues(q)
	if err != nil {
		return err
	}
	for _, c := range comms {
		res = append(res, communicationToView(c))
	}
	return ctx.OK(res)
}
