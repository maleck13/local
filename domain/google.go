package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	e "github.com/maleck13/local/errors"
	"github.com/Sirupsen/logrus"
)

type TokenInfoRetriever interface {
	Retrieve(token string) (*GoogleTokenInfo, error)
}

type GoogleApi struct {
	Config *config.Config
}

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

func NewGoogleApi(config *config.Config) *GoogleApi {
	return &GoogleApi{
		Config: config,
	}
}

func (ga *GoogleApi) Authenticate(token, email string) error {
	info, err := ga.Retrieve(token)
	if err != nil {
		return err
	}
	clientId := ga.Config.Google.ClientID

	if info.EmailVerified != "true" {
		return e.NewServiceError("failed to authenticate with google", http.StatusUnauthorized)
	}
	if info.Aud != clientId {
		return e.NewServiceError("failed to authenticate with google. Client id mismatch", http.StatusUnauthorized)
	}
	if info.Email != email {
		return e.NewServiceError("failed to authenticate with google. user mismatch", http.StatusUnauthorized)
	}
	userRepo := &data.UserRepo{}
	user, err := userRepo.FindOneByFieldAndValue("Email", email)
	if err != nil {
		return err
	}
	if nil == user {
		return e.NewServiceError("failed to find user ", http.StatusNotFound)
	}

	return nil
}

func (ga *GoogleApi) Retrieve(token string) (*GoogleTokenInfo, error) {
	url := fmt.Sprintf(ga.Config.Google.ValidatorURL, token)
	res, err := http.Get(url)
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

func (ga *GoogleApi) decorateUserFromGoogleToken(user *app.User, info *GoogleTokenInfo) (*data.User, error) {
	logrus.Info("google token ", info)
	u := data.NewUserFromRequest(user)
	u.Email = info.Email
	u.FirstName = info.GivenName
	u.SecondName = info.FamilyName
	u.Type = "local"
	i, err := strconv.ParseInt(info.Exp, 10, 64)
	if err != nil {
		return nil, e.NewServiceError("failed to parse exp time from google ", http.StatusInternalServerError)
	}

	u.LoginExpires = int(i)
	return u, nil
}

func (ga *GoogleApi) Register(userSp *app.User) (*data.User, error) {
	tokenInfo, err := ga.Retrieve(userSp.Token)
	if err != nil {
		return nil, err
	}
	user, err := ga.decorateUserFromGoogleToken(userSp, tokenInfo)
	if err != nil {
		return nil, err
	}
	userRepo := data.UserRepo{}
	exist, err := userRepo.FindOneByFieldAndValue("Email", user.Email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, e.NewServiceError("user already exists ", http.StatusConflict)
	}
	if err := userRepo.Save(user); err != nil {
		return nil, e.NewServiceError("failed to register user "+err.Error(), http.StatusInternalServerError)
	}
	return user, nil
}
