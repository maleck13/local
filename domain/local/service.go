package local

import (
	"fmt"
	"net/http"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	e "github.com/maleck13/local/errors"
	"github.com/maleck13/local/external"
)

// Service handles the business logic around locals
type Service struct {
	Config   *config.Config
	UserRepo domain.UserFinderDeleterSaver
}

//Registerer defines methods for registration Service is an implementor
type Registerer interface {
	Register(*app.User) (*domain.User, error)
}

// NewService returns a configured local service
func NewService(config *config.Config, repo domain.UserRepo) *Service {
	return &Service{
		Config:   config,
		UserRepo: repo,
	}
}

// Register a user with locals
func (ls Service) Register(user *app.User) (*domain.User, error) {
	exist, err := ls.UserRepo.FindOneByFieldAndValue("Email", user.Email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, e.NewServiceError("user already exists ", http.StatusConflict)
	}
	u := domain.NewUser(user)
	if err := ls.UserRepo.SaveUpdate(u); err != nil {
		return nil, e.NewServiceError("failed to register user "+err.Error(), http.StatusInternalServerError)
	}
	return u, nil
}

// AddProviderData data takes a user and calls out to a provider such as (google) and adds additonal data to that user
func (ls Service) AddProviderData(user *app.User) (*app.User, error) {
	if user.SignupType == "google" {
		gapi := external.NewGoogleAPI(config.Conf)
		user, err := gapi.AddGoogleData(user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return user, nil
}

// Update updates the fields of a users profile
func (ls Service) Update(update *app.UpdateUser) (*domain.User, error) {
	existing, err := ls.UserRepo.FindOneByFieldAndValue("id", update.ID)
	if err != nil {
		return nil, err
	}
	existing.Area = update.Area
	existing.Email = update.Email
	existing.FirstName = update.FirstName
	existing.SecondName = update.SecondName
	fmt.Println(existing.User)
	if err := ls.UserRepo.SaveUpdate(existing); err != nil {
		return nil, err
	}
	return existing, nil
}
