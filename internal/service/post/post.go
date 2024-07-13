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
	return entity.ListPostResponse{}, nil
}

func (r *PostService) GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error) {
	return entity.GetPostResponse{}, nil
}

func (r *PostService) Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error) {
	return entity.CreatePostResponse{}, nil
}

func (r *PostService) Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error) {
	return entity.UpdatePostResponse{}, nil
}

func (r *PostService) UpdateColumns(ctx context.Context, post entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	return entity.UpdatePostResponse{}, nil
}

func (r *PostService) Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error) {
	return entity.DeletePostResponse{}, nil
}

func (r *PostService) AddFile(ctx context.Context, fileURL string, postID int) error {
	return r.postRepo.AddFile(ctx, fileURL, postID)
}
