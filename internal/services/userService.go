package services

import "github.com/musishere/Blog/internal/repositories"

type UserServive struct {
	repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserServive {
	return &UserServive{repo: r}
}
