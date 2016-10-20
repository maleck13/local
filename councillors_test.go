package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/domain"
	pt "github.com/maleck13/local/domain/testing"
	"github.com/maleck13/local/test"
)

func init() {
	test.SetUpConfig()
}

func councillorsSetup(t *testing.T) func() {
	councillorRepo := domain.NewCouncillorRepo(config.Conf, domain.NewAdminActor(), &domain.AuthorisationService{})
	var tearDown = func() {
		t.Log("cleaning up councillors_test")
		if err := councillorRepo.DeleteAllByKeyValue("FirstName", "john"); err != nil {
			t.Log("failed to clean up !!!!!!! ")
		}
	}
	for i := 0; i < 2; i++ {
		email := fmt.Sprintf("test@test%d.com", i)
		c := pt.MakeTestCouncillor("john", "smith", email, "21 somewhere", "anarea", "right of center")
		if err := councillorRepo.SaveUpdate(c); err != nil {
			t.Error("error setting up ", err.Error())
		}
	}
	return tearDown
}

func TestListCouncillorsByCountyAndArea(t *testing.T) {
	if !*test.IntegrationEnabled {
		t.Skip("Integration disabled")
	}
	//setup a server
	mux := buildService(config.Conf)
	buildCouncillorController(mux)
	testServer := httptest.NewServer(mux.Mux)
	defer testServer.Close()
	testURL := testServer.URL
	// setup some councillors
	tearDown := councillorsSetup(t)
	defer tearDown()

	tests := []struct {
		Name        string
		Endpoint    string
		StatusCode  int
		ExpectError bool
		Assert      func(cs app.GoaLocalCouncillorCollection) error
	}{
		{
			Name:       "test get councillors by county",
			Endpoint:   "/councillors?county=TESTCOUNTY",
			StatusCode: 200,
			Assert: func(cs app.GoaLocalCouncillorCollection) error {
				if len(cs) != 2 {
					return fmt.Errorf("expected 2 councillors got %d ", len(cs))
				}
				for _, c := range cs {
					if c.County != "TESTCOUNTY" {
						return fmt.Errorf("expected all councillors to have the same county")
					}
				}
				return nil
			},
		},
		{
			Name:       "test get councillors by county and area",
			Endpoint:   `/councillors?county=TESTCOUNTY&area=anarea`,
			StatusCode: 200,
			Assert: func(cs app.GoaLocalCouncillorCollection) error {
				if len(cs) != 2 {
					return fmt.Errorf("expected 2 councillors got %d ", len(cs))
				}
				for _, c := range cs {
					if c.County != "TESTCOUNTY" {
						return fmt.Errorf("expected all councillors to have the same county")
					}
					if c.Area != "anarea" {
						return fmt.Errorf("expected the area to be an area")
					}
				}
				return nil
			},
		},
		{
			Name:       "test get councillors by county and area non found",
			Endpoint:   `/councillors?county=TESTCOUNTY&area="notthere"`,
			StatusCode: 200,
			Assert: func(cs app.GoaLocalCouncillorCollection) error {
				return nil
			},
		},
	}

	for _, tr := range tests {
		t.Run(tr.Name, func(t *testing.T) {
			var councillors = app.GoaLocalCouncillorCollection{}
			endpoint := testURL + tr.Endpoint
			req, err := http.NewRequest("GET", endpoint, nil)
			if err != nil {
				t.Fatal("failed to create request", err.Error())
			}
			pt.AddJwtToken(config.Conf, "local", req)
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal("failed to get councillors", err.Error())
			}
			defer res.Body.Close()
			if res.StatusCode != tr.StatusCode {
				t.Fatalf("incorrect status code %d != %d", res.StatusCode, tr.StatusCode)
			}
			dec := json.NewDecoder(res.Body)
			if err := dec.Decode(&councillors); err != nil {
				t.Fatal("failed to decode councillors", err.Error())
			}
			if err := tr.Assert(councillors); err != nil {
				t.Fatal("asserts failed ", err.Error())
			}
		})
	}

}
