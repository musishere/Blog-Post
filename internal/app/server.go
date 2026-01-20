package app

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/musishere/Blog/config"
)

type Server struct {
	E   *echo.Echo
	DB  *gorm.DB
	Cfg *config.AppConfig
}

func StartServer(cfg *config.AppConfig, db *gorm.DB) *Server {
	e := echo.New()
	RegisterRoutes(e, db, cfg)

	return &Server{
		E: e, DB: db, Cfg: cfg,
	}
}

func (s *Server) Start(addr string) error {
	return s.E.Start(addr)
}
