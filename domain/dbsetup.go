package domain

import (
	"github.com/Sirupsen/logrus"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/maleck13/local/errors"
	"golang.org/x/crypto/bcrypt"
	r "gopkg.in/dancannon/gorethink.v2"
)

var tables = []string{data.COUNCILLORS_TABLE, data.USER_TABLE}

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
		if db == data.DB_NAME {
			logrus.Info("db exists skipping creation")
			return nil
		}
	}
	logrus.Info("db does not exist creating new db " + data.DB_NAME)
	_, err = r.DBCreate(data.DB_NAME).RunWrite(session)
	if err != nil {
		return errors.NewServiceError("failed to list dbs ", 500)
	}
	return nil
}

func CreateTables(session *r.Session) error {
	session.Use(data.DB_NAME)
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
			if t == data.COUNCILLORS_TABLE {
				if _, err := r.Table(data.COUNCILLORS_TABLE).IndexCreate("GeoLoc", r.IndexCreateOpts{Geo: true}).RunWrite(session); err != nil {
					return errors.NewServiceError("failed to create index on table "+t+err.Error(), 500)
				}
			}
		}
	}
	return nil
}

func CreateAdminUser(config *config.Config, session *r.Session) error {
	userRepo := UserRepo{Config: config, Authorisor: Access{}, Actor: AdminActor{}}
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
	session.Use(data.DB_NAME)
	adminUserModel := &User{
		User: &app.User{
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
	session, err := data.AdminSession(config)
	if err != nil {
		return err
	}
	_, err = r.DBDrop(data.DB_NAME).RunWrite(session)
	if err != nil {
		return errors.NewServiceError("failed to drop db "+err.Error(), 500)
	}
	return nil
}
