package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	UserRepo           UserRepo
}

type googleTokenInfo struct {
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
func NewGoogleAPI(config *config.Config, userRepo UserRepo) *GoogleAPI {
	return &GoogleAPI{
		Config:             config,
		TokenInfoRetriever: http.DefaultClient,
		UserRepo:           userRepo,
	}
}

//Authenticate takes a token retrieved from google and validates it with the google api
func (ga *GoogleAPI) Authenticate(token, email string) (*User, error) {
	info, err := ga.retrieve(token)
	if err != nil {
		return nil, err
	}
	clientID := ga.Config.Google.ClientID

	if info.EmailVerified != "true" {
		return nil, e.NewServiceError("failed to authenticate with google", http.StatusUnauthorized)
	}
	if info.Aud != clientID {
		return nil, e.NewServiceError("failed to authenticate with google. Client id mismatch", http.StatusUnauthorized)
	}
	if info.Email != email {
		return nil, e.NewServiceError("failed to authenticate with google. user mismatch", http.StatusUnauthorized)
	}
	user, err := ga.UserRepo.FindOneByFieldAndValue("Email", email)
	if err != nil {
		return nil, err
	}
	if nil == user {
		return nil, e.NewServiceError("failed to find user ", http.StatusNotFound)
	}

	return user, nil
}

func (ga *GoogleAPI) retrieve(token string) (*googleTokenInfo, error) {
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

	var gTokenInfo = &googleTokenInfo{}
	json.Unmarshal(body, gTokenInfo)
	return gTokenInfo, nil
}

//FillUserDetailsFromGoogleToken takes a user and adds the info from the google token response such as email and FirstName
func (ga *GoogleAPI) fillUserDetailsFromGoogleToken(user *app.User, info *googleTokenInfo) (*User, error) {
	u := NewUserFromRequest(user)
	u.Email = info.Email
	u.FirstName = info.GivenName
	u.SecondName = info.FamilyName
	u.User.Type = "local"
	i, err := strconv.ParseInt(info.Exp, 10, 64)
	if err != nil {
		return nil, e.NewServiceError("failed to parse exp time from google ", http.StatusInternalServerError)
	}

	u.LoginExpires = int(i)
	return u, nil
}

//Register takes a user retrieves google info about that user and saves it to the database
func (ga *GoogleAPI) Register(userSp *app.User) (*User, error) {
	tokenInfo, err := ga.retrieve(userSp.Token)
	if err != nil {
		return nil, err
	}
	user, err := ga.fillUserDetailsFromGoogleToken(userSp, tokenInfo)
	if err != nil {
		return nil, err
	}

	exist, err := ga.UserRepo.FindOneByFieldAndValue("Email", user.Email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, e.NewServiceError("user already exists ", http.StatusConflict)
	}
	if err := ga.UserRepo.Save(user); err != nil {
		return nil, e.NewServiceError("failed to register user "+err.Error(), http.StatusInternalServerError)
	}
	return user, nil
}
