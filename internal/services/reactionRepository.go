package services

import (
	"github.com/musishere/Blog/internal/repositories"
)

type ReactionService struct {
	repo *repositories.ReactionRepository
}

func NewReactionService(r *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{repo: r}
}
