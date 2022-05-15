package service

import (
	"voice-out-be/internal/dto"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/storage"
	"voice-out-be/internal/util"
)

type authService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) *authService {

	return &authService{
		storage: storage,
	}
}

func (a *authService) RegisterUser(registerUserReq *request.RegisterUserRequest) (response.Code, error) {

	hashedPassword, err := util.HashPassword(registerUserReq.Password)

	if err != nil {
		return response.ServerError, err
	}

	user := &dto.RegisterUser{
		FirstName: registerUserReq.FirstName,
		LastName:  registerUserReq.LastName,
		Username:  registerUserReq.Username,
		Email:     registerUserReq.Email,
		Password:  hashedPassword,
	}

	_, err = a.storage.CreateUser(user)

	if err != nil {
		return response.ServerError, err
	}

	return response.Success, nil
}
