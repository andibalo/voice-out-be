package voiceout

import (
	"github.com/labstack/echo/v4"
	"voice-out-be/internal/config"
	"voice-out-be/internal/storage"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.AppConfig) *Server {

	storage.New(cfg)

	return &Server{
		echo: echo.New(),
	}
}

func (server *Server) Start(addr string) error {
	return server.echo.Start(":" + addr)
}
