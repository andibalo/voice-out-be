package repositories

import (
	"voice-out-be/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {

	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) SaveUser(user *model.User) error {

	err := u.db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	err := u.db.Where("email = ?", email).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByID(userID string) (*model.User, error) {
	user := &model.User{}

	err := u.db.Where("id = ?", userID).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}
