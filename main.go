//go:generate goagen bootstrap -d github.com/maleck13/local/design

package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
)

var confPath = flag.String("config", "config/config-local.json", "config path")
var port = flag.String("port", ":3001", "port to listen on")

func main() {
	// Create service
	service := goa.New("locals")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(false))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	conf := initConfig()
	setupDb(conf)

	uc := NewUserController(service)
	app.MountUserController(service, uc)

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}

func initConfig() *config.Config {
	conf, err := config.LoadConfig(*confPath)
	if err != nil {
		logrus.Fatal("failed to load config ", err.Error())
	}
	return conf
}

func setupDb(config *config.Config) {
	sess, err := data.AdminSession(config)
	if err != nil {
		logrus.Fatal("failed to get db session ", err.Error())
	}
	if err := data.CreateDb(sess); err != nil {
		logrus.Fatal("failed to create db  ", err.Error())
	}
	if err := data.CreateTables(sess); err != nil {
		logrus.Fatal("failed to create tables  ", err.Error())
	}
	if err := data.CreateAdminUser(config, sess); err != nil {
		logrus.Fatal("failed to admin user  ", err.Error())
	}

}
