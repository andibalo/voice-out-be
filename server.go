package voiceout

import (
	"voice-out-be/internal/api/v1/handlers"
	"voice-out-be/internal/config"
	"voice-out-be/internal/service"
	"voice-out-be/internal/storage"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	postService := service.NewPostService(cfg, store)
	userService := service.NewUserService(cfg, store)

	authHandler := handlers.NewAuthHandler(authService)
	postHandler := handlers.NewPostHandler(postService)
	userHandler := handlers.NewUserHandler(userService)

	registerHandlers(e, &handlers.HealthCheck{}, authHandler, postHandler, userHandler)

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
