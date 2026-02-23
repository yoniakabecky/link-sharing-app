package services

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
)

type UserServices struct {
	repo *repositories.UserRepository
}

func NewUserServices(repo *repositories.UserRepository) *UserServices {
	return &UserServices{
		repo: repo,
	}
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (s *UserServices) Register(ctx context.Context, u *models.RegisterUser) (*models.ResponseUser, error) {
	err := validate.Struct(u)
	if err != nil {
		return nil, errors.New("validation error: " + err.Error())
	}
	return s.repo.Register(ctx, u)
}
