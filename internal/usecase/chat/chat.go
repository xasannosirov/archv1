package chat

import (
	"archv1/internal/entity"
	"archv1/internal/service/chat"
	"context"
)

type ChatUseCase struct {
	chatService chat.ChatServiceI
}

func NewChatUseCase(chatService chat.ChatServiceI) ChatUseCaseI {
	return &ChatUseCase{
		chatService: chatService,
	}
}

func (ch *ChatUseCase) UserGroups(ctx context.Context, userID int64) ([]entity.GetGroupResponse, error) {
	return ch.chatService.UserGroups(ctx, userID)
}

func (ch *ChatUseCase) GroupUsers(ctx context.Context, groupID int64) ([]entity.GetUserResponse, error) {
	return ch.chatService.GroupUsers(ctx, groupID)
}

func (ch *ChatUseCase) GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error) {
	return ch.chatService.GetGroup(ctx, groupID)
}

func (ch *ChatUseCase) CreateGroup(ctx context.Context, group entity.CreateGroupRequest) (entity.CreateGroupResponse, error) {
	return ch.chatService.CreateGroup(ctx, group)
}

func (ch *ChatUseCase) UpdateGroup(ctx context.Context, group entity.UpdateGroupRequest) (entity.UpdateGroupResponse, error) {
	return ch.chatService.UpdateGroup(ctx, group)
}

func (ch *ChatUseCase) UpdateGroupColumns(ctx context.Context, fields entity.UpdateGroupColumns) (entity.UpdateGroupResponse, error) {
	return ch.chatService.UpdateGroupColumns(ctx, fields)
}

func (ch *ChatUseCase) DeleteGroup(ctx context.Context, groupID, deletedBy int64) (entity.DeleteGroupResponse, error) {
	return ch.chatService.DeleteGroup(ctx, groupID, deletedBy)
}

func (ch *ChatUseCase) AddUserToGroup(ctx context.Context, userID, groupID int64) error {
	return ch.chatService.AddUserToGroup(ctx, userID, groupID)
}

func (ch *ChatUseCase) RemoveUserFromGroup(ctx context.Context, userID, groupID int64) error {
	return ch.chatService.RemoveUserFromGroup(ctx, userID, groupID)
}

func (ch *ChatUseCase) CreateChat(ctx context.Context, receiverID int64, chatType string) (entity.CreatedChatResponse, error) {
	return ch.chatService.CreateChat(ctx, receiverID, chatType)
}

func (ch *ChatUseCase) DeleteChat(ctx context.Context, chatID int64) error {
	return ch.chatService.DeleteChat(ctx, chatID)
}

func (ch *ChatUseCase) UserChats(ctx context.Context, userID int64) (entity.UserChatsResponse, error) {
	return ch.chatService.UserChats(ctx, userID)
}

func (ch *ChatUseCase) SendMessage(ctx context.Context, message entity.SendMessageRequest) error {
	return ch.chatService.SendMessage(ctx, message)
}

func (ch *ChatUseCase) UpdateMessage(ctx context.Context, message entity.UpdateMessageRequest) error {
	return ch.chatService.UpdateMessage(ctx, message)
}

func (ch *ChatUseCase) DeleteMessage(ctx context.Context, messageID int64) error {
	return ch.chatService.DeleteMessage(ctx, messageID)
}

func (ch *ChatUseCase) GetChatMessages(ctx context.Context, chatID int64) (entity.ChatMessagesResponse, error) {
	return ch.chatService.GetChatMessages(ctx, chatID)
}

func (ch *ChatUseCase) GetChat(ctx context.Context, chatID int64) (entity.Chat, error) {
	return ch.chatService.GetChat(ctx, chatID)
}

func (ch *ChatUseCase) GetMessage(ctx context.Context, messageID int64) (entity.Message, error) {
	return ch.chatService.GetMessage(ctx, messageID)
}
