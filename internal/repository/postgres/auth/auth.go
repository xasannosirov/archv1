package auth

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"errors"
	"fmt"
)

type Repo struct {
	DB *postgres.DB
}

func NewAuthRepo(DB *postgres.DB) AuthRepository {
	return &Repo{
		DB: DB,
	}
}

func (r *Repo) UniqueUsername(ctx context.Context, username string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE username = ? AND deleted_at IS NULL;`

	var count int64
	if err := r.DB.QueryRowContext(ctx, query, username).Scan(&count); err != nil {
		return true, err
	}

	return count != 0, nil
}

func (r *Repo) UpdateToken(ctx context.Context, id int, token string) error {
	result, err := r.DB.NewUpdate().
		Table("users").
		Set("refresh = ?", token).
		Where("deleted_at IS NULL AND id = ?", id).
		Exec(ctx)

	if err != nil {
		fmt.Println(err.Error(), 111)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (entity.GetUserResponse, error) {
	var response entity.GetUserResponse

	selectQuery := fmt.Sprintf(`
	SELECT 
	    id, 
	    username, 
	    password, 
	    role, 
	    status, 
	    refresh 
	FROM users 
	WHERE username = ? AND deleted_at IS NULL AND status = TRUE`)

	err := r.DB.QueryRowContext(ctx, selectQuery, username).Scan(
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

func (r *Repo) GetUserByToken(ctx context.Context, token string) (entity.GetUserResponse, error) {
	var response entity.GetUserResponse

	selectQuery := fmt.Sprintf(`
	SELECT 
	    id, 
	    username, 
	    password, 
	    role, 
	    status, 
	    refresh 
	FROM users 
	WHERE refresh = ? AND deleted_at IS NULL AND status = TRUE`)

	err := r.DB.QueryRowContext(ctx, selectQuery, token).Scan(
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
