package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/request"
)

type Auth struct{}

func NewAuthHandler() *Auth {

	return &Auth{}
}

func (a *Auth) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.AuthAPIPath)

	r.POST(constants.RegisterAPIPath, a.register)
}

func (a *Auth) register(c echo.Context) error {
	user := new(request.RegisterUserRequest)

	if err := c.Bind(user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
