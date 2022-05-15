package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;type:varchar(64)"`
	LastName  string `gorm:"not null;type:varchar(64)"`
	Username  string `gorm:"not null;type:varchar(64)"`
	Email     string `gorm:"not null;unique;type:varchar(64)"`
	Password  string `gorm:"not null;type:varchar(64)"`
}
