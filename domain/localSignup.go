package domain

import (
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
)

//LocalSignup handles signup users not coming from external service
type LocalSignup struct {
	Config   *config.Config
	UserRepo UserRepo
}

//Register adds a new user to the site
func (LocalSignup) Register(*app.User) (*User, error) {
	return nil, nil
}

//NewLocalSignup interact with local db
func NewLocalSignup(config *config.Config, userRepo UserRepo) LocalSignup {
	return LocalSignup{
		Config:   config,
		UserRepo: userRepo,
	}
}
