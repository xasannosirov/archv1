package auth

import (
	"archv1/internal/service/auth"
	"context"
)

type AuthUseCase struct {
	authService auth.AuthServiceI
}

func NewAuthUseCase(service auth.AuthServiceI) AuthUseCaseI {
	return &AuthUseCase{
		authService: service,
	}
}

func (a *AuthUseCase) UniqueUsername(ctx context.Context, username string) (bool, error) {
	return a.authService.UniqueUsername(ctx, username)
}

func (a *AuthUseCase) UpdateToken(ctx context.Context, id int, token string) error {
	return a.authService.UpdateToken(ctx, id, token)
}
