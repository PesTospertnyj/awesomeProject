package api

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/awesomeProject/internal/config"
)

type Server struct {
	DB     *gorm.DB
	config config.API
}

func NewServer(db *gorm.DB, cfg config.API) *Server {
	return &Server{DB: db, config: cfg}
}

func (s *Server) Run() {
	router := NewRouter(s.DB).build()
	logrus.Fatal(router.StartAutoTLS(fmt.Sprintf(":%d", s.config.Port)))
}
