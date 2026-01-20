package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/musishere/Blog/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email=?", email).First((&user)).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
