package testing

import (
	"fmt"

	"github.com/maleck13/local/app"
	"github.com/maleck13/local/domain"
)

const (
	TESTCOUNTY = "TESTCOUNTY"
)

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

func (muf MockUserFinder) FindOneByTypeAndEmail(uType, email string) (*domain.User, error) {
	fmt.Println("MockUserFinder", muf.user)
	return muf.user, muf.err
}

// NewUserFinder returns configured UserFinder
func NewUserFinder(u *domain.User, us []*domain.User, err error) domain.UserFinder {
	return MockUserFinder{
		user:  u,
		users: us,
		err:   err,
	}
}

// MockUserFinderSaver implements UserFinderSaver interface
type MockUserFinderSaver struct {
	MockUserFinder
	SaveUpdateAssert func(u *domain.User)
}

// SaveUpdate mock implements UserSaver
func (mfs MockUserFinderSaver) SaveUpdate(u *domain.User) error {
	if mfs.err != nil {
		return mfs.err
	}
	if "" == u.ID {
		u.ID = "mockid"
	}
	if mfs.SaveUpdateAssert != nil {
		mfs.SaveUpdateAssert(u)
	}
	return nil
}

// NewUserFinderSaver configures and creates a new UserFinderSaver
func NewUserFinderSaver(u *domain.User, us []*domain.User, err error) MockUserFinderSaver {
	mf := MockUserFinder{
		user:  u,
		users: us,
		err:   err,
	}
	return MockUserFinderSaver{MockUserFinder: mf}
}

func MakeTestUser(fn, sn, email, area, uType, id string) *domain.User {
	appUser := &app.User{
		ID:         id,
		FirstName:  fn,
		SecondName: sn,
		Email:      email,
		Area:       area,
		Type:       uType,
	}
	return domain.NewUser(appUser)
}

func MakeTestUpdateUser(id, area, email, fn, sn string) *app.UpdateUser {
	return &app.UpdateUser{
		Area:       area,
		Email:      email,
		FirstName:  fn,
		ID:         id,
		SecondName: sn,
	}
}

func MakeTestCouncillor(fn, sn, email, address, area, party string) *domain.Councillor {
	appCouncillor := &app.GoaLocalCouncillor{
		FirstName:  fn,
		SecondName: sn,
		Email:      email,
		Address:    address,
		Area:       area,
		Party:      party,
		County:     TESTCOUNTY,
	}
	return domain.NewCouncillor(appCouncillor)
}
