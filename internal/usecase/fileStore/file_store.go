package fileStore

import (
	"archv1/internal/entity"
	"archv1/internal/service/fileStore"
	"context"
)

type FilesStoreUseCase struct {
	fileStoreService fileStore.FilesStoreServiceI
}

func NewFilesStoreUseCase(service fileStore.FilesStoreServiceI) FilesStoreUseCaseI {
	return &FilesStoreUseCase{
		fileStoreService: service,
	}
}

func (f *FilesStoreUseCase) ListFolder(ctx context.Context, filter entity.Filter, lang string) (entity.ListFolderResponse, error) {
	return f.fileStoreService.ListFolder(ctx, filter, lang)
}

func (f *FilesStoreUseCase) GetFolder(ctx context.Context, folderID int, lang string) (entity.GetFolderResponse, error) {
	return f.fileStoreService.GetFolder(ctx, folderID, lang)
}

func (f *FilesStoreUseCase) CreateFolder(ctx context.Context, folder entity.CreateFolderRequest) (entity.CreateFolderResponse, error) {
	return f.fileStoreService.CreateFolder(ctx, folder)
}

func (f *FilesStoreUseCase) UpdateFolder(ctx context.Context, folder entity.UpdateFolderRequest) (entity.UpdateFolderResponse, error) {
	return f.fileStoreService.UpdateFolder(ctx, folder)
}

func (f *FilesStoreUseCase) UpdateFolderColumns(ctx context.Context, fields entity.UpdateFolderColumnsRequest) (entity.UpdateFolderResponse, error) {
	return f.fileStoreService.UpdateFolderColumns(ctx, fields)
}

func (f *FilesStoreUseCase) DeleteFolder(ctx context.Context, folderID, deletedBy int) (entity.DeleteFolderResponse, error) {
	return f.fileStoreService.DeleteFolder(ctx, folderID, deletedBy)
}

func (f *FilesStoreUseCase) ListFile(ctx context.Context, filter entity.Filter, lang string) (entity.ListFileResponse, error) {
	return f.fileStoreService.ListFile(ctx, filter, lang)
}

func (f *FilesStoreUseCase) GetFile(ctx context.Context, fileID int, lang string) (entity.GetFileResponse, error) {
	return f.fileStoreService.GetFile(ctx, fileID, lang)
}

func (f *FilesStoreUseCase) CreateFile(ctx context.Context, file entity.CreateFileRequest) (entity.CreateFileResponse, error) {
	return f.fileStoreService.CreateFile(ctx, file)
}

func (f *FilesStoreUseCase) UpdateFile(ctx context.Context, file entity.UpdateFileRequest) (entity.UpdateFileResponse, error) {
	return f.fileStoreService.UpdateFile(ctx, file)
}

func (f *FilesStoreUseCase) UpdateFileColumns(ctx context.Context, fields entity.UpdateFileColumnsRequest) (entity.UpdateFileResponse, error) {
	return f.fileStoreService.UpdateFileColumns(ctx, fields)
}

func (f *FilesStoreUseCase) DeleteFile(ctx context.Context, fileID, deletedBy int) (entity.DeleteFileResponse, error) {
	return f.fileStoreService.DeleteFile(ctx, fileID, deletedBy)
}
