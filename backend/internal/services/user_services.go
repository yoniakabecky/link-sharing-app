package services

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/pkg/password"
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

func (s *UserServices) Register(ctx context.Context, u *models.UserAuthInput) (*models.ResponseUser, error) {
	err := validate.Struct(u)
	if err != nil {
		return nil, errors.New("validation error: " + err.Error())
	}

	hash, err := password.HashPassword(u.Password)
	if err != nil {
		return nil, errors.New("error hashing password: " + err.Error())
	}
	u.Password = hash

	return s.repo.Register(ctx, u)
}

func (s *UserServices) Login(ctx context.Context, u *models.UserAuthInput) (*models.ResponseUser, error) {
	err := validate.Struct(u)
	if err != nil {
		return nil, errors.New("validation error: " + err.Error())
	}

	user, err := s.repo.GetUserByEmail(ctx, u.Email)
	if err != nil {
		return nil, errors.New("user not found: " + err.Error())
	}

	err = password.ComparePassword(u.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid password: " + err.Error())
	}

	ru := models.ResponseUser{
		ID:    user.ID,
		Email: user.Email,
	}

	return &ru, nil
}
