package auth

import (
	"archv1/internal/entity"
	"archv1/internal/repository/postgres/auth"
	"context"
)

type AuthService struct {
	authRepo auth.AuthRepository
}

func NewAuthService(menuRepo auth.AuthRepository) AuthServiceI {
	return &AuthService{
		authRepo: menuRepo,
	}
}

func (m *AuthService) UniqueUsername(ctx context.Context, username string) (bool, error) {
	return m.authRepo.UniqueUsername(ctx, username)
}

func (a *AuthService) UpdateToken(ctx context.Context, id int, token string) error {
	return a.authRepo.UpdateToken(ctx, id, token)
}

func (a *AuthService) GetUserByUsername(ctx context.Context, username string) (entity.GetUserResponse, error) {
	return a.authRepo.GetUserByUsername(ctx, username)
}

func (a *AuthService) GetUserByToken(ctx context.Context, token string) (entity.GetUserResponse, error) {
	return a.authRepo.GetUserByToken(ctx, token)
}
