package service

import (
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
)

type AuthService interface {
	RegisterUser(registerUserReq *request.RegisterUserRequest) (response.Code, error)
}
