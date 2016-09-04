package domain

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maleck13/local/config"
	"github.com/satori/go.uuid"
)

type JWT struct {
	Config *config.Config
}

type localClaim struct {
	Email  string
	scopes string
	jwt.StandardClaims
}

//CreateToken creates a new JWToken
func (j *JWT) CreateToken(user *User) (string, error) {
	in2 := time.Now().Add(time.Hour * 2).Unix()

	t := jwt.New(jwt.SigningMethodHS512)
	t.Claims = jwt.MapClaims{
		"id":     user.ID,
		"iss":    "Locals",              // who creates the token and signs it
		"aud":    "Locals",              // to whom the token is intended to be sent
		"exp":    in2,                   // time when the token will expire (10 minutes from now)
		"jti":    uuid.NewV4().String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),     // when the token was issued/created (now)
		"nbf":    2,                     // time before which the token is not yet valid (2 minutes ago)
		"sub":    user.Email,            // the subject/principal is whom the token is about
		"scopes": "api:access",          // token scope - not a standard claim
		"type":   user.User.Type,        //can be local, councillor or admin
	}
	return t.SignedString([]byte(j.Config.Jwt.Secret))
}

//Authenticate implements the Authenticator interface for jwt token
func (j *JWT) Authenticate(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return j.Config.Jwt.Secret, nil
	})
	if err != nil {
		return errValidationFailed("not authorised")
	}
	return nil
}
