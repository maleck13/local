package data

import (
	"github.com/Sirupsen/logrus"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/errors"
	"golang.org/x/crypto/bcrypt"
	r "gopkg.in/dancannon/gorethink.v2"
)

const (
	DB_NAME           = "locals"
	COUNCILLORS_TABLE = "local_councillors"
	USER_TABLE        = "users"
)

var (
	session *r.Session
	tables  = []string{COUNCILLORS_TABLE, USER_TABLE}
)

func getSession(conf *config.Config) (*r.Session, error) {
	var err error
	if nil == session {
		session, err = r.Connect(r.ConnectOpts{
			Addresses: conf.Database.Addrs,
		})
		if err != nil {
			return nil, err
		}
	}
	return session, nil
}

func AdminSession(conf *config.Config) (*r.Session, error) {
	return getSession(conf)
}

func DbSessionDestroy() {
	if nil != session {
		session.Close()
	}
}

func CreateDb(session *r.Session) error {
	lDbs, err := r.DBList().Run(session)
	if err != nil {
		return errors.NewServiceError("failed to list dbs ", 500)
	}

	var response []string

	if err := lDbs.All(&response); err != nil {
		return errors.NewServiceError("failed to list dbs ", 500)
	}

	for _, db := range response {
		if db == DB_NAME {
			logrus.Info("db exists skipping creation")
			return nil
		}
	}
	logrus.Info("db does not exist creating new db " + DB_NAME)
	_, err = r.DBCreate(DB_NAME).RunWrite(session)
	if err != nil {
		return errors.NewServiceError("failed to list dbs ", 500)
	}
	return nil
}

func CreateTables(session *r.Session) error {
	session.Use(DB_NAME)
	tCurs, err := r.TableList().Run(session)
	if err != nil {
		return errors.NewServiceError("failed to list tables ", 500)
	}
	var existing []string
	if err := tCurs.All(&existing); err != nil {
		return errors.NewServiceError("failed to list dbs ", 500)
	}
	for _, t := range tables {
		tFound := false
		for _, e := range existing {
			if e == t {
				tFound = true
			}
		}
		if !tFound {
			logrus.Info("creating table ", t)
			if _, err := r.TableCreate(t).RunWrite(session); err != nil {
				return errors.NewServiceError("failed to create table "+t+" "+err.Error(), 500)
			}
			if t == COUNCILLORS_TABLE {
				if _, err := r.Table(COUNCILLORS_TABLE).IndexCreate("GeoLoc", r.IndexCreateOpts{Geo: true}).RunWrite(session); err != nil {
					return errors.NewServiceError("failed to create index on table "+t+err.Error(), 500)
				}
			}
		}
	}
	return nil
}

func CreateAdminUser(config *config.Config, session *r.Session) error {
	userRepo := UserRepo{Config: config}
	adminUser := config.Admin.User
	adminPass := config.Admin.Auth
	exist, err := userRepo.FindOneByFieldAndValue("Email", adminUser)
	if err != nil {
		return err
	}
	if nil != exist {
		return nil
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(adminPass), 10)
	if err != nil {
		return err
	}
	session.Use(DB_NAME)
	adminUserModel := &User{
		User:&app.User{
			Email: adminUser,
			Type:  "admin",
			Token: string(enc),
		},
	}
	if err := userRepo.Save(adminUserModel); err != nil {
		return err
	}
	return nil

}

func DropDb(config *config.Config) error {
	session, err := getSession(config)
	if err != nil {
		return err
	}
	_, err = r.DBDrop(DB_NAME).RunWrite(session)
	if err != nil {
		return errors.NewServiceError("failed to drop db "+err.Error(), 500)
	}
	return nil
}

func DbSession(conf *config.Config) (*r.Session, error) {
	var err error
	if session == nil {
		session, err = getSession(conf)
	}
	session.Use(DB_NAME)
	return session, err
}
