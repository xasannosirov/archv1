package chat

import (
	"archv1/internal/entity"
	"context"
)

type ChatUseCaseI interface {
	UserGroups(ctx context.Context, userID int64) ([]entity.GetGroupResponse, error)
	GroupUsers(ctx context.Context, groupID int64) ([]entity.GetUserResponse, error)
	GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error)
	CreateGroup(ctx context.Context, group entity.CreateGroupRequest) (entity.CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, group entity.UpdateGroupRequest) (entity.UpdateGroupResponse, error)
	UpdateGroupColumns(ctx context.Context, fields entity.UpdateGroupColumns) (entity.UpdateGroupResponse, error)
	DeleteGroup(ctx context.Context, groupID, deletedBy int64) (entity.DeleteGroupResponse, error)
	AddUserToGroup(ctx context.Context, userID, groupID int64) error
	RemoveUserFromGroup(ctx context.Context, userID, groupID int64) error
	CreateChat(ctx context.Context, receiverID, creator int64, chatType string) (entity.CreatedChatResponse, error)
	DeleteChat(ctx context.Context, chatID int64) error
	UserChats(ctx context.Context, userID int64) (entity.UserChatsResponse, error)
	SendMessage(ctx context.Context, message entity.SendMessageRequest) error
	UpdateMessage(ctx context.Context, message entity.UpdateMessageRequest) error
	DeleteMessage(ctx context.Context, messageID int64) error
	GetChatMessages(ctx context.Context, chatID int64) (entity.ChatMessagesResponse, error)
	GetChat(ctx context.Context, chatID int64) (entity.Chat, error)
	GetMessage(ctx context.Context, messageID int64) (entity.Message, error)
}
