package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/service"
)

type Auth struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *Auth {

	return &Auth{
		authService: authService,
	}
}

func (a *Auth) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.AuthAPIPath)

	r.POST(constants.RegisterAPIPath, a.registerUser)
}

func (a *Auth) registerUser(c echo.Context) error {
	registerUserReq := &request.RegisterUserRequest{}

	if err := c.Bind(registerUserReq); err != nil {
		return err
	}

	code, err := a.authService.RegisterUser(registerUserReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error")
	}

	resp := response.NewResponse(code, "")

	return c.JSON(http.StatusOK, resp)
}
