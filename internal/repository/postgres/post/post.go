package post

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
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
	var response entity.ListPostResponse
	offset := filter.Limit * (filter.Page - 1)

	selectQuery := fmt.Sprintf(`
	SELECT
		id,
		title ->> '%s',
		content ->> '%s',
		short_content ->> '%s',
		slug,
		status,
		user_id,
		files
	FROM
	    posts
	`, lang, lang, lang)

	whereQuery := ` WHERE deleted_at IS NULL AND status = TRUE `

	filterQuery := fmt.Sprintf(" LIMIT %d OFFSET %d", filter.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+whereQuery+filterQuery)
	if err != nil {
		return entity.ListPostResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			title        string
			content      string
			shortContent string
			post         entity.GetPostResponse
		)
		err = rows.Scan(
			&post.ID,
			&title,
			&content,
			&shortContent,
			&post.Slug,
			&post.Status,
			&post.UserID,
			&post.Files,
		)
		if err != nil {
			return entity.ListPostResponse{}, err
		}

		post.Title = map[string]string{lang: title}
		post.Content = map[string]string{lang: content}
		post.ShortContent = map[string]string{lang: shortContent}

		response.Posts = append(response.Posts, &post)
	}
	if err := rows.Err(); err != nil {
		return entity.ListPostResponse{}, err
	}

	totalQuery := `SELECT count(*) FROM posts WHERE deleted_at IS NULL AND status = TRUE`
	if err := r.DB.QueryRowContext(ctx, totalQuery).Scan(&response.Total); err != nil {
		return entity.ListPostResponse{}, err
	}

	return response, nil
}

func (r *Repo) GetByID(ctx context.Context, postID int, lang string) (entity.GetPostResponse, error) {
	var (
		title        string
		content      string
		shortContent string
		response     entity.GetPostResponse
	)

	selectQuery := fmt.Sprintf(`
	SELECT
		id,
		title ->> '%s',
		content ->> '%s',
		short_content ->> '%s',
		slug,
		status,
		user_id,
		files
	FROM
	    posts
	`, lang, lang, lang)

	whereQuery := ` WHERE deleted_at IS NULL AND status = TRUE AND id = ?`

	var files pq.StringArray
	err := r.DB.QueryRowContext(ctx, selectQuery+whereQuery, postID).Scan(
		&response.ID,
		&title,
		&content,
		&shortContent,
		&response.Slug,
		&response.Status,
		&response.UserID,
		&files,
	)
	if err != nil {
		return entity.GetPostResponse{}, err
	}

	response.Files = files

	response.Title = map[string]string{lang: title}
	response.Content = map[string]string{lang: content}
	response.ShortContent = map[string]string{lang: shortContent}

	return response, nil
}

func (r *Repo) Create(ctx context.Context, post entity.CreatePostRequest) (entity.CreatePostResponse, error) {
	var (
		title        []byte
		content      []byte
		shortContent []byte
		response     entity.CreatePostResponse
	)

	titleBytes, err := json.Marshal(&post.Title)
	if err != nil {
		return entity.CreatePostResponse{}, err
	}
	contentBytes, err := json.Marshal(&post.Content)
	if err != nil {
		return entity.CreatePostResponse{}, err
	}
	shortContentBytes, err := json.Marshal(&post.ShortContent)
	if err != nil {
		return entity.CreatePostResponse{}, err
	}

	err = r.DB.NewInsert().
		Model(&entity.Posts{
			Title:        string(titleBytes),
			Content:      string(contentBytes),
			ShortContent: string(shortContentBytes),
			Slug:         post.Slug,
			Status:       post.Status,
			UserID:       post.UserID,
			CreatedBy:    &post.CreatedBy,
		}).
		Returning("id, title, content, short_content, slug, user_id, files").
		Scan(ctx,
			&response.ID,
			&title,
			&content,
			&shortContent,
			&response.Slug,
			&response.Status,
			&response.UserID,
			&response.Files,
		)

	if err != nil {
		return entity.CreatePostResponse{}, err
	}

	if err := json.Unmarshal(title, &response.Title); err != nil {
		return entity.CreatePostResponse{}, err
	}
	if err := json.Unmarshal(content, &response.Content); err != nil {
		return entity.CreatePostResponse{}, err
	}
	if err := json.Unmarshal(shortContent, &response.ShortContent); err != nil {
		return entity.CreatePostResponse{}, err
	}

	return response, nil
}

