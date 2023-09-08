package api

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func (s *Server) Run() {
	router := NewRouter(s.DB).build()
	logrus.Fatal(router.StartAutoTLS(":8080"))
}
