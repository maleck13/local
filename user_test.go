package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/test"
)

var (
	testServer      *httptest.Server
	existingEmail   = "testcreate@test.com"
	testCreateEmail = "testdontexist@test.com"
)

const (
	DEFAULT_USER_AREA   = "timbooktwo"
	DEFAULT_USER_LAT    = 52.3
	DEFAULT_USER_LON    = -7.32
	DEFAULT_SIGNUP_TYPE = "google"
)

type user struct {
	Config   *config.Config
	userRepo domain.UserRepo
}

func newUser(conf *config.Config) user {
	return user{
		Config:   conf,
		userRepo: domain.NewUserRepo(conf, domain.NewAdminActor(), domain.AuthorisationService{}),
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
	if err := u.userRepo.SaveUpdate(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u user) Delete(email string) error {
	return u.userRepo.DeleteByFieldAndValue("Email", email)
}

func setUp(t *testing.T) func(emails []string) {
	conf, err := config.LoadConfig("./config/config.json")
	if err != nil {
		t.Fail()
	}
	mux := buildService(conf)
	testServer = httptest.NewServer(mux.Mux)
	userSetup := newUser(conf)
	_, err = userSetup.Setup(existingEmail, "john", "smith")
	if err != nil {
		t.Error("setup failed", err.Error())
	}
	return func(emails []string) {
		testServer.Close()
		for _, e := range emails {
			userSetup.Delete(e)
		}
	}
}

func userJSON(email, firstName, secondName string) io.Reader {
	return strings.NewReader(`{
        "email":"` + email + `",
        "signupType":"local",
        "token":"testtoken",
        "firstName":"` + firstName + `",
        "secondName":"` + secondName + `"
    }`)
}

func badUserJSON() io.Reader {
	return strings.NewReader(`{
        "email":"someemail@test.com"
    }
    `)
}

func TestHttpCreateUser(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skipf("integration not enabled")
	}
	tearDown := setUp(t)
	defer tearDown([]string{existingEmail, testCreateEmail})

	t.Run("should be a bad request with missing params", func(t *testing.T) {
		res, err := http.Post(testServer.URL+"/user/signup", "application/json", badUserJSON())
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 400 {
			t.Fatal("expected a 400 error code")
		}
	})
	t.Run("should not add user that exists", func(t *testing.T) {
		res, err := http.Post(testServer.URL+"/user/signup", "application/json", userJSON(existingEmail, "john", "smith"))
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 409 {
			t.Fatal("expected a conflict")
		}
	})
	t.Run("should add user that doesnt exist", func(t *testing.T) {
		res, err := http.Post(testServer.URL+"/user/signup", "application/json", userJSON(testCreateEmail, "john", "smith"))
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 201 {
			t.Fatal("expected a created response but got ", res.StatusCode)
		}
	})
}
