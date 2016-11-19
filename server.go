package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/CrowdSurge/banner"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/imdario/mergo"
	"github.com/intoxicated/instagram-analyzer-api/api"
	"github.com/intoxicated/instagram-analyzer-api/config"
	"github.com/mangoplate/mp-go-libs/logger"
	"github.com/mangoplate/mp-go-libs/logrus/hook"
)

func initLogger(c *config.Configuration) (*api.Resource, error) {
	appLogLevel, err := logrus.ParseLevel(c.Log.Level)
	if err != nil {
		panic(err)
	}

	appLogger, err := logger.GetLogger(c.Log.Out, c.Log.Path)
	if err != nil {
		return nil, err
	}

	appLogger.Formatter = new(logrus.JSONFormatter)
	appLogger.Hooks = make(logrus.LevelHooks)
	appLogger.Level = appLogLevel
	appLogger.Hooks.Add(&hook.ErrorHook{})

	statLogLevel, err := logrus.ParseLevel(c.StatLog.Level)
	if err != nil {
		panic(err)
	}

	statLogger, err := logger.GetLogger(c.StatLog.Out, c.StatLog.Path)
	if err != nil {
		return nil, err
	}

	statLogger.Formatter = new(logrus.JSONFormatter)
	statLogger.Hooks = make(logrus.LevelHooks)
	statLogger.Level = statLogLevel

	return &api.Resource{
		AppLogger:  appLogger,
		StatLogger: statLogger,
	}, nil
}

func initResource() (*api.Resource, error) {
	c, err := config.LoadConfiguration()
	if err != nil {
		return nil, err
	}

	loggerRes, err := initLogger(c)
	if err != nil {
		return nil, err
	}

	resource := api.Resource{
		Config: c,
		DB:     config.InitDB(c),
	}

	mergo.Merge(&resource, loggerRes)
	return &resource, nil
}

func DefaultServer() *api.Resource {
	// Default resource
	rc, err := initResource()
	if err != nil {
		rc.AppLogger.Panic(err)
	}
	//rc.AppLogger.WithFields(logrus.Fields{
	//	"resource": rc,
	//}).Debug()

	rc.Server = NewServer(rc)
	return rc
}

func NewServer(rc *api.Resource) *http.Server {
	api := api.V1{
		Resource: rc,
		Whitelist: []*regexp.Regexp{
			regexp.MustCompile(`/health_check`),
			regexp.MustCompile(`/login`),
			regexp.MustCompile(`/request_token`),
		},
	}

	r := mux.NewRouter()
	//routings
	r.HandleFunc("/login", api.LogInWithInstagram).Methods("GET")
	r.HandleFunc("/request_token", api.Authorize).Methods("GET")
	r.HandleFunc("/search", api.SearchMedia).Methods("GET")

	//add middlewares
	router := api.Authenticate(r)
	router = api.AddDefaultHeader(router)
	router = handlers.LoggingHandler(os.Stdout, router)
	//router = handlers.CORS()(router)

	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0:%d", rc.Config.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server
}

func main() {
	srv := DefaultServer()
	banner.Print("intoxicated")
	fmt.Print("==============================================================\n")
	fmt.Print("=====               Instagram analyzer API               =====\n")
	fmt.Print("=====               Created by intoxicated               =====\n")
	fmt.Print("=====                   Version 0.0.1                    =====\n")
	fmt.Print("==============================================================\n")

	log.Fatal(srv.Server.ListenAndServe())

}
