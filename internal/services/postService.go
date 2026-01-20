package services

import (
	"github.com/jinzhu/gorm"
	"github.com/musishere/Blog/internal/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
	db   *gorm.DB
}

func NewPostService(r *repositories.PostRepository, db *gorm.DB) *PostService {
	return &PostService{
		repo: r,
		db:   db,
	}
}
