package storage

import (
	"voice-out-be/internal/dto"
	"voice-out-be/internal/model"
)

func (s *Store) CreateUser(in *dto.RegisterUser) (*model.User, error) {

	user := &model.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
	}

	err := s.userRepository.SaveUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CreatePost(in *dto.CreatePost) (*model.Post, error) {

	post := &model.Post{
		From:   in.From,
		To:     in.To,
		Body:   in.Body,
		UserID: in.UserID,
	}

	err := s.postRepository.SavePost(post)

	if err != nil {
		return nil, err
	}

	return post, nil
}
