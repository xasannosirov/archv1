package user

import (
	"archv1/internal/entity"
	"archv1/internal/repository/postgres/user"
	"context"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) UserServiceI {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) List(ctx context.Context, filter entity.Filter) (entity.ListUserResponse, error) {
	userResponse, err := u.userRepo.List(ctx, filter)
	if err != nil {
		return entity.ListUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserService) GetByID(ctx context.Context, userID int) (entity.GetUserResponse, error) {
	userResponse, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return entity.GetUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserService) Create(ctx context.Context, user entity.CreateUserRequest) (entity.CreateUserResponse, error) {
	userResponse, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return entity.CreateUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserService) Update(ctx context.Context, user entity.UpdateUserRequest) (entity.UpdateUserResponse, error) {
	userResponse, err := u.userRepo.Update(ctx, user)
	if err != nil {
		return entity.UpdateUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserService) UpdateColumns(ctx context.Context, columns entity.UpdateUserColumnsRequest) (entity.UpdateUserResponse, error) {
	userResponse, err := u.userRepo.UpdateColumns(ctx, columns)
	if err != nil {
		return entity.UpdateUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserService) Delete(ctx context.Context, userID int) (entity.DeleteUserResponse, error) {
	userResponse, err := u.userRepo.Delete(ctx, userID)
	if err != nil {
		return entity.DeleteUserResponse{}, err
	}

	return userResponse, nil
}
