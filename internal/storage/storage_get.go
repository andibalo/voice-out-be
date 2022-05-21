package storage

import (
	"errors"
	"voice-out-be/internal/model"
	"voice-out-be/internal/voerrors"

	"gorm.io/gorm"
)

func (s *Store) FindUserByEmail(email string) (*model.User, error) {

	user, err := s.userRepository.GetUserByEmail(email)

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, voerrors.ErrNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *Store) FindAllPosts() (*[]model.Post, error) {

	posts, err := s.postRepository.GetAllPosts()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Store) FindAllPostsByUserID(userID string) (*[]model.Post, error) {

	posts, err := s.postRepository.GetAllPostsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
