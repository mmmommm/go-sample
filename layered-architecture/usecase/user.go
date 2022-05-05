package usecase

import (
	"context"

	"github.com/mmmommm/go-sample/layered-architecture/domain"
	"github.com/mmmommm/go-sample/layered-architecture/repository"
)

type UserGetCaseInput struct {
	UserID string
}

type UserGetCase func(ctx context.Context, in *UserGetCaseInput) (*domain.User, error)

func ProvideUserGetCase(repo repository.UserRepository) UserGetCase {
	return func(ctx context.Context, in *UserGetCaseInput) (*domain.User, error) {
		repo.FindByID(in.UserID)
		return nil, nil
	}
}