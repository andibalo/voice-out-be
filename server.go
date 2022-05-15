package voiceout

import (
	"github.com/labstack/echo/v4"
	"voice-out-be/internal/api/v1/handlers"
	"voice-out-be/internal/config"
	"voice-out-be/internal/storage"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.AppConfig) *Server {

	e := echo.New()

	storage.New(cfg)

	authHandler := handlers.NewAuthHandler()

	registerHandlers(e, &handlers.HealthCheck{}, authHandler)

	return &Server{
		echo: e,
	}
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(":" + addr)
}

type Handler interface {
	AddRoutes(e *echo.Echo)
}

func registerHandlers(e *echo.Echo, handlers ...Handler) {
	for _, handler := range handlers {
		handler.AddRoutes(e)
	}
}
