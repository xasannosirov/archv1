package user

import (
	"archv1/internal/entity"
	"context"
)

type UserRepository interface {
	List(ctx context.Context, filter entity.Filter) (entity.ListUserResponse, error)
	GetByID(ctx context.Context, userID int) (entity.GetUserResponse, error)
	Create(ctx context.Context, user entity.CreateUserRequest) (entity.CreateUserResponse, error)
	Update(ctx context.Context, user entity.UpdateUserRequest) (entity.UpdateUserResponse, error)
	UpdateColumns(ctx context.Context, user entity.UpdateUserColumnsRequest) (entity.UpdateUserResponse, error)
	Delete(ctx context.Context, userID, deletedBy int) (entity.DeleteUserResponse, error)
}
