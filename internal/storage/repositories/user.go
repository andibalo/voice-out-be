package repositories

import (
	"gorm.io/gorm"
	"voice-out-be/internal/model"
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