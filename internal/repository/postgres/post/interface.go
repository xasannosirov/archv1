package post

import (
	"archv1/internal/entity"
	"context"
)

type PostRepository interface {
	List(ctx context.Context, filter entity.Filter, lang string) (entity.ListPostResponse, error)
	GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error)
	Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error)
	Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error)
	UpdateColumns(ctx context.Context, post entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error)
	Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error)
}
