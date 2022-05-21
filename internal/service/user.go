package service

import (
	"errors"
	"voice-out-be/internal/model"
	"voice-out-be/internal/response"
	"voice-out-be/internal/storage"
	"voice-out-be/internal/voerrors"

	"go.uber.org/zap"
)

type userService struct {
	config Config
	store  storage.Storage
}

func NewUserService(config Config, store storage.Storage) *userService {

	return &userService{
		config: config,
		store:  store,
	}
}

func (s *userService) FetchCurrentUser(userID string) (response.Code, *response.FetchUserResponse, error) {

	s.config.Logger().Info("FetchCurrentUser: finding user by id")
	user, err := s.store.FindUserByID(userID)

	if err != nil {
		if !errors.Is(err, voerrors.ErrNotFound) {
			s.config.Logger().Error("FetchCurrentUser: error finding user by email", zap.Error(err))
			return response.ServerError, nil, err
		}
	}

	resp := s.mapUserToUserResp(user)

	return response.Success, resp, nil
}

func (s *userService) mapUserToUserResp(user *model.User) *response.FetchUserResponse {

	return &response.FetchUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
}
