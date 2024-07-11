package auth

import (
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
