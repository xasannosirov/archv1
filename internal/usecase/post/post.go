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

func (m *PostUseCase) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListPostResponse, error) {
	response, err := m.postService.List(ctx, filter, lang)
	if err != nil {
		return entity.ListPostResponse{}, err
	}

	return response, nil
}

func (r *PostUseCase) GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error) {
	return entity.GetPostResponse{}, nil
}

func (r *PostUseCase) Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error) {
	return entity.CreatePostResponse{}, nil
}

func (r *PostUseCase) Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error) {
	return entity.UpdatePostResponse{}, nil
}

func (r *PostUseCase) UpdateColumns(ctx context.Context, post entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	return entity.UpdatePostResponse{}, nil
}

func (r *PostUseCase) Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error) {
	return entity.DeletePostResponse{}, nil
}

func (r *PostUseCase) AddFile(ctx context.Context, fileURL string, postID int) error {
	return r.postService.AddFile(ctx, fileURL, postID)
}
