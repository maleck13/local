package domain_test

import (
	"fmt"
	"testing"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	"github.com/maleck13/local/test"
)

var (
	existingCouncillor *domain.Councillor
)

func makeCouncillor(fn, sn, email, address, area, party string) *domain.Councillor {
	appCouncillor := &app.Councillor{
		FirstName:  fn,
		SecondName: sn,
		Email:      email,
		Address:    address,
		Area:       area,
		Party:      party,
	}
	return domain.NewCouncillor(appCouncillor)
}

func councillorRepoTestSetup(t *testing.T) func() {
	councillorRepo := domain.NewCouncillorRepo(config.Conf, domain.NewAdminActor(), &domain.AuthorisationService{})
	var tearDown = func() {
		t.Log("cleaning up councillorRepoTest")
		if err := councillorRepo.DeleteAllByKeyValue("FirstName", "john"); err != nil {
			t.Log("failed to clean up !!!!!!! ")
		}
	}
	existingCouncillor = makeCouncillor("john", "smith", "exits@test.com", "21 somewhere", "some area", "right of center")
	if err := councillorRepo.SaveUpdate(existingCouncillor); err != nil {
		t.Fatal(err.Error())
	}
	return tearDown
}

func TestSaveUpdateCouncillor(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("integration tests not enabled")
	}
	councillorRepo := domain.NewCouncillorRepo(config.Conf, domain.NewAdminActor(), &domain.AuthorisationService{})
	//setup returns a tearDown function
	tearDown := councillorRepoTestSetup(t)
	//tearDown
	defer tearDown()
	//define our tests
	tests := []struct {
		Name        string
		Councillor  *domain.Councillor
		ExpectError bool
		Assert      func(c1 *domain.Councillor, c2 *domain.Councillor) error
	}{
		{
			Name:        "test saves new councillor",
			Councillor:  makeCouncillor("john", "smith", "jsmith@test.com", "21 somewhere address", "some area", "right of center"),
			ExpectError: false,
			Assert: func(c1 *domain.Councillor, c2 *domain.Councillor) error {
				if c1.Email != c2.Email {
					return fmt.Errorf("epected %s to equal %s ", c1.Email, c2.Email)
				}
				return nil
			},
		},
		{
			Name:        "test updates existing councillor",
			Councillor:  existingCouncillor,
			ExpectError: false,
			Assert: func(c1 *domain.Councillor, c2 *domain.Councillor) error {
				if c1.ID != c2.ID {
					return fmt.Errorf("expected ids to match %s != %s", c1.ID, c2.ID)
				}
				if c1.Email != c2.Email {
					return fmt.Errorf("expected emails to match %s != %s", c1.Email, c2.Email)
				}
				return nil
			},
		},
		{
			Name:        "test wont create duplicate councillor",
			Councillor:  makeCouncillor("john", "smith", "jsmith@test.com", "21 somewhere address", "some area", "right of center"),
			ExpectError: true,
		},
	}
	//run tests
	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			err := councillorRepo.SaveUpdate(v.Councillor)
			if err == nil && v.ExpectError {
				t.Fatal("expected an error but recived none")
			} else if err != nil && !v.ExpectError {
				t.Fatal("did not expect an error but recieved one ", err.Error())
			}
			if v.Assert != nil {
				//assert councillor present in db
				c, err := councillorRepo.FindOneByKeyValue("Email", v.Councillor.Email)
				if err != nil {
					t.Fatal("did not expect an error reading councillor ", err.Error())
				}
				if err := v.Assert(c, v.Councillor); err != nil {
					t.Fatal(err.Error())
				}
			}

		})
	}
}

func TestDeleteCouncillor(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("integration tests not enabled")
	}
	//setup
	tearDown := councillorRepoTestSetup(t)
	defer tearDown()
	councillorRepo := domain.NewCouncillorRepo(config.Conf, domain.NewAdminActor(), &domain.AuthorisationService{})

	tests := []struct {
		Name        string
		Councillor  *domain.Councillor
		ExpectError bool
		Assert      func(c1 *domain.Councillor) error
	}{
		{
			Name:        "test deletes existing councillor",
			Councillor:  existingCouncillor,
			ExpectError: false,
			Assert: func(c *domain.Councillor) error {
				if nil != c {
					return fmt.Errorf("did not expect to find a councillor after delete ")
				}
				return nil
			},
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			err := councillorRepo.DeleteAllByKeyValue("Email", v.Councillor.Email)
			if err == nil && v.ExpectError {
				t.Fatal("expected an error but recived none")
			} else if err != nil && !v.ExpectError {
				t.Fatal("did not expect an error but recieved one ", err.Error())
			}
			if v.Assert != nil {
				//assert councillor present in db
				c, err := councillorRepo.FindOneByKeyValue("Email", v.Councillor.Email)
				if err != nil {
					t.Fatal("did not expect an error reading councillor ", err.Error())
				}
				if err := v.Assert(c); err != nil {
					t.Fatal(err.Error())
				}
			}
		})
	}

}

func TestFindCouncillorByKeyValue(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("integration tests not enabled")
	}
	//setup
	tearDown := councillorRepoTestSetup(t)
	defer tearDown()
	councillorRepo := domain.NewCouncillorRepo(config.Conf, domain.NewAdminActor(), &domain.AuthorisationService{})
	tests := []struct {
		Name        string
		Field       string
		Value       string
		ExpectError bool
		Assert      func(c *domain.Councillor) error
	}{
		{
			Name:        "test find existing councillor",
			Field:       "Email",
			Value:       existingCouncillor.Email,
			ExpectError: false,
			Assert: func(c *domain.Councillor) error {
				if c.Email != existingCouncillor.Email {
					return fmt.Errorf("expected to find correct councillor emails do not match %s != %s", c.Email, existingCouncillor.Email)
				}
				return nil
			},
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			fc, err := councillorRepo.FindOneByKeyValue(v.Field, v.Value)
			if v.ExpectError && err == nil {
				t.Fatal("expected an error but gone none")
			} else if !v.ExpectError && err != nil {
				t.Fatal("unexpected error ", err.Error())
			}
			if err := v.Assert(fc); err != nil {
				t.Fatal("did not expect an assert error ", err.Error())
			}
		})
	}

}
