package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/intoxicated/instagram-analyzer-api/config"
)

type Resource struct {
	Config     *config.Configuration
	DB         *config.DB
	Server     *http.Server
	AppLogger  *logrus.Logger
	StatLogger *logrus.Logger

	AccessToken string
}
