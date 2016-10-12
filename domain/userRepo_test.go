package domain_test

import (
	"fmt"
	"testing"

	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	pt "github.com/maleck13/local/domain/testing"
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

func setUpUserRepoTest(t *testing.T) func() {
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	user := pt.MakeTestUser("John", "Smith", existingLocalUser, "some area", "local")
	if err := userRepo.SaveUpdate(user); err != nil {
		t.Log("failed to setUpUserRepoTest", err.Error())
	}
	return func() {
		if err := userRepo.DeleteByFieldAndValue("FirstName", "John"); err != nil {
			t.Log("failed to tear down users", err.Error())
		}
	}
}

func TestFindOneByFieldAndValue(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("integration disabled")
	}
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
	if !*test.IntegrationEnabled {
		t.Skip("integration disabled")
	}
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
	if !*test.IntegrationEnabled {
		t.Skip("integration disabled")
	}
	tearDown := setUpUserRepoTest(t)
	defer tearDown()
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	exists, err := userRepo.FindOneByFieldAndValue("Email", existingLocalUser)
	if err != nil {
		t.Fatal("error fetching existing user", err.Error())
	}
	//prepare our table tests
	var tests = []struct {
		Name        string
		User        *domain.User
		ExpectError bool
		Assert      func(u1 *domain.User) error
	}{
		{
			Name:        "test save new user",
			User:        pt.MakeTestUser("John", "Smith", "save@test.com", "somewhere", "local"),
			ExpectError: false,
			Assert: func(u1 *domain.User) error {
				u2, err := userRepo.FindOneByFieldAndValue("Email", u1.Email)
				if err != nil {
					return err
				}
				if u1.Email != u2.Email {
					return fmt.Errorf("expected emails to match %s != %s", u1.Email, u2.Email)
				}
				return nil
			},
		},
		{
			Name:        "test save existing user",
			User:        exists,
			ExpectError: false,
			Assert: func(u1 *domain.User) error {
				u2, err := userRepo.FindOneByFieldAndValue("Email", u1.Email)
				if err != nil {
					return err
				}
				if u1.Email != u2.Email {
					return fmt.Errorf("expected emails to match %s != %s", u1.Email, u2.Email)
				}
				if u1.ID != u2.ID {
					return fmt.Errorf("expected ids to match %s != %s", u1.ID, u2.ID)
				}
				return nil
			},
		},
		{
			Name:        "test fails creating a new existing user",
			User:        pt.MakeTestUser("John", "smith", existingLocalUser, "somewhere", "local"),
			ExpectError: true,
			Assert: func(u1 *domain.User) error {
				return nil
			},
		},
	}
	//run tests
	for _, tv := range tests {
		t.Run(tv.Name, func(t *testing.T) {
			err := userRepo.SaveUpdate(tv.User)
			if tv.ExpectError && err == nil {
				t.Fatal("expected an error but gone none")
			}
			if !tv.ExpectError && err != nil {
				t.Fatal("did not expect an error but got one: ", err.Error())
			}
			if err := tv.Assert(tv.User); err != nil {
				t.Fatal("asserts failed ", err.Error())
			}
		})
	}
}

func TestDeleteByFieldAndValue(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("integration disabled")
	}
	tearDown := setUpUserRepoTest(t)
	defer tearDown()
	userRepo := domain.NewUserRepo(config.Conf, domain.NewAdminActor(), domain.AuthorisationService{})
	exists, err := userRepo.FindOneByFieldAndValue("Email", existingLocalUser)
	if err != nil {
		t.Fatal("error fetching existing user", err.Error())
	}
	tests := []struct {
		Name        string
		Field       string
		Value       string
		ExpectError bool
		Assert      func(key, val string) error
	}{
		{
			Name:        "test deletes existing user",
			Field:       "Email",
			Value:       exists.Email,
			ExpectError: false,
			Assert: func(key, val string) error {
				exists, err := userRepo.FindOneByFieldAndValue(key, val)
				if err != nil {
					return err
				}
				if exists != nil {
					return fmt.Errorf("did not expect to find a user just deleted")
				}
				return nil
			},
		},
		{
			Name:        "test deletes non existing user",
			Field:       "Email",
			Value:       "idont@test.com",
			ExpectError: false,
			Assert: func(key, val string) error {
				exists, err := userRepo.FindOneByFieldAndValue(key, val)
				if err != nil {
					return err
				}
				if exists != nil {
					return fmt.Errorf("did not expect to find a user that doesn't exist ")
				}
				return nil
			},
		},
	}

	//run tests
	for _, tv := range tests {
		t.Run(tv.Name, func(t *testing.T) {
			err := userRepo.DeleteByFieldAndValue(tv.Field, tv.Value)
			if tv.ExpectError && err == nil {
				t.Fatal("expected an error but gone none")
			}
			if !tv.ExpectError && err != nil {
				t.Fatal("did not expect an error but got one: ", err.Error())
			}
			if err := tv.Assert(tv.Field, tv.Value); err != nil {
				t.Fatal("asserts failed", err.Error())
			}
		})
	}

}
