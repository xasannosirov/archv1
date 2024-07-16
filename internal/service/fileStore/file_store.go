package fileStore

import (
	"archv1/internal/entity"
	"archv1/internal/repository/postgres/fileStore"
	"context"
)

type FilesStoreService struct {
	fileStoreRepo fileStore.FilesStoreRepository
}

func NewFilesStoreService(fileStoreRepo fileStore.FilesStoreRepository) FilesStoreServiceI {
	return &FilesStoreService{
		fileStoreRepo: fileStoreRepo,
	}
}

func (f *FilesStoreService) ListFolder(ctx context.Context, filter entity.Filter) (entity.ListFolderResponse, error) {
	return f.fileStoreRepo.ListFolder(ctx, filter)
}

func (f *FilesStoreService) GetFolder(ctx context.Context, folderID int) (entity.GetFolderResponse, error) {
	return f.fileStoreRepo.GetFolder(ctx, folderID)
}

func (f *FilesStoreService) CreateFolder(ctx context.Context, folder entity.CreateFolderRequest) (entity.CreateFolderResponse, error) {
	return f.fileStoreRepo.CreateFolder(ctx, folder)
}

func (f *FilesStoreService) UpdateFolder(ctx context.Context, folder entity.UpdateFolderRequest) (entity.UpdateFolderResponse, error) {
	return f.fileStoreRepo.UpdateFolder(ctx, folder)
}

func (f *FilesStoreService) UpdateFolderColumns(ctx context.Context, fields entity.UpdateFolderColumnsRequest) (entity.UpdateFolderResponse, error) {
	return f.fileStoreRepo.UpdateFolderColumns(ctx, fields)
}

func (f *FilesStoreService) DeleteFolder(ctx context.Context, folderID, deletedBy int) (entity.DeleteFolderResponse, error) {
	return f.fileStoreRepo.DeleteFolder(ctx, folderID, deletedBy)
}

func (f *FilesStoreService) ListFile(ctx context.Context, filter entity.Filter) (entity.ListFileResponse, error) {
	return f.fileStoreRepo.ListFile(ctx, filter)
}

func (f *FilesStoreService) GetFile(ctx context.Context, fileID int) (entity.GetFileResponse, error) {
	return f.fileStoreRepo.GetFile(ctx, fileID)
}

func (f *FilesStoreService) CreateFile(ctx context.Context, file entity.CreateFileRequest) (entity.CreateFileResponse, error) {
	return f.fileStoreRepo.CreateFile(ctx, file)
}

func (f *FilesStoreService) UpdateFile(ctx context.Context, file entity.UpdateFileRequest) (entity.UpdateFileResponse, error) {
	return f.fileStoreRepo.UpdateFile(ctx, file)
}

func (f *FilesStoreService) UpdateFileColumns(ctx context.Context, fields entity.UpdateFileColumnsRequest) (entity.UpdateFileResponse, error) {
	return f.fileStoreRepo.UpdateFileColumns(ctx, fields)
}

func (f *FilesStoreService) DeleteFile(ctx context.Context, fileID, deletedBy int) (entity.DeleteFileResponse, error) {
	return f.fileStoreRepo.DeleteFile(ctx, fileID, deletedBy)
}
