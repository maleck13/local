package councillor

import (
	"errors"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/domain"
)

type CouncillorSaver interface {
	SaveUpdate(*domain.Councillor) error
}

type CouncillorFinder interface {
	FindOneByKeyValue(key string, value interface{}) (*domain.Councillor, error)
}

type CouncillorSaverFinder interface {
	CouncillorSaver
	CouncillorFinder
}

type councillorService struct {
	CouncillorSaverFinder CouncillorSaverFinder
}

func NewService(sf CouncillorSaverFinder) councillorService {
	return councillorService{
		CouncillorSaverFinder: sf,
	}
}

func (cs councillorService) SetCouncillorUID(cid, uid string) error {
	existing, err := cs.CouncillorSaverFinder.FindOneByKeyValue("id", cid)
	if err != nil {
		return err
	}
	existing.UserID = uid
	return cs.CouncillorSaverFinder.SaveUpdate(existing)
}

func (cs councillorService) Update(id string, incoming *app.CouncillorUpdate) (*domain.Councillor, error) {
	existing, err := cs.CouncillorSaverFinder.FindOneByKeyValue("id", id)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("no existing councillor")
	}
	existing.Address = incoming.Address
	existing.Area = incoming.Area
	existing.Email = incoming.Email
	existing.Facebook = incoming.Facebook
	existing.FirstName = incoming.FirstName
	existing.SecondName = incoming.SecondName
	existing.Twitter = incoming.Twitter
	existing.Phone = incoming.Phone
	existing.Web = incoming.Web
	if err := cs.CouncillorSaverFinder.SaveUpdate(existing); err != nil {
		return nil, err
	}
	return existing, nil
}
