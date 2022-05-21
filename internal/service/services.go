package service

import (
	"voice-out-be/internal/model"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"

	"go.uber.org/zap"
)

type AuthService interface {
	RegisterUser(registerUserReq *request.RegisterUserRequest) (code response.Code, token string, err error)
	Login(loginReq *request.LoginRequest) (code response.Code, token string, err error)
	GenerateJWT(name string, email string, userID string) (jwtToken string, err error)
}

type PostService interface {
	CreatePost(createPostReq *request.CreatePostRequest, userID string) (code response.Code, err error)
	FetchAllPosts() (code response.Code, posts *[]model.Post, err error)
}

type Config interface {
	Logger() *zap.Logger
}
