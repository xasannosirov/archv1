package menu

import (
	"archv1/internal/entity"
	"archv1/internal/pkg/repo/postgres"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Repo struct {
	DB *postgres.DB
}

func NewMenuRepo(DB *postgres.DB) MenuRepository {
	return &Repo{
		DB: DB,
	}
}

func (r *Repo) getChildMenus(ctx context.Context, parentID int, lang string) ([]interface{}, error) {
	selectQuery := fmt.Sprintf(`
	SELECT 
	    id, 
	    title ->> '%s' as title, 
	    content ->> '%s' as content, 
	    is_static, 
	    sort, 
	    parent_id, 
	    slug, 
	    path 
	FROM menus`, lang, lang)

	filterQuery := fmt.Sprintf(" WHERE deleted_at IS NULL AND parent_id = %d ORDER BY sort", parentID)

	rows, err := r.DB.QueryContext(ctx, selectQuery+filterQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []interface{}
	for rows.Next() {
		var (
			title   string
			content string
			menu    entity.ParentMenuWithChildren
		)
		err = rows.Scan(
			&menu.ParentMenu.ID,
			&title,
			&content,
			&menu.ParentMenu.IsStatic,
			&menu.ParentMenu.Sort,
			&menu.ParentMenu.ParentID,
			&menu.ParentMenu.Slug,
			&menu.ParentMenu.Path,
		)
		if err != nil {
			return nil, err
		}

		menu.ParentMenu.Title = map[string]string{lang: title}
		menu.ParentMenu.Content = map[string]string{lang: content}

		menu.Children, err = r.getChildMenus(ctx, menu.ParentMenu.ID, lang)
		if err != nil {
			return nil, err
		}

		response = append(response, &menu)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repo) GetSiteMenus(ctx context.Context, filter entity.Filter, lang string) (entity.SiteMenuListResponse, error) {
	var response entity.SiteMenuListResponse
	offset := filter.Limit * (filter.Page - 1)

	selectQuery := fmt.Sprintf(`
	SELECT 
    	id, 
    	title ->> '%s' as title, 
    	content ->> '%s' as content, 
    	is_static, 
   		sort, 
    	parent_id, 
    	slug, 
    	path 
	FROM menus`, lang, lang)

	filterQuery := fmt.Sprintf(
		" WHERE deleted_at IS NULL AND parent_id IS NULL ORDER BY sort LIMIT %d OFFSET %d",
		filter.Limit, offset,
	)

	rows, err := r.DB.QueryContext(ctx, selectQuery+filterQuery)
	if err != nil {
		return entity.SiteMenuListResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			title              string
			content            string
			parentWithChildren entity.ParentMenuWithChildren
		)
		err := rows.Scan(
			&parentWithChildren.ParentMenu.ID,
			&title,
			&content,
			&parentWithChildren.ParentMenu.IsStatic,
			&parentWithChildren.ParentMenu.Sort,
			&parentWithChildren.ParentMenu.ParentID,
			&parentWithChildren.ParentMenu.Slug,
			&parentWithChildren.ParentMenu.Path,
		)
		if err != nil {
			return entity.SiteMenuListResponse{}, err
		}

		parentWithChildren.ParentMenu.Title = map[string]string{lang: title}
		parentWithChildren.ParentMenu.Content = map[string]string{lang: content}

		children, err := r.getChildMenus(ctx, parentWithChildren.ParentMenu.ID, lang)
		if err != nil {
			return entity.SiteMenuListResponse{}, err
		}

		parentWithChildren.Children = children
		response.SiteMenus = append(response.SiteMenus, parentWithChildren)
	}

	if err = rows.Err(); err != nil {
		return entity.SiteMenuListResponse{}, err
	}

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM menus WHERE deleted_at IS NULL AND parent_id IS NULL")

	if err := r.DB.QueryRowContext(ctx, countQuery).Scan(&response.Total); err != nil {
		return entity.SiteMenuListResponse{}, err
	}

	return response, nil
}

func (r *Repo) List(ctx context.Context, filter entity.Filter, lang string) (entity.ListMenuResponse, error) {
	var response entity.ListMenuResponse
	offset := filter.Limit * (filter.Page - 1)

	selectQuery := fmt.Sprintf(`
	SELECT 
	    id, 
	    title ->> '%s', 
	    content ->> '%s', 
	    is_static, 
	    sort, 
	    parent_id, 
	    slug, 
	    path 
	FROM menus`, lang, lang)

	whereQuery := ` WHERE deleted_at IS NULL`

	limitQuery := fmt.Sprintf(" LIMIT %d OFFSET %d", filter.Limit, offset)

	rows, err := r.DB.QueryContext(ctx, selectQuery+whereQuery+limitQuery)
	if err != nil {
		return entity.ListMenuResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			title   string
			content string
			menu    entity.GetMenuResponse
		)
		err := rows.Scan(
			&menu.ID,
			&title,
			&content,
			&menu.IsStatic,
			&menu.Sort,
			&menu.ParentID,
			&menu.Slug,
			&menu.Path,
		)
		if err != nil {
			return entity.ListMenuResponse{}, err
		}

		menu.Title = map[string]string{lang: title}
		menu.Content = map[string]string{lang: content}

		response.Menus = append(response.Menus, &menu)
	}

	if err := rows.Err(); err != nil {
		return entity.ListMenuResponse{}, err
	}

	totalQuery := `SELECT COUNT(*) FROM menus WHERE deleted_at IS NULL`
	if err := r.DB.QueryRowContext(ctx, totalQuery).Scan(&response.Total); err != nil {
		return entity.ListMenuResponse{}, err
	}

	return response, nil
}

func (r *Repo) GetByID(ctx context.Context, menuID int, lang string) (entity.GetMenuResponse, error) {
	var (
		content  string
		title    string
		response entity.GetMenuResponse
	)

	selectQuery := fmt.Sprintf(`
	SELECT 
	    id, 
	    title ->> '%s', 
	    content ->> '%s', 
	    is_static, 
	    sort, 
	    parent_id, 
	    slug, 
	    path 
	FROM menus`, lang, lang)

	whereQuery := ` WHERE deleted_at IS NULL AND id = ?`

	err := r.DB.QueryRow(selectQuery+whereQuery, menuID).Scan(
		&response.ID,
		&title,
		&content,
		&response.IsStatic,
		&response.Sort,
		&response.ParentID,
		&response.Slug,
		&response.Path,
	)
	if err != nil {
		return entity.GetMenuResponse{}, err
	}

	response.Title = map[string]string{lang: title}
	response.Content = map[string]string{lang: content}

	return response, nil
}

func (r *Repo) Create(ctx context.Context, menu entity.CreateMenuRequest) (entity.CreateMenuResponse, error) {
	var (
		content  []byte
		title    []byte
		response entity.CreateMenuResponse
	)

	titleBytes, err := json.Marshal(menu.Title)
	if err != nil {
		return entity.CreateMenuResponse{}, err
	}
	contentBytes, err := json.Marshal(menu.Content)
	if err != nil {
		return entity.CreateMenuResponse{}, err
	}

	err = r.DB.NewInsert().
		Model(&entity.Menus{
			Title:    string(titleBytes),
			Content:  string(contentBytes),
			IsStatic: menu.IsStatic,
			Sort:     menu.Sort,
			ParentID: menu.ParentID,
			Status:   menu.Status,
			Slug:     menu.Slug,
			Path:     menu.Path,
		}).
		Returning("id, content, title, is_static, sort, parent_id, slug, path").
		Scan(ctx,
			&response.ID,
			&title,
			&content,
			&response.IsStatic,
			&response.Sort,
			&response.ParentID,
			&response.Slug,
			&response.Path,
		)

	if err != nil {
		return entity.CreateMenuResponse{}, err
	}

	if err := json.Unmarshal(content, &response.Content); err != nil {
		return entity.CreateMenuResponse{}, err
	}
	if err := json.Unmarshal(title, &response.Title); err != nil {
		return entity.CreateMenuResponse{}, err
	}

	return response, nil
}

func (r *Repo) Update(ctx context.Context, menu entity.UpdateMenuRequest) (entity.UpdateMenuResponse, error) {
	var (
		content  []byte
		title    []byte
		response entity.UpdateMenuResponse
	)

	titleBytes, err := json.Marshal(menu.Title)
	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}
	contentBytes, err := json.Marshal(menu.Content)
	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	err = r.DB.NewUpdate().Model(&entity.Menus{
		ID:       &menu.ID,
		Title:    string(titleBytes),
		Content:  string(contentBytes),
		IsStatic: menu.IsStatic,
		Sort:     menu.Sort,
		ParentID: menu.ParentID,
		Status:   menu.Status,
		Slug:     menu.Slug,
		Path:     menu.Path,
	}).
		Where("deleted_at IS NULL AND status = TRUE AND id = ?", menu.ID).
		Returning("id, content, title, is_static, sort, parent_id, slug, path").
		Scan(ctx,
			&response.ID,
			&title,
			&content,
			&response.IsStatic,
			&response.Sort,
			&response.ParentID,
			&response.Slug,
			&response.Path,
		)

	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	if err := json.Unmarshal(content, &response.Content); err != nil {
		return entity.UpdateMenuResponse{}, err
	}
	if err := json.Unmarshal(title, &response.Title); err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	return response, nil
}

