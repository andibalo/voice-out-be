package service

import (
	"go.uber.org/zap"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
)

type AuthService interface {
	RegisterUser(registerUserReq *request.RegisterUserRequest) (response.Code, error)
	GenerateJWT(name string, email string) (jwtToken string, err error)
}

type Config interface {
	Logger() *zap.Logger
}
