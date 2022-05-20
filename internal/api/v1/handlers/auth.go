package handlers

import (
	"net/http"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/service"
	"voice-out-be/internal/voerrors"

	"github.com/labstack/echo/v4"
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
	r.POST(constants.LoginAPIPath, a.login)
}

func (a *Auth) registerUser(c echo.Context) error {
	registerUserReq := &request.RegisterUserRequest{}

	if err := c.Bind(registerUserReq); err != nil {
		return err
	}

	err := registerUserReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return a.failedAuthResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, err := a.authService.RegisterUser(registerUserReq)

	if err != nil {
		return a.failedAuthResponse(c, code, err, "")
	}

	token, err := a.authService.GenerateJWT(registerUserReq.Username, registerUserReq.Email)

	if err != nil {
		return a.failedAuthResponse(c, "", err, "")
	}

	resp := response.NewResponse(code, token)

	return c.JSON(http.StatusOK, resp)
}

func (a *Auth) login(c echo.Context) error {
	loginReq := &request.LoginRequest{}

	if err := c.Bind(loginReq); err != nil {
		return err
	}

	err := loginReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return a.failedAuthResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, token, err := a.authService.Login(loginReq)

	if err != nil {
		return a.failedAuthResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, token)

	return c.JSON(http.StatusOK, resp)
}

func (a *Auth) failedAuthResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
	if code == "" {
		code = voerrors.MapErrorsToCode(err)
	}

	resp := response.Wrapper{
		ResponseCode: code,
		Status:       code.GetStatus(),
		Message:      code.GetMessage(),
	}

	if errorMsg != "" {
		resp.SetResponseMessage(errorMsg)
	}

	return c.JSON(voerrors.MapErrorsToStatusCode(err), resp)
}
