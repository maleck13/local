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
func (u user) Setup(email, firstName, secondName, uType string) (*domain.User, error) {
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
			Type:       uType,
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
	buildUserController(mux)
	testServer = httptest.NewServer(mux.Mux)
	userSetup := newUser(conf)
	_, err = userSetup.Setup(existingEmail, "john", "smith", "councillor")
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

func TestCouncillorSignup(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("Integration disabled")
	}
	//setup a server
	tearDown := setUp(t)
	defer tearDown([]string{existingEmail, testCreateEmail})
	tests := []struct {
		Name       string
		Endpoint   string
		StatusCode int
	}{
		{
			Name:       "test councillor exists with given mail",
			Endpoint:   "/user/councillor/signup",
			StatusCode: 200,
		},
	}
	for _, tr := range tests {
		t.Run(tr.Name, func(t *testing.T) {
			endpoint := testServer.URL + tr.Endpoint
			req, err := http.NewRequest("POST", endpoint, strings.NewReader(`{"email":"`+existingEmail+`"}`))
			if err != nil {
				t.Fatal("failed to create request", err.Error())
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal("failed to check councillors email", err.Error())
			}
			defer res.Body.Close()
			if tr.StatusCode != res.StatusCode {
				t.Fatalf("expected status %d but got %d", tr.StatusCode, res.StatusCode)
			}
		})
	}

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
			t.Fatalf("expected a 400 error code but got %d ", res.StatusCode)
		}
	})
	t.Run("should not add user that exists", func(t *testing.T) {
		res, err := http.Post(testServer.URL+"/user/signup", "application/json", userJSON(existingEmail, "john", "smith"))
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 409 {
			t.Fatalf("expected a conflict but got %d", res.StatusCode)
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

func TestLoginUser(t *testing.T) {

}
