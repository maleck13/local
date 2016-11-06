package councillor

import (
	"errors"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/domain"
)

type Saver interface {
	SaveUpdate(*domain.Councillor) error
}

type Finder interface {
	FindOneByKeyValue(key string, value interface{}) (*domain.Councillor, error)
}

type SaverFinder interface {
	Saver
	Finder
}

type councillorService struct {
	SaverFinder SaverFinder
}

func NewService(sf SaverFinder) councillorService {
	return councillorService{
		SaverFinder: sf,
	}
}

func (cs councillorService) SetCouncillorUID(cid, uid string) error {
	existing, err := cs.SaverFinder.FindOneByKeyValue("id", cid)
	if err != nil {
		return err
	}
	existing.UserID = uid
	return cs.SaverFinder.SaveUpdate(existing)
}

func (cs councillorService) Update(id string, incoming *app.CouncillorUpdate) (*domain.Councillor, error) {
	existing, err := cs.SaverFinder.FindOneByKeyValue("id", id)
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
	if err := cs.SaverFinder.SaveUpdate(existing); err != nil {
		return nil, err
	}
	return existing, nil
}
