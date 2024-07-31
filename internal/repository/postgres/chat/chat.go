package chat

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"database/sql"
	"fmt"
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
	query := fmt.Sprintf(`
	SELECT
		g.id,
		g.name,
		g.username,
		g.description
	FROM
	    group_users AS gu
	INNER JOIN
		groups AS g ON g.id = gu.group_id
	INNER JOIN 
	    users u ON u.id = gu.user_id
	WHERE
	    g.deleted_at IS NULL AND gu.deleted_at IS NULL
	AND gu.user_id = '%d'
	`, userID)

	var response []entity.GetGroupResponse

	rows, err := ch.DB.QueryContext(ctx, query)
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

		if nullDescription.Valid {
			group.Description = nullDescription.String
		}

		response = append(response, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (ch *ChatRepo) GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error) {
	query := fmt.Sprintf(`
	SELECT
		id,
		name,
		username,
		description
	FROM
	    groups
	WHERE
		deleted_at IS NULL AND id = '%d'
	`, groupID)

	var (
		nullDescription sql.NullString
		result          entity.GetGroupResponse
	)
	err := ch.DB.QueryRowContext(ctx, query).Scan(
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

	err := ch.DB.NewUpdate().Table("groups").
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

	updater := ch.DB.NewUpdate().Table("groups")

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
	res, err := ch.DB.NewUpdate().Table("groups").
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

	if rows == 0 {
		return entity.DeleteGroupResponse{}, sql.ErrNoRows
	}

	return entity.DeleteGroupResponse{
		Message: "success",
	}, nil
}

func (ch *ChatRepo) AddUserToGroup(ctx context.Context, userID, groupID int64) error {
	checkQuery := fmt.Sprintf(`SELECT COUNT(*) FROM group_users WHERE group_id = '%d' AND user_id = '%d'`, groupID, userID)

	var count int
	if err := ch.DB.QueryRowContext(ctx, checkQuery).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		query := fmt.Sprintf(`INSERT INTO group_users (group_id, user_id) VALUES ('%d', '%d')`, groupID, userID)

		result, err := ch.DB.ExecContext(ctx, query)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return sql.ErrNoRows
		}

		return nil
	} else {
		query := fmt.Sprintf(`UPDATE group_users SET deleted_at = NULL WHERE group_id = '%d' AND user_id = '%d'`, groupID, userID)

		result, err := ch.DB.ExecContext(ctx, query)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return sql.ErrNoRows
		}

		return nil
	}

}

func (ch *ChatRepo) RemoveUserFromGroup(ctx context.Context, userID, groupID int64) error {
	query := fmt.Sprintf(`UPDATE group_users  SET deleted_at = NOW() WHERE group_id = '%d' AND user_id = '%d' AND deleted_at IS NULL`, groupID, userID)

	result, err := ch.DB.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ch *ChatRepo) CreateChat(ctx context.Context, creator int64, chatType string) (entity.CreatedChatResponse, error) {
	query := `INSERT INTO chat (creator, type) VALUES ($1, $2) RETURNING id, creator, type`

	var response entity.CreatedChatResponse
	err := ch.DB.QueryRowContext(ctx, query).Scan(
		&response.ChatId,
		&response.Creator,
		&response.ChatType,
	)
	if err != nil {
		return entity.CreatedChatResponse{}, err
	}

	return response, nil
}

func (ch *ChatRepo) DeleteChat(ctx context.Context, chatID int64) error {
	query := `UPDATE chat SET deleted_at = NOW() WHERE id = $1`

	result, err := ch.DB.ExecContext(ctx, query, chatID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (ch *ChatRepo) UserChats(ctx context.Context, userID int64) (entity.UserChatsResponse, error) {
	query := `SELECT id, type, creator FROM chat WHERE creator = $1 AND deleted_at IS NULL`

	var response entity.UserChatsResponse

	rows, err := ch.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return entity.UserChatsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat struct {
			ChatID   int    `json:"chat_id"`
			ChatType string `json:"chat_type"`
			Creator  int    `json:"creator"`
		}

		err = rows.Scan(
			&chat.ChatID,
			&chat.ChatType,
			&chat.Creator,
		)
		if err != nil {
			return entity.UserChatsResponse{}, err
		}

		response.Chats = append(response.Chats, chat)
	}

	if err := rows.Err(); err != nil {
		return entity.UserChatsResponse{}, err
	}

	return response, nil
}
