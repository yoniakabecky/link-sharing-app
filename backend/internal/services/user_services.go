package services

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/config"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/models"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/pkg/jwt"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/pkg/password"
	"github.com/yoniakabecky/link-sharing-app/backend/internal/repositories"
)

type UserServices struct {
	repo             *repositories.UserRepository
	refreshTokenRepo *repositories.RefreshTokenRepository
}

func NewUserServices(repo *repositories.UserRepository, refreshTokenRepo *repositories.RefreshTokenRepository) *UserServices {
	return &UserServices{
		repo:             repo,
		refreshTokenRepo: refreshTokenRepo,
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

// GetUserByID returns the user as ResponseUser (id, email only).
func (s *UserServices) GetUserByID(ctx context.Context, id int) (*models.ResponseUser, error) {
	u, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.ResponseUser{ID: u.ID, Email: u.Email}, nil
}

// IssueTokens generates a new access token and refresh token for the user, storing the refresh token hash in DB.
func (s *UserServices) IssueTokens(ctx context.Context, user *models.ResponseUser) (accessToken, refreshToken string, err error) {
	cfg := config.Load()
	secret := []byte(cfg.JWT.Key)
	accessToken, err = jwt.GenerateJWT(secret, strconv.Itoa(user.ID))
	if err != nil {
		return "", "", err
	}
	refreshToken, err = generateOpaqueToken()
	if err != nil {
		return "", "", err
	}
	hash := sha256Hash(refreshToken)
	expiresAt := time.Now().Add(time.Duration(cfg.JWT.RefreshExp) * time.Second)
	if err := s.refreshTokenRepo.Create(ctx, user.ID, hash, expiresAt.Format("2006-01-02 15:04:05")); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

// ValidateRefreshAndIssue validates the refresh token, deletes it (rotation), and issues a new token pair.
func (s *UserServices) ValidateRefreshAndIssue(ctx context.Context, token string) (accessToken, refreshToken string, user *models.ResponseUser, err error) {
	hash := sha256Hash(token)
	rt, err := s.refreshTokenRepo.GetByTokenHash(ctx, hash)
	if err != nil {
		return "", "", nil, errors.New("invalid refresh token")
	}
	if time.Now().After(rt.ExpiresAt) {
		_ = s.refreshTokenRepo.DeleteByID(ctx, rt.ID)
		return "", "", nil, errors.New("refresh token expired")
	}
	u, err := s.repo.GetUserByID(ctx, rt.UserID)
	if err != nil {
		return "", "", nil, err
	}
	_ = s.refreshTokenRepo.DeleteByID(ctx, rt.ID)
	user = &models.ResponseUser{ID: u.ID, Email: u.Email}
	accessToken, refreshToken, err = s.IssueTokens(ctx, user)
	if err != nil {
		return "", "", nil, err
	}
	return accessToken, refreshToken, user, nil
}

func generateOpaqueToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func sha256Hash(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}
