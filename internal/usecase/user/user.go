package user

import (
	"archv1/internal/entity"
	service "archv1/internal/service/user"
	"context"
)

type UserUseCase struct {
	userService service.UserServiceI
}

func NewUserUseCase(service service.UserServiceI) UserUseCaseI {
	return &UserUseCase{
		userService: service,
	}
}

func (u *UserUseCase) List(ctx context.Context, filter entity.Filter) (entity.ListUserResponse, error) {
	userResponse, err := u.userService.List(ctx, filter)
	if err != nil {
		return entity.ListUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserUseCase) GetByID(ctx context.Context, userID int) (entity.GetUserResponse, error) {
	userResponse, err := u.userService.GetByID(ctx, userID)
	if err != nil {
		return entity.GetUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserUseCase) Create(ctx context.Context, user entity.CreateUserRequest) (entity.CreateUserResponse, error) {
	userResponse, err := u.userService.Create(ctx, user)
	if err != nil {
		return entity.CreateUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserUseCase) Update(ctx context.Context, user entity.UpdateUserRequest) (entity.UpdateUserResponse, error) {
	userResponse, err := u.userService.Update(ctx, user)
	if err != nil {
		return entity.UpdateUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserUseCase) UpdateColumns(ctx context.Context, columns entity.UpdateUserColumnsRequest) (entity.UpdateUserResponse, error) {
	userResponse, err := u.userService.UpdateColumns(ctx, columns)
	if err != nil {
		return entity.UpdateUserResponse{}, err
	}

	return userResponse, nil
}

func (u *UserUseCase) Delete(ctx context.Context, userID int) (entity.DeleteUserResponse, error) {
	userResponse, err := u.userService.Delete(ctx, userID)
	if err != nil {
		return entity.DeleteUserResponse{}, err
	}

	return userResponse, nil
}
