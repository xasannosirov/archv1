package fileStore

import (
	"archv1/internal/entity"
	"context"
)

type FilesStoreServiceI interface {
	ListFolder(ctx context.Context, filter entity.Filter, lang string) (entity.ListFolderResponse, error)
	GetFolder(ctx context.Context, folderID int, lang string) (entity.GetFolderResponse, error)
	CreateFolder(ctx context.Context, folder entity.CreateFolderRequest) (entity.CreateFolderResponse, error)
	UpdateFolder(ctx context.Context, folder entity.UpdateFolderRequest) (entity.UpdateFolderResponse, error)
	UpdateFolderColumns(ctx context.Context, fields entity.UpdateFolderColumnsRequest) (entity.UpdateFolderResponse, error)
	DeleteFolder(ctx context.Context, folderID, deletedBy int) (entity.DeleteFolderResponse, error)
	ListFile(ctx context.Context, filter entity.Filter, lang string) (entity.ListFileResponse, error)
	GetFile(ctx context.Context, fileID int, lang string) (entity.GetFileResponse, error)
	CreateFile(ctx context.Context, file entity.CreateFileRequest) (entity.CreateFileResponse, error)
	UpdateFile(ctx context.Context, file entity.UpdateFileRequest) (entity.UpdateFileResponse, error)
	UpdateFileColumns(ctx context.Context, fields entity.UpdateFileColumnsRequest) (entity.UpdateFileResponse, error)
	DeleteFile(ctx context.Context, fileID, deletedBy int) (entity.DeleteFileResponse, error)
}
