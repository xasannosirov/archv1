package auth

import (
	"context"
)

type AuthRepository interface {
	UniqueUsername(ctx context.Context, username string) (bool, error)
	UpdateToken(ctx context.Context, id int, token string) error
}
