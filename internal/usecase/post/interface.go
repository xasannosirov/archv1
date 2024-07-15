package post

import (
	"archv1/internal/entity"
	"context"
)

type PostUseCaseI interface {
	List(ctx context.Context, filter entity.Filter, lang string) (entity.ListPostResponse, error)
	GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error)
	Create(ctx context.Context, menu entity.CreatePostRequest) (entity.CreatePostResponse, error)
	Update(ctx context.Context, menu entity.UpdatePostRequest) (entity.UpdatePostResponse, error)
	UpdateColumns(ctx context.Context, fields entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error)
	Delete(ctx context.Context, menuID, deletedBy int) (entity.DeletePostResponse, error)
	AddFile(ctx context.Context, fileURL string, postID, updatedBy int) error
}
