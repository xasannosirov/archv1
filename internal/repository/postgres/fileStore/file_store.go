package fileStore

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
)

type Repo struct {
	DB *postgres.DB
}

func NewFileStoreRepo(DB *postgres.DB) FilesStoreRepository {
	return &Repo{
		DB: DB,
	}
}

func (r *Repo) ListFolder(ctx context.Context, filter entity.Filter, lang string) (entity.ListFolderResponse, error) {
	return entity.ListFolderResponse{}, nil
}

func (r *Repo) GetFolder(ctx context.Context, folderID int, lang string) (entity.GetFolderResponse, error) {
	return entity.GetFolderResponse{}, nil
}

func (r *Repo) CreateFolder(ctx context.Context, folder entity.CreateFolderRequest) (entity.CreateFolderResponse, error) {
	return entity.CreateFolderResponse{}, nil
}

func (r *Repo) UpdateFolder(ctx context.Context, folder entity.UpdateFolderRequest) (entity.UpdateFolderResponse, error) {
	return entity.UpdateFolderResponse{}, nil
}

func (r *Repo) UpdateFolderColumns(ctx context.Context, fields entity.UpdateFolderColumnsRequest) (entity.UpdateFolderResponse, error) {
	return entity.UpdateFolderResponse{}, nil
}

func (r *Repo) DeleteFolder(ctx context.Context, folderID, deletedBy int) (entity.DeleteFolderResponse, error) {
	return entity.DeleteFolderResponse{}, nil
}

func (r *Repo) ListFile(ctx context.Context, filter entity.Filter, lang string) (entity.ListFileResponse, error) {
	return entity.ListFileResponse{}, nil
}

func (r *Repo) GetFile(ctx context.Context, fileID int, lang string) (entity.GetFileResponse, error) {
	return entity.GetFileResponse{}, nil
}

func (r *Repo) CreateFile(ctx context.Context, file entity.CreateFileRequest) (entity.CreateFileResponse, error) {
	return entity.CreateFileResponse{}, nil
}

func (r *Repo) UpdateFile(ctx context.Context, file entity.UpdateFileRequest) (entity.UpdateFileResponse, error) {
	return entity.UpdateFileResponse{}, nil
}

func (r *Repo) UpdateFileColumns(ctx context.Context, fields entity.UpdateFileColumnsRequest) (entity.UpdateFileResponse, error) {
	return entity.UpdateFileResponse{}, nil
}

func (r *Repo) DeleteFile(ctx context.Context, fileID, deletedBy int) (entity.DeleteFileResponse, error) {
	return entity.DeleteFileResponse{}, nil
}
