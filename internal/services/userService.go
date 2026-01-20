package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/musishere/Blog/internal/auth"
	"github.com/musishere/Blog/internal/models"
	"github.com/musishere/Blog/internal/repositories"
)

type UserServive struct {
	repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserServive {
	return &UserServive{repo: r}
}

func (service *UserServive) Register(username, email, password string) (*models.User, error) {
	user := &models.User{
		ID:          uuid.New(),
		UserName:    username,
		Email:       email,
		DisplayName: username,
		Role:        "reader",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	hashedPassword, err := auth.HashedPassword(password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword
	if err := service.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserServive) Login(email, password string) (*models.User, string, error) {
	user, err := service.repo.GetByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if !auth.CheckPassword(user.Password, password) {
		return nil, "", err
	}

	token, _ := auth.GenerateJsonWebToken(user.ID.String(), user.Email, user.Role)
	return user, token, nil

}
