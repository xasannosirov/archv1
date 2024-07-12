package auth

import (
	"archv1/internal/entity"
	"context"
)

type AuthUseCaseI interface {
	UniqueUsername(ctx context.Context, username string) (bool, error)
	UpdateToken(ctx context.Context, id int, token string) error
	GetUserByUsername(ctx context.Context, username string) (entity.GetUserResponse, error)
}
