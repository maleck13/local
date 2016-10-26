//go:generate goagen bootstrap -d github.com/maleck13/local/design

package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/maleck13/local/app"
	"github.com/maleck13/local/config"
	"github.com/maleck13/local/data"
	"github.com/maleck13/local/domain"
)

var confPath = flag.String("config", "config/config-local.json", "config path")
var port = flag.String("port", ":3000", "port to listen on")

func buildService(conf *config.Config) *goa.Service {

	// Create service
	service := goa.New("locals")
	// Mount middleware
	service.Mux.Handle("GET", "/", assetHandler())
	service.Mux.Handle("GET", "/*.js", assetHandler())
	service.Mux.Handle("GET", "/app/*", assetHandler())
	service.Mux.Handle("GET", "/vendor/*", assetHandler())

	service.Use(middleware.RequestIDWithHeader(middleware.RequestIDHeader))
	service.Use(middleware.LogRequest(false))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseJWTMiddleware(service, NewJWTMiddleware(conf))
	return service
}

func buildCommunicationsController(service *goa.Service) {
	app.MountCommunicationsController(service, NewCommunicationsController(service))
}

func buildCouncillorController(service *goa.Service) {
	app.MountCouncillorsController(service, NewCouncillorsController(service))
}

func buildSwaggerController(service *goa.Service) {
	app.MountSwaggerController(service, NewSwaggerController(service))
}

func buildAdminController(service *goa.Service) {
	app.MountAdminController(service, NewAdminController(service))
}

func buildUserController(service *goa.Service) {
	app.MountUserController(service, NewUserController(service))
}

func main() {
	//set up config
	conf := initConfig()
	setupDb(conf)
	service := buildService(conf)
	buildCommunicationsController(service)
	buildCouncillorController(service)
	buildSwaggerController(service)
	buildAdminController(service)
	buildUserController(service)
	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}

func assetHandler() goa.MuxHandler {
	base := "./web/dist/"
	h := http.FileServer(http.Dir(base))
	return func(rw http.ResponseWriter, req *http.Request, v url.Values) {
		fmt.Println(req.URL.Path)
		if req.URL.Path == "/login" {
			req.URL.Path = "/"
		}
		if req.URL.Path == "/passwordreset" {
			req.URL.Path = "/"
		}
		h.ServeHTTP(rw, req)
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
	if err := domain.CreateDb(sess); err != nil {
		logrus.Fatal("failed to create db  ", err.Error())
	}
	if err := domain.CreateTables(sess); err != nil {
		logrus.Fatal("failed to create tables  ", err.Error())
	}
	if err := domain.CreateAdminUser(config, sess); err != nil {
		logrus.Fatal("failed to admin user  ", err.Error())
	}

}
