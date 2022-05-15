package storage

import "voice-out-be/internal/model"

func (s *Store) FindUserByEmail(email string) (*model.User, error) {

	user, err := s.userRepository.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
