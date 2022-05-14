package voiceout

import (
	"github.com/labstack/echo/v4"
	"voice-out-be/internal/config"
)

type Server struct {
	echo   *echo.Echo
	config *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	return &Server{
		echo:   echo.New(),
		config: cfg,
	}
}

func (server *Server) Start(addr string) error {
	return server.echo.Start(":" + addr)
}
