package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		Database Database `envPrefix:"DATABASE_"`
		API      API      `envPrefix:"API_"`
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}

	API struct {
		Port int
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	opts := env.Options{UseFieldNameByDefault: true}

	// Load env vars.
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Print the loaded data.
	logrus.Debugf("%+v\n", cfg)

	return cfg, nil
}
