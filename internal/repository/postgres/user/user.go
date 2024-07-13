package user

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Repo struct {
	DB *postgres.DB
}

func NewUserRepo(DB *postgres.DB) UserRepository {
	return &Repo{
		DB: DB,
	}
}

func (r *Repo) List(ctx context.Context, filter entity.Filter) (entity.ListUserResponse, error) {
	var response entity.ListUserResponse
	offset := filter.Limit * (filter.Page - 1)

	selectQuery := `
	SELECT
		id,
		username,
		password,
		role,
		status,
		refresh
	FROM users
	`

	whereQuery := ` WHERE deleted_at IS NULL AND status = TRUE`

	limitQuery := fmt.Sprintf(" LIMIT %d OFFSET %d", filter.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+whereQuery+limitQuery)
	if err != nil {
		return entity.ListUserResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.GetUserResponse
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Status,
			&user.Refresh,
		)
		if err != nil {
			return entity.ListUserResponse{}, err
		}

		response.Users = append(response.Users, &user)
	}

	if err := rows.Err(); err != nil {
		return entity.ListUserResponse{}, err
	}

	totalQuery := `SELECT COUNT(*) FROM users WHERE deleted_at IS NULL AND status = TRUE`
	if err := r.DB.QueryRowContext(ctx, totalQuery).Scan(&response.Total); err != nil {
		return entity.ListUserResponse{}, err
	}

	return response, nil
}

func (r *Repo) GetByID(ctx context.Context, userID int) (entity.GetUserResponse, error) {
	var response entity.GetUserResponse

	selectQuery := `
	SELECT 
		id, 
		username, 
		password, 
		role,
		status,
		refresh
	FROM users`

	whereQuery := ` WHERE deleted_at IS NULL AND status = TRUE AND id = ?`

	err := r.DB.QueryRowContext(ctx, selectQuery+whereQuery, userID).Scan(
		&response.Id,
		&response.Username,
		&response.Password,
		&response.Role,
		&response.Status,
		&response.Refresh,
	)
	if err != nil {
		return entity.GetUserResponse{}, err
	}

	return response, nil
}

func (r *Repo) Create(ctx context.Context, user entity.CreateUserRequest) (entity.CreateUserResponse, error) {
	var response entity.CreateUserResponse

	err := r.DB.NewInsert().
		Model(&entity.Users{
			Username: user.Username,
			Password: user.Password,
			Role:     user.Role,
			Status:   user.Status,
		}).
		Returning("id, username, password, role, status").
		Scan(ctx, &response.Id, &response.Username, &response.Password, &response.Role, &response.Status)

	if err != nil {
		return entity.CreateUserResponse{}, err
	}

	reSetQuery := "UPDATE users SET created_by = ? WHERE username = ?"

	_, err = r.DB.ExecContext(ctx, reSetQuery, response.Id, user.Username)
	if err != nil {
		return entity.CreateUserResponse{}, err
	}

	return response, nil
}

func (r *Repo) Update(ctx context.Context, user entity.UpdateUserRequest) (entity.UpdateUserResponse, error) {
	var response entity.UpdateUserResponse

	err := r.DB.NewUpdate().
		Model(&entity.Users{
			ID:        &user.Id,
			Username:  user.Username,
			Password:  user.Password,
			Role:      user.Role,
			Status:    user.Status,
			UpdatedBy: &user.UpdatedBy,
		}).Where("deleted_at IS NULL AND status = TRUE AND id = ?", user.Id).
		Returning("id, username, password, role, status, refresh").
		Scan(
			ctx,
			&response.Id,
			&response.Username,
			&response.Password,
			&response.Role,
			&response.Status,
			&response.Refresh,
		)

	if err != nil {
		return entity.UpdateUserResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateColumns(ctx context.Context, request entity.UpdateUserColumnsRequest) (entity.UpdateUserResponse, error) {
	var response entity.UpdateUserResponse

	updater := r.DB.NewUpdate().Table("users")
	for key, value := range request.Fields {
		if key == "username" {
			updater.Set(key+" = ?", value)
		} else if key == "password" {
			updater.Set(key+" = ?", value)
		} else if key == "role" {
			updater.Set(key+" = ?", value)
		} else if key == "status" {
			updater.Set(key+" = ?", value)
		} else if key == "refresh" {
			updater.Set(key+" = ?", value)
		}
	}

	updater.Set("updated_at = NOW()")
	updater.Set("updated_by = ?", request.Fields["updated_by"])

	err := updater.Where("deleted_at IS NULL AND status = TRUE AND id = ?", request.ID).
		Returning("id, username, password, role, status, refresh").
		Scan(
			ctx,
			&response.Id,
			&response.Username,
			&response.Password,
			&response.Role,
			&response.Status,
			&response.Refresh,
		)

	if err != nil {
		return entity.UpdateUserResponse{}, err
	}

	return response, nil
}

func (r *Repo) Delete(ctx context.Context, userID, deletedBy int) (entity.DeleteUserResponse, error) {
	result, err := r.DB.NewUpdate().
		Set("deleted_at = NOW()").
		Set("deleted_by = ?", deletedBy).
		Table("users").
		Where("deleted_at IS NULL AND id = ?", userID).
		Exec(ctx)

	if err != nil {
		return entity.DeleteUserResponse{}, err
	}

	rowEffects, err := result.RowsAffected()
	if err != nil {
		return entity.DeleteUserResponse{}, err
	}

	if rowEffects == 0 {
		return entity.DeleteUserResponse{}, errors.New("user not found")
	}

	return entity.DeleteUserResponse{
		Message: "success",
	}, nil
}
