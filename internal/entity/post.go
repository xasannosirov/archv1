package entity

type Posts struct {
	ID           int      `json:"id" bun:"id"`
	Title        string   `json:"title" bun:"title"`
	Content      string   `json:"content" bun:"content"`
	ShortContent string   `json:"short_content" bun:"short_content"`
	Slug         string   `json:"slug" bun:"slug"`
	Status       bool     `json:"status" bun:"status"`
	UserID       int      `json:"user_id" bun:"user_id"`
	Files        []string `json:"files" bun:"files"`
	CreatedBy    *int     `json:"created_by" bun:"created_by"`
	UpdatedBy    *int     `json:"updated_by" bun:"updated_by"`
}

type CreatePostRequest struct {
	Title        map[string]string `json:"title"`
	Content      map[string]string `json:"content"`
	ShortContent map[string]string `json:"short_content"`
	Slug         string            `json:"slug"`
	Status       bool              `json:"status"`
	UserID       int               `json:"user_id"`
	CreatedBy    int               `json:"-"`
}

type CreatePostResponse struct {
	ID           int               `json:"id"`
	Title        map[string]string `json:"title"`
	Content      map[string]string `json:"content"`
	ShortContent map[string]string `json:"short_content"`
	Slug         string            `json:"slug"`
	Status       bool              `json:"status"`
	UserID       int               `json:"user_id"`
	Files        []string          `json:"files"`
}

type UpdatePostRequest struct {
	ID           int               `json:"id"`
	Title        map[string]string `json:"title"`
	Content      map[string]string `json:"content"`
	ShortContent map[string]string `json:"short_content"`
	Slug         string            `json:"slug"`
	Status       bool              `json:"status"`
	UserID       int               `json:"user_id"`
	Files        []string          `json:"files"`
	UpdatedBy    int               `json:"-"`
}

type UpdatePostResponse struct {
	ID           int               `json:"id"`
	Title        map[string]string `json:"title"`
	Content      map[string]string `json:"content"`
	ShortContent map[string]string `json:"short_content"`
	Slug         string            `json:"slug"`
	Status       bool              `json:"status"`
	UserID       int               `json:"user_id"`
	Files        []string          `json:"files"`
}

type UpdatePostColumnsRequest struct {
	ID     int               `json:"id"`
	Fields map[string]string `json:"fields"`
}

type DeletePostResponse struct {
	Message string `json:"message"`
}

type GetPostResponse struct {
	ID           int               `json:"id"`
	Title        map[string]string `json:"title"`
	Content      map[string]string `json:"content"`
	ShortContent map[string]string `json:"short_content"`
	Slug         string            `json:"slug"`
	Status       bool              `json:"status"`
	UserID       int               `json:"user_id"`
	Files        []string          `json:"files"`
}

type ListPostResponse struct {
	Posts []*GetPostResponse `json:"menus"`
	Total int64              `json:"total"`
}
