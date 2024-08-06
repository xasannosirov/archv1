package chat

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type RepoChat struct {
	DB *postgres.DB
}

func NewChatRepo(db *postgres.DB) ChatRepository {
	return &RepoChat{
		DB: db,
	}
}

func (ch *RepoChat) UserGroups(ctx context.Context, userID int64) ([]entity.GetGroupResponse, error) {
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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			_ = err
		}
	}(rows)

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

func (ch *RepoChat) GroupUsers(ctx context.Context, groupID int64) ([]entity.GetUserResponse, error) {
	query := fmt.Sprintf(`
	SELECT
		u.id,
		u.username,
		u.role,
		u.status
	FROM
	    group_users AS gu
	INNER JOIN
		groups AS g ON g.id = gu.group_id
	INNER JOIN 
	    users u ON u.id = gu.user_id
	WHERE
	    g.deleted_at IS NULL AND gu.deleted_at IS NULL AND u.deleted_at IS NULL
	AND gu.group_id = '%d' AND u.status = TRUE
	`, groupID)

	var response []entity.GetUserResponse

	rows, err := ch.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			_ = err
		}
	}(rows)

	for rows.Next() {
		var (
			user entity.GetUserResponse
		)
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Role,
			&user.Status,
		)

		if err != nil {
			return nil, err
		}

		response = append(response, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (ch *RepoChat) GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error) {
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

func (ch *RepoChat) CreateGroup(ctx context.Context, group entity.CreateGroupRequest) (entity.CreateGroupResponse, error) {
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

func (ch *RepoChat) UpdateGroup(ctx context.Context, group entity.UpdateGroupRequest) (entity.UpdateGroupResponse, error) {
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

func (ch *RepoChat) UpdateGroupColumns(ctx context.Context, fields entity.UpdateGroupColumns) (entity.UpdateGroupResponse, error) {
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

func (ch *RepoChat) DeleteGroup(ctx context.Context, groupID, deletedBy int64) (entity.DeleteGroupResponse, error) {
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

func (ch *RepoChat) AddUserToGroup(ctx context.Context, userID, groupID int64) error {
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

func (ch *RepoChat) RemoveUserFromGroup(ctx context.Context, userID, groupID int64) error {
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

func (ch *RepoChat) addAdmin(ctx context.Context, adminID, groupID int64) error {
	query := fmt.Sprintf(`INSERT INTO group_users (group_id, user_id) VALUES ('%d', '%d')`, groupID, adminID)

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

func (ch *RepoChat) CreateChat(ctx context.Context, receiverID, creator int64, chatType string) (entity.CreatedChatResponse, error) {
	query := fmt.Sprintf(`INSERT INTO chat (receiver_id, chat_type) VALUES ('%d', '%s') RETURNING id, receiver_id, chat_type`, receiverID, chatType)

	var response entity.CreatedChatResponse
	err := ch.DB.QueryRowContext(ctx, query).Scan(
		&response.ChatId,
		&response.ReceiverID,
		&response.ChatType,
	)
	if err != nil {
		return entity.CreatedChatResponse{}, err
	}

	if chatType == "group" {
		err := ch.addAdmin(ctx, creator, int64(response.ChatId))
		if err != nil {
			return entity.CreatedChatResponse{}, err
		}
	}

	return response, nil
}

func (ch *RepoChat) DeleteChat(ctx context.Context, chatID int64) error {
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

func (ch *RepoChat) UserChats(ctx context.Context, userID int64) (entity.UserChatsResponse, error) {
	query := `SELECT id, chat_type, receiver_id FROM chat WHERE receiver_id = $1 AND deleted_at IS NULL`

	var response entity.UserChatsResponse

	rows, err := ch.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return entity.UserChatsResponse{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			_ = err
		}
	}(rows)

	for rows.Next() {
		var chat struct {
			ChatID     int    `json:"chat_id"`
			ChatType   string `json:"chat_type"`
			ReceiverID int    `json:"receiver_id"`
		}

		err = rows.Scan(
			&chat.ChatID,
			&chat.ChatType,
			&chat.ReceiverID,
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

func (ch *RepoChat) SendMessage(ctx context.Context, message entity.SendMessageRequest) error {
	query := fmt.Sprintf(`
	INSERT INTO messages (chat_id, content, message_type, sender) VALUES ('%d', '%s', '%s', '%d')`,
		message.ChatID,
		message.Message,
		message.MessageType,
		message.Sender,
	)

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

func (ch *RepoChat) UpdateMessage(ctx context.Context, message entity.UpdateMessageRequest) error {
	query := fmt.Sprintf(`
	UPDATE messages SET  chat_id = '%d', content = '%s' WHERE id = '%d' AND deleted_at IS NULL`,
		message.ChatID,
		message.NewMessage,
		message.MessageID,
	)

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

func (ch *RepoChat) DeleteMessage(ctx context.Context, messageID int64) error {
	query := fmt.Sprintf(`UPDATE messages SET  deleted_at = NOW() WHERE id = '%d' AND deleted_at IS NULL`, messageID)

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

func (ch *RepoChat) GetChatMessages(ctx context.Context, chatID int64) (entity.ChatMessagesResponse, error) {
	query := fmt.Sprintf(
		`SELECT id, chat_id, content, sender, message_type FROM messages WHERE chat_id = '%d' AND deleted_at IS NULL;`,
		chatID)

	rows, err := ch.DB.QueryContext(ctx, query)
	if err != nil {
		return entity.ChatMessagesResponse{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			_ = err
		}
	}(rows)

	var response entity.ChatMessagesResponse
	for rows.Next() {
		var message struct {
			MessageID   int    `json:"message_id"`
			ChatID      int    `json:"chat_id"`
			Sender      int    `json:"sender"`
			Message     string `json:"content"`
			MessageType string `json:"message_type"`
		}

		err = rows.Scan(
			&message.MessageID,
			&message.ChatID,
			&message.Message,
			&message.Sender,
			&message.MessageType,
		)
		if err != nil {
			return entity.ChatMessagesResponse{}, err
		}

		response.Messages = append(response.Messages, message)
	}

	if err := rows.Err(); err != nil {
		return entity.ChatMessagesResponse{}, err
	}

	return response, nil
}

func (ch *RepoChat) GetChat(ctx context.Context, chatID int64) (entity.Chat, error) {
	query := fmt.Sprintf(`
	SELECT id, chat_type, receiver_id FROM chat	WHERE id = '%d' AND deleted_at IS NULL
	`, chatID)

	var response entity.Chat

	row := ch.DB.QueryRowContext(ctx, query)

	err := row.Scan(&response.ID, &response.ChatType, &response.ReceiverID)
	if err != nil {
		return entity.Chat{}, err
	}

	return response, nil
}

func (ch *RepoChat) GetMessage(ctx context.Context, messageID int64) (entity.Message, error) {
	query := fmt.Sprintf(`
	SELECT id, chat_id, content, message_type, sender FROM messages	WHERE id = '%d' AND deleted_at IS NULL
	`, messageID)

	var response entity.Message

	row := ch.DB.QueryRowContext(ctx, query)

	err := row.Scan(&response.ID, &response.ChatId, &response.Content, &response.MessageType, &response.Sender)
	if err != nil {
		return entity.Message{}, err
	}

	return response, nil
}
