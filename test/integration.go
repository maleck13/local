package test

import (
	"flag"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
)

const (
	DEFAULT_USER_AREA   = "timbooktwo"
	DEFAULT_USER_LAT    = 52.3
	DEFAULT_USER_LON    = -7.32
	DEFAULT_SIGNUP_TYPE = "google"
)

var (
	//IntegrationEnabled is flag for running the integration tests
	IntegrationEnabled = flag.Bool("integration", false, "enabled integration tests")
)

type user struct {
	Config   *config.Config
	userRepo domain.UserRepo
}

func NewUser(conf *config.Config) user {
	return user{
		Config:   conf,
		userRepo: domain.NewUserRepo(conf, domain.AdminActor{}, domain.Access{}),
	}
}

//SetupUser sets up a user with a set location and area
func (u user) Setup(email, firstName, secondName string) (*domain.User, error) {
	var lon = new(float64)
	var lat = new(float64)
	*lon = DEFAULT_USER_LON
	*lat = DEFAULT_USER_LAT
	user := &domain.User{
		User: &app.User{
			Area:       DEFAULT_USER_AREA,
			Email:      email,
			FirstName:  firstName,
			SecondName: secondName,
			Location: &app.Location{
				Lon: lon,
				Lat: lat,
			},
			SignupType: DEFAULT_SIGNUP_TYPE,
		},
	}
	if err := u.userRepo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u user) Delete(email string) error {
	return u.userRepo.DeleteByFieldAndValue("Email", email)
}
