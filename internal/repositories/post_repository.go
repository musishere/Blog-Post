package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/musishere/Blog/internal/models"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetAll() ([]models.Post, error) {
	var post []models.Post

	err := r.db.Order("created_at_desc").Find(&post).Error

	return post, err
}

func (r *PostRepository) GetByID(id string) (*models.Post, error) {
	var post models.Post

	err := r.db.Preload("Comments").Preload("Reactions").Preload("Tags").Where("id=?", id).First(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) DeletePostByID(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var post models.Post
		if err := tx.Preload("Tags").First(&post, "id=?", id).Error; err != nil {
			return nil
		}
	})
}
