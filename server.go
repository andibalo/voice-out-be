package voiceout

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"voice-out-be/internal/api/v1/handlers"
	"voice-out-be/internal/config"
	"voice-out-be/internal/service"
	"voice-out-be/internal/storage"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(cfg *config.AppConfig) *Server {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	store := storage.New(cfg)

	authService := service.NewAuthService(cfg, store)

	authHandler := handlers.NewAuthHandler(authService)

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
