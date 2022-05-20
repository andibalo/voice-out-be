package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        string `gorm:"primaryKey"`
	From      string `gorm:"not null;type:varchar(64)"`
	To        string `gorm:"not null;type:varchar(64)"`
	Body      string `gorm:"not null;type:varchar(64)"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
