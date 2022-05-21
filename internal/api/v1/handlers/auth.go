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

func (h *Auth) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.AuthAPIPath)

	r.POST(constants.RegisterAPIPath, h.registerUser)
	r.POST(constants.LoginAPIPath, h.login)
}

func (h *Auth) registerUser(c echo.Context) error {
	registerUserReq := &request.RegisterUserRequest{}

	if err := c.Bind(registerUserReq); err != nil {
		return err
	}

	err := registerUserReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedAuthResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, token, err := h.authService.RegisterUser(registerUserReq)

	if err != nil {
		return h.failedAuthResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, token)

	return c.JSON(http.StatusOK, resp)
}

func (h *Auth) login(c echo.Context) error {
	loginReq := &request.LoginRequest{}

	if err := c.Bind(loginReq); err != nil {
		return err
	}

	err := loginReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedAuthResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, token, err := h.authService.Login(loginReq)

	if err != nil {
		return h.failedAuthResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, token)

	return c.JSON(http.StatusOK, resp)
}

func (h *Auth) failedAuthResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
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
