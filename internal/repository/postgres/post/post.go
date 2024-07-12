package post

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
)

type Repo struct {
	DB *postgres.DB
}

func NewPostRepo(DB *postgres.DB) PostRepository {
	return &Repo{
		DB: DB,
	}
}

func (r *Repo) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListPostResponse, error) {
	return entity.ListPostResponse{}, nil
}

func (r *Repo) GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error) {
	return entity.GetPostResponse{}, nil
}

func (r *Repo) Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error) {
	return entity.CreatePostResponse{}, nil
}

func (r *Repo) Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error) {
	return entity.UpdatePostResponse{}, nil
}

func (r *Repo) UpdateColumns(ctx context.Context, post entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	return entity.UpdatePostResponse{}, nil
}

func (r *Repo) Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error) {
	return entity.DeletePostResponse{}, nil
}
