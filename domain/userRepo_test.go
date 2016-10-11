package domain_test

import (
	"fmt"
	"testing"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/test"
)

const (
	existingLocalUser = "mylocaluser@test.com"
)

func init() {
	if !*test.IntegrationEnabled {
		fmt.Println("integration is disabled")
	}
}

func makeTestUser(fn, sn, email, area, uType string) *domain.User {
	appUser := &app.User{
		FirstName:  fn,
		SecondName: sn,
		Email:      email,
		Area:       area,
		Type:       uType,
	}
	return domain.NewUserFromRequest(appUser)
}

func setUpUserRepoTest(t *testing.T) func() {
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	user := makeTestUser("John", "Smith", existingLocalUser, "some area", "local")
	if err := userRepo.SaveUpdate(user); err != nil {
		t.Fatal("failed to setUpUserRepoTest", err.Error())
	}
	return func() {
		if err := userRepo.DeleteByFieldAndValue("FirstName", "John"); err != nil {
			t.Log("failed to tear down users", err.Error())
		}
	}
}

func TestFindOneByFieldAndValue(t *testing.T) {
	tearDown := setUpUserRepoTest(t)
	defer tearDown()
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	var tests = []struct {
		Name        string
		Field       string
		Value       string
		ExpectError bool
		Assert      func(u *domain.User) error
	}{
		{
			Name:        "test find existing user ok",
			Field:       "Email",
			Value:       existingLocalUser,
			ExpectError: false,
			Assert: func(u *domain.User) error {
				if u.Email != existingLocalUser {
					return fmt.Errorf("expected email to match %s != %s ", u.Email, existingLocalUser)
				}
				return nil
			},
		},
		{
			Name:        "test find user that doesn't exist",
			Field:       "Email",
			Value:       "idont@exist.com",
			ExpectError: false,
			Assert: func(u *domain.User) error {
				if u != nil {
					return fmt.Errorf("expected not to find a user")
				}
				return nil
			},
		},
	}
	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			u, err := userRepo.FindOneByFieldAndValue(v.Field, v.Value)
			if v.ExpectError && nil == err {
				t.Fatal("expected an error but got none")
			} else if !v.ExpectError && err != nil {
				t.Fatalf("did not expect an error but got %s", err.Error())
			}
			if err := v.Assert(u); err != nil {
				t.Fatal("assert error ", err.Error())
			}
		})
	}
}

func TestFindAllByTypeAndArea(t *testing.T) {
	tearDown := setUpUserRepoTest(t)
	defer tearDown()
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	var tests = []struct {
		Name        string
		Value       string
		ExpectError bool
		Assert      func(us []*domain.User) error
	}{
		{
			Name:        "test find all works by area",
			Value:       "some area",
			ExpectError: false,
			Assert: func(us []*domain.User) error {
				if len(us) != 1 {
					return fmt.Errorf("expected on one user got %d", len(us))
				}
				u := us[0]
				if u.Area != "some area" {
					return fmt.Errorf("expected area to match")
				}
				return nil
			},
		},
		{
			Name:        "test find all return nil when non found",
			Value:       "not an area",
			ExpectError: false,
			Assert: func(us []*domain.User) error {
				if nil != us {
					return fmt.Errorf("expected no result values")
				}
				return nil
			},
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			us, err := userRepo.FindAllByTypeAndArea("local", v.Value)
			if v.ExpectError && err == nil {
				t.Fatal("expected an error but got none")
			} else if !v.ExpectError && err != nil {
				t.Fatal(" didnt expect an error but got one ", err.Error())
			}
			if err := v.Assert(us); err != nil {
				t.Fatal("assert failed ", err.Error())
			}
		})
	}

}

func TestSaveUpdate(t *testing.T) {
	tearDown := setUpUserRepoTest(t)
	defer tearDown()
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	var tests = []struct {
		Name        string
		Value       string
		ExpectError bool
		Assert      func(us []*domain.User) error
	}{}
}

func TestDeleteByFieldAndValue(t *testing.T) {

}
