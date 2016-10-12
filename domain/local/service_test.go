package local_test

import (
	"fmt"
	"testing"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/domain/local"
	pt "github.com/maleck13/local/domain/testing"
	"github.com/maleck13/local/test"
)

func init() {
	if *test.IntegrationEnabled {
		fmt.Println("integration tests enabled")
	}
}

func TestRegister(t *testing.T) {

	var tests = []struct {
		Name           string
		ExpectError    bool
		ShouldFindUser bool
		Assert         func(*domain.User) error
	}{
		{
			Name:           "test register succeeds",
			ExpectError:    false,
			ShouldFindUser: false,
			Assert: func(u *domain.User) error {
				if u.ID == "" {
					return fmt.Errorf("expected an id to be set on user")
				}
				if u.Email != "test@test.com" {
					return fmt.Errorf("expected users email to match")
				}
				return nil
			},
		},
		{
			Name:           "test register wont register existing user",
			ExpectError:    true,
			ShouldFindUser: true,
			Assert: func(u *domain.User) error {
				return nil
			},
		},
	}

	for _, tv := range tests {
		t.Run(tv.Name, func(t *testing.T) {
			var userRepo = pt.NewUserFinderSaver(nil, nil, nil)
			user := pt.MakeTestUser("John", "Smith", "test@test.com", "somewhere", "local", "")
			if tv.ShouldFindUser {
				userRepo = pt.NewUserFinderSaver(user, nil, nil)
			}
			service := local.NewService(config.Conf, userRepo)
			registered, err := service.Register(user.User)
			if tv.ExpectError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tv.ExpectError && err != nil {
				t.Fatal("did not expect an error but got ", err.Error())
			}
			if err := tv.Assert(registered); err != nil {
				t.Fatal("assert failed ", err.Error())
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	var tests = []struct {
		Name           string
		ExpectError    bool
		ShouldFindUser bool
		Assert         func(du *domain.User, uu *app.UpdateUser) error
		Update         *app.UpdateUser
		User           *domain.User
	}{
		{
			Name:           "test update works ok",
			ExpectError:    false,
			ShouldFindUser: true,
			Update:         pt.MakeTestUpdateUser("id", "somearea", "test@test.com", "John", "Blithe"),
			User:           pt.MakeTestUser("John", "Smith", "test@test.com", "somewhere", "local", "id"),
			Assert: func(du *domain.User, uu *app.UpdateUser) error {
				if du.ID != uu.ID {
					return fmt.Errorf("expected the id to be the same %s != %s ", du.ID, uu.ID)
				}
				if du.Area != uu.Area {
					return fmt.Errorf("expected the areas to be the same %s != %s ", du.Area, uu.Area)
				}
				return nil
			},
		},
		{
			Name:           "test update fails if no user",
			ExpectError:    true,
			ShouldFindUser: false,
			Update:         pt.MakeTestUpdateUser("id", "somearea", "test@test.com", "John", "Blithe"),
			User:           nil,
			Assert: func(du *domain.User, uu *app.UpdateUser) error {
				if nil != du && nil != uu {
					return fmt.Errorf("did not expect to get an update or existing user")
				}
				return nil
			},
		},
	}

	for _, tv := range tests {
		t.Run(tv.Name, func(t *testing.T) {
			var userRepo = pt.NewUserFinderSaver(nil, nil, nil)
			if tv.ShouldFindUser {
				userRepo = pt.NewUserFinderSaver(tv.User, nil, nil)
			}
			service := local.NewService(config.Conf, userRepo)
			updated, err := service.Update(tv.Update)
			if tv.ExpectError && err == nil {
				t.Fatal("expected error but got none")
			}
			if !tv.ExpectError && err != nil {
				t.Fatal("did not expect an error but got ", err.Error())
			}
			if err := tv.Assert(updated, tv.Update); err != nil {
				t.Fatal("assert failed ", err.Error())
			}
		})
	}
}