func (r *Repo) Update(ctx context.Context, post entity.UpdatePostRequest) (entity.UpdatePostResponse, error) {
	var (
		title        []byte
		content      []byte
		shortContent []byte
		response     entity.UpdatePostResponse
	)

	titleBytes, err := json.Marshal(&post.Title)
	if err != nil {
		return entity.UpdatePostResponse{}, err
	}
	contentBytes, err := json.Marshal(&post.Content)
	if err != nil {
		return entity.UpdatePostResponse{}, err
	}
	shortContentBytes, err := json.Marshal(&post.ShortContent)
	if err != nil {
		return entity.UpdatePostResponse{}, err
	}

	err = r.DB.NewUpdate().
		Model(&entity.Posts{
			ID:           post.ID,
			Title:        string(titleBytes),
			Content:      string(contentBytes),
			ShortContent: string(shortContentBytes),
			Slug:         post.Slug,
			Status:       post.Status,
			UserID:       post.UserID,
			Files:        post.Files,
			UpdatedBy:    &post.UpdatedBy,
		}).
		Where("deleted_at IS NULL AND status = TRUE AND id = ?", post.ID).
		Returning("id, title, content, short_content, slug, user_id, files").
		Scan(ctx,
			&response.ID,
			&title,
			&content,
			&shortContent,
			&response.Slug,
			&response.Status,
			&response.UserID,
			&response.Files,
		)

	if err != nil {
		return entity.UpdatePostResponse{}, err
	}

	if err := json.Unmarshal(title, &response.Title); err != nil {
		return entity.UpdatePostResponse{}, err
	}
	if err := json.Unmarshal(content, &response.Content); err != nil {
		return entity.UpdatePostResponse{}, err
	}
	if err := json.Unmarshal(shortContent, &response.ShortContent); err != nil {
		return entity.UpdatePostResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateColumns(ctx context.Context, fields entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	var (
		title        string
		content      string
		shortContent string
		response     entity.UpdatePostResponse
	)

	updater := r.DB.NewUpdate().Table("menus")

	for key, value := range fields.Fields {
		if key == "title" {
			titleBytes, err := json.Marshal(value)
			if err != nil {
				return entity.UpdatePostResponse{}, err
			}

			updater.Set(key+" = ?", string(titleBytes))
		} else if key == "content" {
			contentBytes, err := json.Marshal(value)
			if err != nil {
				return entity.UpdatePostResponse{}, err
			}

			updater.Set(key+" = ?", string(contentBytes))
		} else if key == "short_content" {
			shorContentBytes, err := json.Marshal(value)
			if err != nil {
				return entity.UpdatePostResponse{}, err
			}

			updater.Set(key+" = ?", string(shorContentBytes))
		} else if key == "slug" {
			updater.Set(key+" = ?", value)
		} else if key == "status" {
			updater.Set(key+" = ?", value)
		} else if key == "user_id" {
			updater.Set(key+" = ?", value)
		} else if key == "files" {
			updater.Set(key+" = ?", value)
		}
	}

	err := updater.Where("deleted_at IS NULL AND status = TRUE AND id = ?", fields.ID).
		Returning("id, title, content, short_content, slug, user_id, files").
		Scan(ctx,
			&response.ID,
			&title,
			&content,
			&shortContent,
			&response.Slug,
			&response.Status,
			&response.UserID,
			&response.Files,
		)

	if err != nil {
		return entity.UpdatePostResponse{}, err
	}

	if err := json.Unmarshal([]byte(title), &response.Title); err != nil {
		return entity.UpdatePostResponse{}, err
	}
	if err := json.Unmarshal([]byte(content), &response.Content); err != nil {
		return entity.UpdatePostResponse{}, err
	}
	if err := json.Unmarshal([]byte(shortContent), &response.ShortContent); err != nil {
		return entity.UpdatePostResponse{}, err
	}

	return response, nil
}

func (r *Repo) Delete(ctx context.Context, postID, deletedBy int) (entity.DeletePostResponse, error) {
	result, err := r.DB.NewUpdate().
		Set("deleted_at = NOW()").
		Set("deleted_by = ?", deletedBy).
		Table("posts").
		Where("deleted_at IS NULL AND id = ?", postID).
		Exec(ctx)

	if err != nil {
		return entity.DeletePostResponse{}, err
	}

	rowEffects, err := result.RowsAffected()
	if err != nil {
		return entity.DeletePostResponse{}, err
	}

	if rowEffects == 0 {
		return entity.DeletePostResponse{}, errors.New("menu not found")
	}

	return entity.DeletePostResponse{
		Message: "success",
	}, nil
}

func (r *Repo) AddFile(ctx context.Context, fileURL string, postID int) error {
	selectQuery := fmt.Sprintf(
		`SELECT files FROM posts WHERE deleted_at IS NULL AND status = TRUE AND id = '%d'`, postID,
	)

	var files pq.StringArray

	if err := r.DB.QueryRowContext(ctx, selectQuery).Scan(&files); err != nil {
		return err
	}

	files = append(files, fileURL)

	insertQuery := fmt.Sprintf(`INSERT INTO posts (files) VALUES ('%v')`, files)

	result, err := r.DB.ExecContext(ctx, insertQuery)
	if err != nil {
		return err
	}

	rowEffects, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowEffects == 0 {
		return errors.New("file not found")
	}

	return nil
}
