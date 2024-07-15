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
			files        pq.StringArray
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
			&files,
		)
		if err != nil {
			return entity.ListPostResponse{}, err
		}

		post.Title = map[string]string{lang: title}
		post.Content = map[string]string{lang: content}
		post.ShortContent = map[string]string{lang: shortContent}

		post.Files = files

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
		files        pq.StringArray
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
		Returning("id, title, content, short_content, slug, status, user_id").
		Scan(ctx,
			&response.ID,
			&title,
			&content,
			&shortContent,
			&response.Slug,
			&response.Status,
			&response.UserID,
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
		files        pq.StringArray
		response     entity.UpdatePostResponse
	)

	err := r.DB.NewUpdate().
		Table("posts").
		Set("title = ?", post.Title).
		Set("content = ?", post.Content).
		Set("short_content = ?", post.ShortContent).
		Set("slug = ?", post.Slug).
		Set("status = ?", post.Status).
		Set("user_id = ?", post.UserID).
		Set("updated_by = ?", post.UpdatedBy).
		Set("updated_at = NOW()").
		Where("deleted_at IS NULL AND status = TRUE AND id = ?", post.ID).
		Returning("id, title, content, short_content, slug, status, user_id, files").
		Scan(ctx,
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

	response.Files = files

	return response, nil
}

func (r *Repo) UpdateColumns(ctx context.Context, fields entity.UpdatePostColumnsRequest) (entity.UpdatePostResponse, error) {
	var (
		title        string
		content      string
		shortContent string
		files        pq.StringArray
		response     entity.UpdatePostResponse
	)

	updater := r.DB.NewUpdate().Table("posts")

	for key, value := range fields.Fields {
		if key == "title" {
			updater.Set(key+" = ?", value)
		} else if key == "content" {
			updater.Set(key+" = ?", value)
		} else if key == "short_content" {
			updater.Set(key+" = ?", value)
		} else if key == "slug" {
			updater.Set(key+" = ?", value)
		} else if key == "status" {
			updater.Set(key+" = ?", value)
		} else if key == "user_id" {
			updater.Set(key+" = ?", value)
		} else if key == "files" {
			updater.Set(key+" = ?", value)
		} else if key == "updated_by" {
			updater.Set(key+" = ?", value)
		}
	}

	updater.Set("updated_at = NOW()")

	err := updater.Where("deleted_at IS NULL AND status = TRUE AND id = ?", fields.ID).
		Returning("id, title, content, short_content, slug, status, user_id, files").
		Scan(ctx,
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

	response.Files = files

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
		return entity.DeletePostResponse{}, errors.New("post not found")
	}

	return entity.DeletePostResponse{
		Message: "success",
	}, nil
}

func (r *Repo) AddFile(ctx context.Context, fileURL string, postID, updatedBy int) error {
	updateQuery := fmt.Sprintf(`
	UPDATE posts 
	SET 
	    files = array_append(files, '%v'),
	    updated_by = '%d'
	WHERE deleted_at IS NULL AND status = TRUE AND id = '%d'
	`, fileURL, updatedBy, postID)

	result, err := r.DB.ExecContext(ctx, updateQuery)
	if err != nil {
		return err
	}

	rowEffects, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowEffects == 0 {
		return errors.New("post not found for append file")
	}

	return nil
}
