package testing

import (
	"net/http"

	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
)

func AddJwtToken(conf *config.Config, tokenType string, req *http.Request) error {
	authenticate := domain.AuthenticationService{
		Config: conf,
	}
	token, err := authenticate.CreateToken("", "test@test.com", tokenType)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	return nil
}
