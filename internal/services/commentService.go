package services

import (
	"github.com/jinzhu/gorm"
	"github.com/musishere/Blog/internal/repositories"
)

type CommentService struct {
	repo *repositories.CommentRepository
	db   *gorm.DB
}

func NewCommentService(r *repositories.CommentRepository, db *gorm.DB) *CommentService {
	return &CommentService{repo: r, db: db}
}
