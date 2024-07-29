package chat

import (
	"archv1/internal/entity"
	"context"
)

type ChatUseCaseI interface {
	UserGroups(ctx context.Context, userID int64) ([]entity.GetGroupResponse, error)
	GetGroup(ctx context.Context, groupID int64) (entity.GetGroupResponse, error)
	CreateGroup(ctx context.Context, group entity.CreateGroupRequest) (entity.CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, group entity.UpdateGroupRequest) (entity.UpdateGroupResponse, error)
	UpdateGroupColumns(ctx context.Context, fields entity.UpdateGroupColumns) (entity.UpdateGroupResponse, error)
	DeleteGroup(ctx context.Context, groupID, deletedBy int64) (entity.DeleteGroupResponse, error)
	AddUserToGroup(ctx context.Context, userID, groupID, createdBy int64) error
	RemoveUserFromGroup(ctx context.Context, userID, groupID, deletedBy int64) error
}
