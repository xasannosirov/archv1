package chat

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"database/sql"
	"time"
)

type ChatRepo struct {
	DB *postgres.DB
}

func NewChatRepo(db *postgres.DB) ChatRepository {
	return &ChatRepo{
		DB: db,
	}
}

func (ch *ChatRepo) UserGroups(ctx context.Context, userID int64) ([]entity.GetGroupResponse, error) {
	query := `
	SELECT
		g.id,
		g.name,
		g.username,
		g.description
	FROM
	    groups AS g
	INNER JOIN
		group_users AS gu ON g.id = gu.group_id
	WHERE
	    g.deleted_at IS NULL AND gu.deleted_at IS NULL
	AND gu.user_id = $1
	`

	var response []entity.GetGroupResponse

	rows, err := ch.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			nullDescription sql.NullString
			group           entity.GetGroupResponse
		)
		err := rows.Scan(
			&group.GroupId,
			&group.Name,
			&group.Username,
			&nullDescription,
		)

		if err != nil {
			return nil, err
		}

		response = append(response, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (ch *ChatRepo) GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error) {
	query := `
	SELECT
		id,
		name,
		username,
		description
	FROM
	    groups
	WHERE
		deleted_at IS NULL AND id = $1;
	`

	var (
		nullDescription sql.NullString
		result          entity.GetGroupResponse
	)
	err := ch.DB.QueryRowContext(ctx, query, groupID).Scan(
		&result.GroupId,
		&result.Name,
		&result.Username,
		&nullDescription,
	)
	if err != nil {
		return entity.GetGroupResponse{}, err
	}

	if nullDescription.Valid {
		result.Description = nullDescription.String
	}

	return result, nil
}

func (ch *ChatRepo) CreateGroup(ctx context.Context, group entity.CreateGroupRequest) (entity.CreateGroupResponse, error) {
	var (
		nullDescription sql.NullString
		response        entity.CreateGroupResponse
	)

	err := ch.DB.NewInsert().
		Model(
			&entity.Groups{
				Name:        group.Name,
				Username:    group.Username,
				Description: &group.Description,
				CreatedBy:   group.CreatedBy,
			}).
		Returning("id, name, username, description").
		Scan(ctx, &response.GroupID, &response.Name, &response.Username, &nullDescription)

	if err != nil {
		return entity.CreateGroupResponse{}, err
	}

	if nullDescription.Valid {
		response.Description = nullDescription.String
	}

	return response, nil
}

func (ch *ChatRepo) UpdateGroup(ctx context.Context, group entity.UpdateGroupRequest) (entity.UpdateGroupResponse, error) {
	var (
		nullDescription sql.NullString
		response        entity.UpdateGroupResponse
	)

	err := ch.DB.NewUpdate().
		Set("name = ?", group.Name).
		Set("description = ?", group.Description).
		Set("username = ?", group.Username).
		Set("updated_at = ?", time.Now()).
		Set("updated_by = ?", group.UpdatedBy).
		Where("deleted_at IS NULL AND id = ?", group.GroupID).
		Returning("id, name, username, description").
		Scan(ctx, &response.GroupID, &response.Name, &response.Username, &nullDescription)

	if err != nil {
		return entity.UpdateGroupResponse{}, err
	}

	if nullDescription.Valid {
		response.Description = nullDescription.String
	}

	return response, nil
}

func (ch *ChatRepo) UpdateGroupColumns(ctx context.Context, fields entity.UpdateGroupColumns) (entity.UpdateGroupResponse, error) {
	var (
		nullDescription sql.NullString
		response        entity.UpdateGroupResponse
	)

	updater := ch.DB.NewUpdate()

	for key, value := range fields.Fields {
		if key == "name" {
			updater.Set(key+" = ?", value)
		} else if key == "username" {
			updater.Set(key+" = ?", value)
		} else if key == "description" {
			updater.Set(key+" = ?", value)
		} else if key == "updated_by" {
			updater.Set(key+" = ?", value)
		}
	}

	updater.Set("deleted_at = ?", time.Now())
	err := updater.Where("deleted_at IS NULL AND id = ?", fields.GroupID).
		Returning("id, name, username, description").
		Scan(ctx, &response.GroupID, &response.Name, &response.Username, &nullDescription)

	if err != nil {
		return entity.UpdateGroupResponse{}, err
	}

	if nullDescription.Valid {
		response.Description = nullDescription.String
	}

	return response, nil
}

func (ch *ChatRepo) DeleteGroup(ctx context.Context, groupID, deletedBy int64) (entity.DeleteGroupResponse, error) {
	res, err := ch.DB.NewUpdate().
		Set("deleted_at = ?", time.Now()).
		Set("deleted_by = ?", deletedBy).
		Where("deleted_at IS NULL AND id = ?", groupID).
		Exec(ctx)

	if err != nil {
		return entity.DeleteGroupResponse{}, err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return entity.DeleteGroupResponse{}, err
	}

	if rows != 0 {
		return entity.DeleteGroupResponse{}, sql.ErrNoRows
	}

	return entity.DeleteGroupResponse{
		Message: "success",
	}, nil
}

func (ch *ChatRepo) AddUserToGroup(ctx context.Context, userID, groupID, createdBy int64) error {
	return nil
}

func (ch *ChatRepo) RemoveUserFromGroup(ctx context.Context, userID, groupID, deletedBy int64) error {
	return nil
}
