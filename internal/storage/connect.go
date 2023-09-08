package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/awesomeProject/internal/config"
)

type Connector interface {
	Connect() (*gorm.DB, error)
}

type PostgresConnector struct {
	URI string
}

func NewPostgresConnector(cfg config.Database) *PostgresConnector {
	return &PostgresConnector{
		URI: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port),
	}
}

func (pc *PostgresConnector) Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(pc.URI), &gorm.Config{})
}
