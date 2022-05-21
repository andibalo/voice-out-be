package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	From      string    `json:"from"  gorm:"not null;type:varchar(64)"`
	To        string    `json:"to"  gorm:"not null;type:varchar(64)"`
	Body      string    `json:"body" gorm:"not null;type:varchar(64)"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
