package service

import (
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"

	"go.uber.org/zap"
)

type AuthService interface {
	RegisterUser(registerUserReq *request.RegisterUserRequest) (response.Code, error)
	Login(loginReq *request.LoginRequest) (code response.Code, token string, err error)
	GenerateJWT(name string, email string) (jwtToken string, err error)
}

type Config interface {
	Logger() *zap.Logger
}
