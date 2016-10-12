package domain_test

import (
	"net/http"
	"testing"

	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	pt "github.com/maleck13/local/domain/testing"
	"github.com/maleck13/local/external"
	"github.com/maleck13/local/test"
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
			Name: "test authenticate fails for email different",
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
		mUser := pt.MakeTestUser("John", "Smith", "test@test.com", "", "local")
		authService := domain.AuthenticationService{Config: config.Conf, UserFinder: pt.NewUserFinder(mUser, nil, nil), Provider: "google", GoogleAPI: gAPI}
		token, err := authService.Authenticate("token", "test@test.com")
		if tv.ExpectError && err == nil {
			t.Fatal("expected an error but gone none")
		}
		if !tv.ExpectError && err != nil {
			t.Fatal("did not expect an error but gone one ", err.Error())
		}
		if !tv.ExpectError && token == "" {
			t.Fatal("expected a token ")
		}
	}

}

func TestAuthenticateViaJWTToken(t *testing.T) {

}

func TestAuthenticateViaLocalDB(t *testing.T) {

}
