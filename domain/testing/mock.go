package testing

import "github.com/maleck13/local/domain"

// MockUserFinder is a mock implemenation of a  domain.UserFinder
type MockUserFinder struct {
	user  *domain.User
	users []*domain.User
	err   error
}

// FindOneByFieldAndValue mock implements UserFinder
func (muf MockUserFinder) FindOneByFieldAndValue(field, val string) (*domain.User, error) {
	return muf.user, muf.err
}

// FindAllByTypeAndArea mock implements UserFinder
func (muf MockUserFinder) FindAllByTypeAndArea(uType, area string) ([]*domain.User, error) {
	return muf.users, muf.err
}

// NewUserFinder returns configured UserFinder
func NewUserFinder(u *domain.User, us []*domain.User, err error) domain.UserFinder {
	return MockUserFinder{
		user:  u,
		users: us,
		err:   err,
	}
}
