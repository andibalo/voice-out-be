package handlers

import (
	"fmt"
	"net/http"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/response"
	"voice-out-be/internal/service"
	"voice-out-be/internal/voerrors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *User {

	return &User{
		userService: userService,
	}
}

func (h *User) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.UserAPIPath)

	r.Use(middleware.JWT([]byte("TEST")))
	r.GET(constants.FetchCurrentUserAPIPath, h.getCurrentUser)
}

func (h *User) getCurrentUser(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userId := fmt.Sprintf("%s", claims["userId"])

	code, respData, err := h.userService.FetchCurrentUser(userId)

	if err != nil {
		return h.failedUserResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, respData)

	return c.JSON(http.StatusOK, resp)
}

func (h *User) failedUserResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
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
