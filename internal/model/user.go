package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	FirstName string `gorm:"not null;type:varchar(64)"`
	LastName  string `gorm:"not null;type:varchar(64)"`
	Username  string `gorm:"not null;type:varchar(64)"`
	Email     string `gorm:"not null;unique;type:varchar(64)"`
	Password  string `gorm:"not null;type:varchar(64)"`
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}
