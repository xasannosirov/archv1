package post

import (
	"archv1/internal/entity"
	"archv1/internal/service/post"
	"context"
)

type PostUseCase struct {
	postService post.PostServiceI
}

func NewPostUseCase(service post.PostServiceI) PostUseCaseI {
	return &PostUseCase{
		postService: service,
	}
}

func (r *PostUseCase) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListPostResponse, error) {
	return r.postService.List(ctx, filter, lang)
}

func (r *PostUseCase) GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error) {
	return r.postService.GetByID(ctx, postID, lang)
}

func (r *PostUseCase) Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error) {
	return r.postService.Create(ctx, post)
}

func (r *PostUseCase) Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error) {
	return r.postService.Update(ctx, post)
}

func (r *PostUseCase) UpdateColumns(ctx context.Context, post entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	return r.postService.UpdateColumns(ctx, post)
}

func (r *PostUseCase) Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error) {
	return r.postService.Delete(ctx, postID, deletedBy)
}

func (r *PostUseCase) AddFile(ctx context.Context, fileURL string, postID int) error {
	return r.postService.AddFile(ctx, fileURL, postID)
}
