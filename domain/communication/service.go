/**
communication agregates together both locals and councillors in order to orchestrate communication between the two parties.

**/

package communication

import (
	"errors"
	"net/http"

	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
)

// Sender is someone looking to communicate
type Sender interface {
	// From will be something like an email address
	Address() string
	Name() string
}

// Reciever is someone recieving communication
type Reciever interface {
	// Address can be something like an email address
	Address() string
	Name() string
}

// Message is a piece of communication usually being passed between a sender and and a reciever
type Message interface {
	Subject() string
	Body() string
	Href() string
	Template() string
}

// HTTPTransporter allows for configured http transport
type HTTPTransporter func(req *http.Request) (*http.Response, error)

// Type is a communication type
type Type string

func (t Type) String() string {
	return string(t)
}

const (
	// Email communication type
	Email Type = "email"
	// Twitter communication type
	Twitter Type = "twitter"
)

//Communicator is something that does the communication
type Communicator interface {
	Communicate(Sender, Reciever, Message) error
	CommunicationType() string
}

//Service is the communication service
type Service struct {
	SaveUpdater   domain.CommunicationSaver
	Communicators []Communicator
	Config        *config.Config
}

type CommunicationService interface {
	Send(conf *config.Config, t Type, s Sender, r Reciever, m Message) error
}

func NewService(conf *config.Config, saveUpdate domain.CommunicationSaver) Service {
	s := Service{
		SaveUpdater: saveUpdate,
	}
	s.Communicators = append(s.Communicators, email{
		Config: conf,
	})
	return s
}

// Send will send an a communication and return a Communication object
func (cs Service) Send(t Type, s Sender, r Reciever, m Message) error {
	var communicator Communicator

	for _, c := range cs.Communicators {
		if c.CommunicationType() == t.String() {
			communicator = c
			break
		}
	}
	if nil == communicator {
		return errors.New("unknown communication type " + string(t))
	}
	if err := communicator.Communicate(s, r, m); err != nil {
		return err
	}
	return nil
}
