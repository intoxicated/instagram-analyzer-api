package config

import "github.com/vrischmann/envconfig"

type Configuration struct {
	Server struct {
		Port uint `envconfig:"default=8080"`
	}

	DB struct {
		Host         string
		Port         uint `envconfig:"default=3306"`
		User         string
		Password     string
		Name         string
		Charset      string `envconfig:"default=utf8mb4"`
		Location     string `envconfig:"default=UTC"`
		MaxIdleConns int    `envconfig:"default=10"`
		MaxOpenConns int    `envconfig:"default=10"`
		LogMode      bool   `envconfig:"default=true"`
	}

	Instagram struct {
		ClientId     string
		ClientSecret string
		AuthUrl      string
		TokenUrl     string
		RedirectUri  string
		ResponseType string
		Scope        string
	}

	Log struct {
		Out   string `envconfig:"default=file" valid:"matches(stdout|stderr|file)"`
		Path  string `envconfig:"default=application.log"`
		Level string `envconfig:"default=info"`
	}

	StatLog struct {
		Out   string `envconfig:"default=stdout" valid:"matches(stdout|stderr|file)"`
		Path  string `envconfig:"default=stats.log"`
		Level string `envconfig:"default=info"`
	}
}

func LoadConfiguration() (c *Configuration, err error) {
	err = envconfig.Init(&c)
	return
}
