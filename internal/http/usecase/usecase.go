package usecase

import "backend/internal/http/repository"

type UseCase struct {
	Repository *repository.Repository
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Repository: repo,
	}
}
