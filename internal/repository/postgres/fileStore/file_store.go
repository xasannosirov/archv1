package fileStore

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"errors"
	"fmt"
)

type Repo struct {
	DB *postgres.DB
}

func NewFileStoreRepo(DB *postgres.DB) FilesStoreRepository {
	return &Repo{
		DB: DB,
	}
}

func (r *Repo) ListFolder(ctx context.Context, filter entity.Filter) (entity.ListFolderResponse, error) {
	offset := filter.Limit * (filter.Page - 1)
	var response entity.ListFolderResponse

	selectQuery := `
	SELECT 
		id, 
		name, 
		parent_id
	FROM
	    folders
	`

	filterQuery := fmt.Sprintf(` WHERE deleted_at IS NULL LIMIT %d OFFSET %d`, filter.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+filterQuery)
	if err != nil {
		return entity.ListFolderResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var folder entity.GetFolderResponse
		err = rows.Scan(
			&folder.ID,
			&folder.Name,
			&folder.ParentID,
		)
		if err != nil {
			return entity.ListFolderResponse{}, err
		}

		response.Folders = append(response.Folders, &folder)
	}

	if err := rows.Err(); err != nil {
		return entity.ListFolderResponse{}, err
	}

	countQuery := `SELECT COUNT(*) FROM folders WHERE deleted_at IS NULL`

	if err := r.DB.QueryRowContext(ctx, countQuery).Scan(&response.Total); err != nil {
		return entity.ListFolderResponse{}, err
	}

	return response, nil
}

func (r *Repo) GetFolder(ctx context.Context, folderID int) (entity.GetFolderResponse, error) {
	var response entity.GetFolderResponse

	selectQuery := `
	SELECT
		id,
		name,
		parent_id
	FROM
	    folders
	`

	filterQuery := fmt.Sprintf(" WHERE id = %d AND deleted_at IS NULL", folderID)

	err := r.DB.QueryRowContext(ctx, selectQuery+filterQuery).Scan(
		&response.ID,
		&response.Name,
		&response.ParentID,
	)
	if err != nil {
		return entity.GetFolderResponse{}, err
	}

	return response, nil
}

