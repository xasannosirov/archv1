package post

import (
	"archv1/internal/entity"
	"archv1/internal/repository/postgres/post"
	"context"
)

type PostService struct {
	postRepo post.PostRepository
}

func NewPostService(menuRepo post.PostRepository) PostServiceI {
	return &PostService{
		postRepo: menuRepo,
	}
}

func (r *PostService) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListPostResponse, error) {
	return r.postRepo.List(ctx, filter, lang)
}

func (r *PostService) GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error) {
	return r.postRepo.GetByID(ctx, postID, lang)
}

func (r *PostService) Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error) {
	return r.postRepo.Create(ctx, post)
}

func (r *PostService) Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error) {
	return r.postRepo.Update(ctx, post)
}

func (r *PostService) UpdateColumns(ctx context.Context, post entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	return r.postRepo.UpdateColumns(ctx, post)
}

func (r *PostService) Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error) {
	return r.postRepo.Delete(ctx, postID, deletedBy)
}

func (r *PostService) AddFile(ctx context.Context, fileURL string, postID, updatedBy int) error {
	return r.postRepo.AddFile(ctx, fileURL, postID, updatedBy)
}
