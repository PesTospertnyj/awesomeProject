package main

import (
	"github.com/sirupsen/logrus"

	"github.com/awesomeProject/internal/api"
	"github.com/awesomeProject/internal/config"
	"github.com/awesomeProject/internal/storage"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	db, err := storage.NewPostgresConnector(cfg.Database).Connect()
	if err != nil {
		logrus.Fatal(err)
	}

	api.NewServer(db, cfg.API).Run()
}
