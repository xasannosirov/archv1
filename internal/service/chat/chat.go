package chat

import (
	"archv1/internal/entity"
	"archv1/internal/repository/postgres/chat"
	"context"
)

type ChatService struct {
	chatRepo chat.ChatRepository
}

func NewChatService(chatRepo chat.ChatRepository) ChatServiceI {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

func (ch *ChatService) UserGroups(ctx context.Context, userID int64) ([]entity.GetGroupResponse, error) {
	return ch.chatRepo.UserGroups(ctx, userID)
}

func (ch *ChatService) GroupUsers(ctx context.Context, groupID int64) ([]entity.GetUserResponse, error) {
	return ch.chatRepo.GroupUsers(ctx, groupID)
}

func (ch *ChatService) GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error) {
	return ch.chatRepo.GetGroup(ctx, groupID)
}

func (ch *ChatService) CreateGroup(ctx context.Context, group entity.CreateGroupRequest) (entity.CreateGroupResponse, error) {
	return ch.chatRepo.CreateGroup(ctx, group)
}

func (ch *ChatService) UpdateGroup(ctx context.Context, group entity.UpdateGroupRequest) (entity.UpdateGroupResponse, error) {
	return ch.chatRepo.UpdateGroup(ctx, group)
}

func (ch *ChatService) UpdateGroupColumns(ctx context.Context, fields entity.UpdateGroupColumns) (entity.UpdateGroupResponse, error) {
	return ch.chatRepo.UpdateGroupColumns(ctx, fields)
}

func (ch *ChatService) DeleteGroup(ctx context.Context, groupID, deletedBy int64) (entity.DeleteGroupResponse, error) {
	return ch.chatRepo.DeleteGroup(ctx, groupID, deletedBy)
}

func (ch *ChatService) AddUserToGroup(ctx context.Context, userID, groupID int64) error {
	return ch.chatRepo.AddUserToGroup(ctx, userID, groupID)
}

func (ch *ChatService) RemoveUserFromGroup(ctx context.Context, userID, groupID int64) error {
	return ch.chatRepo.RemoveUserFromGroup(ctx, userID, groupID)
}

func (ch *ChatService) CreateChat(ctx context.Context, receiverID, creator int64, chatType string) (entity.CreatedChatResponse, error) {
	return ch.chatRepo.CreateChat(ctx, receiverID, creator, chatType)
}

func (ch *ChatService) DeleteChat(ctx context.Context, chatID int64) error {
	return ch.chatRepo.DeleteChat(ctx, chatID)
}

func (ch *ChatService) UserChats(ctx context.Context, userID int64) (entity.UserChatsResponse, error) {
	return ch.chatRepo.UserChats(ctx, userID)
}

func (ch *ChatService) SendMessage(ctx context.Context, message entity.SendMessageRequest) error {
	return ch.chatRepo.SendMessage(ctx, message)
}

func (ch *ChatService) UpdateMessage(ctx context.Context, message entity.UpdateMessageRequest) error {
	return ch.chatRepo.UpdateMessage(ctx, message)
}

func (ch *ChatService) DeleteMessage(ctx context.Context, messageID int64) error {
	return ch.chatRepo.DeleteMessage(ctx, messageID)
}

func (ch *ChatService) GetChatMessages(ctx context.Context, chatID int64) (entity.ChatMessagesResponse, error) {
	return ch.chatRepo.GetChatMessages(ctx, chatID)
}

func (ch *ChatService) GetChat(ctx context.Context, chatID int64) (entity.Chat, error) {
	return ch.chatRepo.GetChat(ctx, chatID)
}

func (ch *ChatService) GetMessage(ctx context.Context, messageID int64) (entity.Message, error) {
	return ch.chatRepo.GetMessage(ctx, messageID)
}
