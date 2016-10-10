package domain

import "github.com/goadesign/goa"

// AuthorisationService authorises access to entities
type AuthorisationService struct{}

//Authorise an action on an entity
func (AuthorisationService) Authorise(entity AccessDefinor, action string, actor Actor) error {
	access := entity.AccessTypes()
	if actor.Id() == entity.Owner() {
		return nil
	}
	var who []string
	if action == "write" {
		who = access["write"]

	} else if action == "read" {
		who = access["read"]
	}
	for _, allowed := range who {
		if actor.Type() == allowed || "any" == allowed {
			return nil
		}
	}
	return goa.ErrUnauthorized("Authorise : unauthorised no access")
}
