package repositories

import (
	"voice-out-be/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {

	return &PostRepository{
		db: db,
	}
}

func (p *PostRepository) SavePost(post *model.Post) error {

	err := p.db.Create(post).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *PostRepository) GetAllPosts() (*[]model.Post, error) {

	var posts *[]model.Post

	err := p.db.Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostRepository) GetAllPostsByUserID(userID string) (*[]model.Post, error) {

	var posts *[]model.Post

	err := p.db.Where("user_id = ?", userID).Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}
