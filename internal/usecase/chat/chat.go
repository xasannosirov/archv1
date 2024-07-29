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

func (ch *ChatUseCase) AddUserToGroup(ctx context.Context, userID, groupID, createdBy int64) error {
	return ch.chatService.AddUserToGroup(ctx, userID, groupID, createdBy)
}

func (ch *ChatUseCase) RemoveUserFromGroup(ctx context.Context, userID, groupID, deletedBy int64) error {
	return ch.chatService.RemoveUserFromGroup(ctx, userID, groupID, deletedBy)
}
