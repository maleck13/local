package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maleck13/local/config"
	"github.com/maleck13/local/test"
)

var (
	testServer    *httptest.Server
	existingEmail = "testcreate@test.com"
)

func setUp(t *testing.T) func(emails []string) {
	conf, err := config.LoadConfig("./config/config.json")
	if err != nil {
		t.Fail()
	}
	mux := buildService(conf)
	testServer = httptest.NewServer(mux.Mux)
	t.Log("setup")

	userSetup := test.NewUser(conf)
	testCreateEmail := "testcreate@test.com"
	return func(emails []string) {
		emails = append(emails, testCreateEmail)
		testServer.Close()
		for _, e := range emails {
			userSetup.Delete(e)
			if _, err := userSetup.Setup(testCreateEmail, "testcreate", "smith"); err != nil {
				t.Fail()
			}
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

func TestCreateUser(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skipf("integration not enabled")
	}
	var()
	tearDown := setUp(t)
	defer tearDown()

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
		res, err := http.Post(testServer.URL+"/user/signup", "application/json", userJSON(, "john", "smith"))
		if err != nil {
			t.Fatalf(err.Error())
		}
		defer res.Body.Close()
		if res.StatusCode != 409 {
			t.Fatal("expected a conflict")
		}
	})
}
