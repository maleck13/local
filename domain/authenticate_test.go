package domain_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	pt "github.com/maleck13/local/domain/testing"
	"github.com/maleck13/local/external"
	"github.com/maleck13/local/test"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	test.SetUpConfig()
}

type MockGoogleInfoRetriever struct {
	requester func() (*http.Response, error)
}

func (mr MockGoogleInfoRetriever) Get(url string) (*http.Response, error) {
	return mr.requester()
}

func TestAuthenticateViaGoogle(t *testing.T) {
	//need a configured google client

	var tests = []struct {
		Name         string
		Response     string
		ResponseCode int
		ExpectError  bool
	}{
		{
			Name: "test authenticates ok",
			Response: `{
                "email_verified":"true",
                "aud":"` + config.Conf.Google.ClientID + `",
                "email":"test@test.com"
            }`,
			ResponseCode: 200,
			ExpectError:  false,
		},
		{
			Name: "test authenticate fails for email_verified",
			Response: `{
                "email_verified":"false",
                "aud":"` + config.Conf.Google.ClientID + `",
                "email":"test@test.com"
            }`,
			ResponseCode: 200,
			ExpectError:  true,
		},
		{
			Name: "test authenticate fails for wrong email",
			Response: `{
                "email_verified":"true",
                "aud":"` + config.Conf.Google.ClientID + `",
                "email":"different@test.com"
            }`,
			ResponseCode: 200,
			ExpectError:  true,
		},
	}

	for _, tv := range tests {
		requester := test.CreateRequest(t, tv.ResponseCode, "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=token", tv.Response)
		gAPI := &external.GoogleAPI{
			Config:             config.Conf,
			TokenInfoRetriever: MockGoogleInfoRetriever{requester: requester},
		}
		mUser := pt.MakeTestUser("John", "Smith", "test@test.com", "", "local", "")
		authService := domain.AuthenticationService{Config: config.Conf, UserFinder: pt.NewUserFinderSaver(mUser, nil, nil), Provider: "google", GoogleAPI: gAPI}
		err := authService.Authenticate("token", "test@test.com")
		if tv.ExpectError && err == nil {
			t.Fatal("expected an error but gone none")
		}
		if !tv.ExpectError && err != nil {
			t.Fatal("did not expect an error but gone one ", err.Error())
		}
	}

}

func TestAuthenticateLocal(t *testing.T) {

	tests := []struct {
		Name        string
		Password    string
		UserName    string
		IsActivated bool
		ExpectError bool
	}{
		{
			Name:        "test local authenticate works",
			Password:    "Password1",
			UserName:    "test@test.com",
			IsActivated: true,
			ExpectError: false,
		},
		{
			Name:        "test local authenticate fails with bad pass",
			Password:    "Password12",
			UserName:    "test@test.com",
			IsActivated: true,
			ExpectError: true,
		},
		{
			Name:        "test local authenticate fails with bad username",
			Password:    "Password12",
			UserName:    "test2@test.com",
			IsActivated: true,
			ExpectError: true,
		},
		{
			Name:        "test local authenticate fails when user not activated",
			Password:    "Password1",
			UserName:    "test@test.com",
			IsActivated: false,
			ExpectError: true,
		},
	}
	for _, tv := range tests {
		t.Run(tv.Name, func(t *testing.T) {
			user := pt.MakeTestUser("John", "Smith", "test@test.com", "somewhere", "local", "")
			user.Active = tv.IsActivated
			encPass, err := bcrypt.GenerateFromPassword([]byte("Password1"), bcrypt.DefaultCost)
			if err != nil {
				t.Fatal("faled to generate hashed password", err.Error())
			}
			user.Token = string(encPass)
			uf := pt.NewUserFinderSaver(user, nil, nil)
			authenticator := domain.AuthenticationService{Config: config.Conf, UserFinder: uf, Provider: "local"}
			err = authenticator.Authenticate(tv.Password, tv.UserName)
			if tv.ExpectError && nil == err {
				t.Fatal("expected an error but got none")
			}
			if !tv.ExpectError && err != nil {
				t.Fatal("did not expect an error but got one ", err.Error())
			}
		})
	}
}

func TestResetPassword(t *testing.T) {
	tests := []struct {
		Name        string
		UserID      string
		NewPass     string
		SignUpType  string
		ExpectError bool
		Error       error
		Assert      func(t *testing.T) func(u *domain.User)
	}{
		{
			Name:        "test reset password resets the users password",
			UserID:      "testid",
			NewPass:     "mypassword",
			SignUpType:  "local",
			ExpectError: false,
			Error:       nil,
			Assert: func(t *testing.T) func(u *domain.User) {
				return func(u *domain.User) {
					if err := bcrypt.CompareHashAndPassword([]byte(u.Token), []byte("mypassword")); err != nil {
						t.Fatalf("did not expect an error %s ", err.Error())
					}
				}
			},
		},
		{
			Name:        "test reset password fails on error",
			UserID:      "testid",
			NewPass:     "mypassword",
			SignUpType:  "local",
			ExpectError: true,
			Error:       fmt.Errorf("an error "),
			Assert: func(t *testing.T) func(*domain.User) {
				return func(u *domain.User) {
					t.Fatal("should not have got here")
				}
			},
		},
		{
			Name:        "test reset password fails on google signup",
			UserID:      "testid",
			NewPass:     "mypassword",
			SignUpType:  "google",
			ExpectError: true,
			Error:       nil,
			Assert: func(t *testing.T) func(*domain.User) {
				return func(u *domain.User) {
					t.Fatal("should not have got here")
				}
			},
		},
	}

	for _, tv := range tests {
		t.Run(tv.Name, func(t *testing.T) {
			user := pt.MakeTestUser("John", "Smith", "test@test.com", "somewhere", tv.SignUpType, tv.UserID)
			uf := pt.NewUserFinderSaver(user, nil, tv.Error)
			uf.SaveUpdateAssert = tv.Assert(t)
			authenticator := domain.AuthenticationService{Config: config.Conf, UserFinder: uf, Provider: "local"}
			err := authenticator.ResetPassword(tv.UserID, tv.NewPass)
			if err != nil && !tv.ExpectError {
				t.Fatal("did not expect an error changing password but got " + err.Error())
			}
			if err == nil && tv.ExpectError {
				t.Fatal("expected an error bug got nil")
			}
		})
	}
}

func TestVerifyVerificationToken(t *testing.T) {

}
