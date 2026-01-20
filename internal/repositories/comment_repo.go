package repositories

import "github.com/jinzhu/gorm"

type CommentRepository struct {
	db *gorm.DB
}
