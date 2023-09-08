package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector interface {
	Connect() (*gorm.DB, error)
}

type PostgresConnector struct {
	URI string
}

func NewPostgresConnector(host, dbName, user, password string, port int) *PostgresConnector {
	return &PostgresConnector{
		URI: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			host, user, password, dbName, port),
	}
}

func (pc *PostgresConnector) Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(pc.URI), &gorm.Config{})
}
