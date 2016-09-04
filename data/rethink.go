package data

import (
	"github.com/maleck13/local/config"
	r "gopkg.in/dancannon/gorethink.v2"
)

const (
	DB_NAME           = "locals"
	COUNCILLORS_TABLE = "local_councillors"
	USER_TABLE        = "users"
)

var (
	session *r.Session
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

func DbSession(conf *config.Config) (*r.Session, error) {
	var err error
	if session == nil {
		session, err = getSession(conf)
	}
	session.Use(DB_NAME)
	return session, err
}
