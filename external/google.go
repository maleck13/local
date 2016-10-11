package external

//interfaces with the external google api in order to authenticate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	e "github.com/maleck13/local/errors"
)

//TokenInfoRetriever retrieves token  info from Google
type TokenInfoRetriever interface {
	Get(url string) (*http.Response, error)
}

//GoogleAPI provides methods for interacting with the google api
type GoogleAPI struct {
	Config             *config.Config
	TokenInfoRetriever TokenInfoRetriever
}

// GoogleTokenInfo encapsulates the response for reading a oAuth Token from Google
type GoogleTokenInfo struct {
	Alg           string `json:"alg"`
	AtHash        string `json:"at_hash"`
	Aud           string `json:"aud"`
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Exp           string `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           string `json:"iat"`
	Iss           string `json:"iss"`
	Kid           string `json:"kid"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}

//NewGoogleAPI interact with google
func NewGoogleAPI(config *config.Config) *GoogleAPI {
	return &GoogleAPI{
		Config:             config,
		TokenInfoRetriever: http.DefaultClient,
	}
}

//Authenticate takes a token retrieved from google and validates it with the google api
func (ga *GoogleAPI) Authenticate(token, email string) error {
	info, err := ga.RetrieveUserInfo(token)
	if err != nil {
		return goa.ErrInternal("failed to retrive token")
	}
	clientID := ga.Config.Google.ClientID

	if info.EmailVerified != "true" {
		return goa.ErrUnauthorized("failed to retrive token")
	}
	if info.Aud != clientID {
		return e.NewServiceError("failed to authenticate with google. Client id mismatch", http.StatusUnauthorized)
	}
	if info.Email != email {
		return e.NewServiceError("failed to authenticate with google. user mismatch", http.StatusUnauthorized)
	}
	return nil
}

func (ga *GoogleAPI) RetrieveUserInfo(token string) (*GoogleTokenInfo, error) {
	url := fmt.Sprintf(ga.Config.Google.ValidatorURL, token)
	res, err := ga.TokenInfoRetriever.Get(url)
	if err != nil {
		return nil, e.NewServiceError("failed to validate token with google", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, e.NewServiceError(err.Error(), http.StatusInternalServerError)
	}

	var gTokenInfo = &GoogleTokenInfo{}
	json.Unmarshal(body, gTokenInfo)
	return gTokenInfo, nil
}

//AddGoogleData takes a user retrieves google info about that user and adds that data to to the user model
func (ga *GoogleAPI) AddGoogleData(u *app.User) (*app.User, error) {
	info, err := ga.RetrieveUserInfo(u.Token)
	if err != nil {
		return nil, err
	}
	u.Email = info.Email
	u.FirstName = info.GivenName
	u.SecondName = info.FamilyName
	u.Type = "local"
	return u, nil
}
