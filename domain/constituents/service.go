package constituents

import (
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/domain"
)

type ConstiuentFinder interface {
	FindAllByTypeAndArea(uType, area string) ([]*domain.User, error)
}

type CouncillorFinder interface {
	FindOneByKeyValue(key string, value interface{}) (*domain.Councillor, error)
}

type CommsFinder interface {
	FindFirstOpenBySenderAndReceipent(sender, receipient string) (*domain.Communication, error)
}

type Service struct {
	ConstituentFinder ConstiuentFinder
	CouncillorFinder  CouncillorFinder
	CommsFinder       CommsFinder
}

func userToConstituent(u *domain.User) *app.GoaLocalConsituentsNocomms {
	return &app.GoaLocalConsituentsNocomms{
		ID:         &u.ID,
		FirstName:  &u.FirstName,
		SecondName: &u.SecondName,
	}
}

func (s Service) ConsistuentsForCouncillor(cid string) (app.GoaLocalConsituentsNocommsCollection, error) {
	//TODO this is highly inefficient and should use joins
	councillor, err := s.CouncillorFinder.FindOneByKeyValue("id", cid)
	if err != nil {
		return nil, err
	}
	users, err := s.ConstituentFinder.FindAllByTypeAndArea("local", councillor.Area)
	if err != nil {
		return nil, err
	}
	constituents := app.GoaLocalConsituentsNocommsCollection{}
	for _, u := range users {
		c := userToConstituent(u)
		has, err := s.HasOpenComms(u.ID, councillor.ID)
		if err != nil {
			return nil, err
		}
		c.HasOpenComms = &has
		constituents = append(constituents, c)
	}
	return constituents, nil
}

// HasOpenComms true if a constituent has open comms with a councillor
func (s Service) HasOpenComms(sender, cid string) (bool, error) {
	c, err := s.CommsFinder.FindFirstOpenBySenderAndReceipent(sender, cid)
	if err != nil {
		return false, err
	}
	if c != nil {
		return true, nil
	}
	return false, nil
}
