package storage

import (
	"errors"
	"gorm.io/gorm"
	"voice-out-be/internal/model"
	"voice-out-be/internal/voerrors"
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