func (r *Repo) CreateFolder(ctx context.Context, folder entity.CreateFolderRequest) (entity.CreateFolderResponse, error) {
	var response entity.CreateFolderResponse

	err := r.DB.NewInsert().
		Model(&entity.Folders{
			Name:      folder.Name,
			ParentID:  folder.ParentID,
			CreatedBy: folder.CreatedBy,
		}).
		Returning("id, name, parent_id").
		Scan(ctx, &response.ID, &response.Name, &response.ParentID)

	if err != nil {
		return entity.CreateFolderResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateFolder(ctx context.Context, folder entity.UpdateFolderRequest) (entity.UpdateFolderResponse, error) {
	var response entity.UpdateFolderResponse

	err := r.DB.NewUpdate().
		Table("folders").
		Set("name = ?", folder.Name).
		Set("parent_id = ?", folder.ParentID).
		Set("updated_by = ?", folder.UpdatedBy).
		Set("updated_at = NOW()").
		Where("deleted_at IS NULL AND id = ?", folder.ID).
		Returning("id, name, parent_id").
		Scan(ctx, &response.ID, &response.Name, &response.ParentID)

	if err != nil {
		return entity.UpdateFolderResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateFolderColumns(ctx context.Context, fields entity.UpdateFolderColumnsRequest) (entity.UpdateFolderResponse, error) {
	var response entity.UpdateFolderResponse

	updater := r.DB.NewUpdate().
		Table("folders")

	for key, value := range fields.Fields {
		if key == "name" {
			updater.Set(key+" = ?", value)
		} else if key == "parent_id" {
			updater.Set(key+" = ?", value)
		} else if key == "updated_by" {
			updater.Set(key+" = ?", value)
		}
	}

	err := updater.Set("updated_at = NOW()").
		Where("deleted_at IS NULL AND id = ?", fields.FolderID).
		Returning("id, name, parent_id").
		Scan(ctx, &response.ID, &response.Name, &response.ParentID)

	if err != nil {
		return entity.UpdateFolderResponse{}, err
	}

	return response, nil
}

func (r *Repo) DeleteFolder(ctx context.Context, folderID, deletedBy int) (entity.DeleteFolderResponse, error) {
	result, err := r.DB.NewUpdate().
		Table("folders").
		Set("deleted_at = NOW()").
		Set("deleted_by = ?", deletedBy).
		Where("deleted_at IS NULL AND id = ?", folderID).
		Exec(ctx)

	if err != nil {
		return entity.DeleteFolderResponse{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.DeleteFolderResponse{}, err
	}

	if rowsAffected == 0 {
		return entity.DeleteFolderResponse{}, errors.New("no rows affected")
	}

	return entity.DeleteFolderResponse{
		Message: "success",
	}, nil
}

func (r *Repo) ListFile(ctx context.Context, filter entity.Filter) (entity.ListFileResponse, error) {
	offset := filter.Limit * (filter.Page - 1)
	var response entity.ListFileResponse

	selectQuery := `
	SELECT 
		id, 
		type, 
		link,
		folder_id
	FROM
	    files
	`

	filterQuery := fmt.Sprintf(` WHERE deleted_at IS NULL LIMIT %d OFFSET %d`, filter.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+filterQuery)
	if err != nil {
		return entity.ListFileResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var file entity.GetFileResponse
		err = rows.Scan(
			&file.ID,
			&file.Type,
			&file.Link,
			&file.FolderID,
		)
		if err != nil {
			return entity.ListFileResponse{}, err
		}

		response.Files = append(response.Files, &file)
	}

	if err := rows.Err(); err != nil {
		return entity.ListFileResponse{}, err
	}

	countQuery := `SELECT COUNT(*) FROM files WHERE deleted_at IS NULL`

	if err := r.DB.QueryRowContext(ctx, countQuery).Scan(&response.Total); err != nil {
		return entity.ListFileResponse{}, err
	}

	return response, nil
}

func (r *Repo) GetFile(ctx context.Context, fileID int) (entity.GetFileResponse, error) {
	var response entity.GetFileResponse

	selectQuery := `
	SELECT
		id,
		type,
		link,
		folder_id
	FROM
	    files
	`

	filterQuery := fmt.Sprintf(" WHERE id = %d AND deleted_at IS NULL", fileID)

	err := r.DB.QueryRowContext(ctx, selectQuery+filterQuery).Scan(
		&response.ID,
		&response.Type,
		&response.Link,
		&response.FolderID,
	)
	if err != nil {
		return entity.GetFileResponse{}, err
	}

	return response, nil
}

func (r *Repo) CreateFile(ctx context.Context, file entity.CreateFileRequest) (entity.CreateFileResponse, error) {
	var response entity.CreateFileResponse

	err := r.DB.NewInsert().
		Model(&entity.Files{
			Type:      file.Type,
			Link:      file.Link,
			FolderID:  file.FolderID,
			CreatedBy: &file.CreatedBy,
		}).
		Returning("id, type, link, folder_id").
		Scan(ctx, &response.ID, &response.Type, &response.Link, &response.FolderID)

	if err != nil {
		return entity.CreateFileResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateFile(ctx context.Context, file entity.UpdateFileRequest) (entity.UpdateFileResponse, error) {
	var response entity.UpdateFileResponse

	err := r.DB.NewUpdate().
		Table("files").
		Set("type = ?", file.Type).
		Set("link = ?", file.Link).
		Set("folder_id = ?", file.FolderID).
		Set("updated_by = ?", file.UpdatedBy).
		Set("updated_at = NOW()").
		Where("deleted_at IS NULL AND id = ?", file.ID).
		Returning("id, type, link, folder_id").
		Scan(ctx, &response.ID, &response.Type, &response.Link, &response.FolderID)

	if err != nil {
		return entity.UpdateFileResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateFileColumns(ctx context.Context, fields entity.UpdateFileColumnsRequest) (entity.UpdateFileResponse, error) {
	var response entity.UpdateFileResponse

	updater := r.DB.NewUpdate().
		Table("files")

	for key, value := range fields.Fields {
		if key == "type" {
			updater.Set(key+" = ?", value)
		} else if key == "link" {
			updater.Set(key+" = ?", value)
		} else if key == "parent_id" {
			updater.Set(key+" = ?", value)
		} else if key == "updated_by" {
			updater.Set(key+" = ?", value)
		}
	}

	err := updater.Set("updated_at = NOW()").
		Where("deleted_at IS NULL AND id = ?", fields.FileID).
		Returning("id, type, link, folder_id").
		Scan(ctx, &response.ID, &response.Type, &response.Link, &response.FolderID)

	if err != nil {
		return entity.UpdateFileResponse{}, err
	}

	return response, nil
}

func (r *Repo) DeleteFile(ctx context.Context, fileID, deletedBy int) (entity.DeleteFileResponse, error) {
	result, err := r.DB.NewUpdate().
		Table("files").
		Set("deleted_at = NOW()").
		Set("deleted_by = ?", deletedBy).
		Where("deleted_at IS NULL AND id = ?", fileID).
		Exec(ctx)

	if err != nil {
		return entity.DeleteFileResponse{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.DeleteFileResponse{}, err
	}

	if rowsAffected == 0 {
		return entity.DeleteFileResponse{}, errors.New("no rows affected")
	}

	return entity.DeleteFileResponse{
		Message: "success",
	}, nil
}