func (r *Repo) UpdateColumns(ctx context.Context, fields entity.UpdateMenuColumnsRequest) (entity.UpdateMenuResponse, error) {
	var (
		content  string
		title    string
		response entity.UpdateMenuResponse
	)
	updater := r.DB.NewUpdate().Table("menus")

	for key, value := range fields.Fields {
		if key == "title" {
			titleBytes, err := json.Marshal(value)
			if err != nil {
				return entity.UpdateMenuResponse{}, err
			}

			updater.Set(key+" = ?", titleBytes)
		} else if key == "content" {
			contentBytes, err := json.Marshal(value)
			if err != nil {
				return entity.UpdateMenuResponse{}, err
			}

			updater.Set(key+" = ?", contentBytes)
		} else if key == "is_static" {
			updater.Set(key+" = ?", value)
		} else if key == "parent_id" {
			updater.Set(key+" = ?", value)
		} else if key == "slug" {
			updater.Set(key+" = ?", value)
		} else if key == "path" {
			updater.Set(key+" = ?", value)
		}
	}

	err := updater.Where("deleted_at IS NULL AND status = TRUE AND id = ?", fields.ID).
		Returning("id, title, content, is_static, sort, parent_id, slug, path").
		Scan(
			ctx,
			&response.ID,
			&title,
			&content,
			&response.IsStatic,
			&response.Sort,
			&response.ParentID,
			&response.Slug,
			&response.Path,
		)

	if err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	if err := json.Unmarshal([]byte(title), &response.Title); err != nil {
		return entity.UpdateMenuResponse{}, err
	}
	if err := json.Unmarshal([]byte(content), &response.Content); err != nil {
		return entity.UpdateMenuResponse{}, err
	}

	return response, nil
}

func (r *Repo) Delete(ctx context.Context, menuID, deletedBy int) (entity.DeleteMenuResponse, error) {
	result, err := r.DB.NewUpdate().
		Set("deleted_at = NOW()").
		Set("deleted_by = ?", deletedBy).
		Table("menus").
		Where("deleted_at IS NULL AND id = ?", menuID).
		Exec(ctx)

	if err != nil {
		return entity.DeleteMenuResponse{}, err
	}

	rowEffects, err := result.RowsAffected()
	if err != nil {
		return entity.DeleteMenuResponse{}, err
	}

	if rowEffects == 0 {
		return entity.DeleteMenuResponse{}, errors.New("menu not found")
	}

	return entity.DeleteMenuResponse{
		Message: "success",
	}, nil
}
