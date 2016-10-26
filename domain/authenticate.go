package domain

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/external"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//Authenticator defines a user authenticator interface Service is an implementor
type Authenticator interface {
	Authenticate(token, id string) (string, error)
	// CreateToken takes the unique userId, the email of the user and the type of the user and returns a jwtToken
	CreateToken(id, email, uType string) (string, error)

	ResetPassword(userID, newPass string) error
	VerifyVerificationToken(token, uid, scope string) error
}

type AuthenticationService struct {
	Config     *config.Config
	UserFinder UserFinderSaver
	Provider   string
	GoogleAPI  *external.GoogleAPI
}

// Authenticate facade around the different authentication types
func (as AuthenticationService) Authenticate(token, id string) error {
	var err error
	switch as.Provider {
	case "google":
		err = as.googleAuthenticate(token, id)
	case "local":
		err = as.localAuthenticate(token, id)
	default:
		return goa.ErrUnauthorized("unknown provider")
	}
	if err != nil {
		return err
	}
	return nil
}

func (as AuthenticationService) googleAuthenticate(token, id string) error {
	clientID := as.Config.Google.ClientID
	info, err := as.GoogleAPI.RetrieveUserInfo(token)
	if err != nil {
		return goa.ErrUnauthorized("failed to retrive token")
	}
	if info.EmailVerified != "true" {
		return goa.ErrUnauthorized("failed validate email from token")
	}
	if info.Aud != clientID {
		return goa.ErrUnauthorized("failed  validate aud from token")
	}
	if info.Email != id {
		return goa.ErrUnauthorized("failed to validate email matched")
	}
	return nil
}

func (as AuthenticationService) localAuthenticate(token, email string) error {
	u, err := as.UserFinder.FindOneByFieldAndValue("Email", email)
	if err != nil {
		return err
	}
	if u == nil {
		return goa.ErrNotFound("no such user")
	}
	if !u.Active {
		return goa.ErrUnauthorized("account not activated.")
	}
	if "" == token {
		return goa.ErrBadRequest("password cannot be empty")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Token), []byte(token)); err != nil {
		return goa.ErrUnauthorized("not authorised")
	}
	return nil
}

//CreateToken creates a new JWToken
func (as AuthenticationService) CreateToken(id, email, uType string, claimModifier func(jwt.MapClaims) jwt.Claims) (string, error) {
	in2 := time.Now().Add(time.Hour * 2).Unix()
	var scopes = []string{"api:access"}
	if uType == "admin" {
		scopes = append(scopes, "admin:access")
	}

	t := jwt.New(jwt.SigningMethodHS512)
	claims := jwt.MapClaims{
		"id":     id,
		"iss":    "Locals",              // who creates the token and signs it
		"aud":    "Locals",              // to whom the token is intended to be sent
		"exp":    in2,                   // time when the token will expire (10 minutes from now)
		"jti":    uuid.NewV4().String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),     // when the token was issued/created (now)
		"nbf":    2,                     // time before which the token is not yet valid (2 minutes ago)
		"sub":    email,                 // the subject/principal is whom the token is about
		"scopes": scopes,                // token scope - not a standard claim
		"type":   uType,                 //can be local, councillor or admin
	}
	t.Claims = claims
	if nil != claimModifier {
		t.Claims = claimModifier(claims)
	}
	return t.SignedString([]byte(as.Config.Jwt.Secret))
}

func (as AuthenticationService) parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: " + t.Header["alg"].(string))
		}
		return []byte(as.Config.Jwt.Secret), nil
	})
}

//VerifyVerificationToken will verify the jwt token is from the user uid and has the required scope
func (as AuthenticationService) VerifyVerificationToken(token, uid, scope string) error {
	jToken, err := as.parseToken(token)
	if err != nil {
		return errors.Wrap(err, "failed to parse token")
	}
	claims := jToken.Claims.(jwt.MapClaims)
	if val, ok := claims["scopes"]; ok {
		if val.(string) != scope {
			return goa.ErrUnauthorized("invalid scope")
		}
	} else {
		return goa.ErrUnauthorized("invalid no key")
	}
	if val, ok := claims["id"]; ok {
		if val.(string) != uid {
			return goa.ErrUnauthorized("invalid mismath uid")
		}
	} else {
		return goa.ErrUnauthorized("invalid  no uid")
	}
	if val, ok := claims["exp"]; ok {
		v := val.(float64)
		if int64(v) < time.Now().Unix() {
			return goa.ErrUnauthorized("invalid  expired")
		}
	} else {
		return goa.ErrUnauthorized("invalid")
	}
	return nil
}

//ResetPassword will reset a password for a user using the passed password and bcrypt hashing.
func (as AuthenticationService) ResetPassword(userID, newPass string) error {
	u, err := as.UserFinder.FindOneByFieldAndValue("id", userID)
	if err != nil {
		return err
	}
	if u == nil {
		return errors.New("no user by that id")
	}
	fmt.Println("ResetPassword ", u.SignupType)
	if u.SignupType == "google" {
		return errors.New("cannot reset password for google signup account")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}
	u.Token = string(hash)
	return as.UserFinder.SaveUpdate(u)
}

// NewAuthenticateService returns an AuthenticationService that is an Authenticator
func NewAuthenticateService(provider string, conf *config.Config, userFinder UserFinderSaver) AuthenticationService {
	return AuthenticationService{
		Config:     conf,
		UserFinder: userFinder,
		Provider:   provider,
		GoogleAPI:  external.NewGoogleAPI(conf),
	}

}
