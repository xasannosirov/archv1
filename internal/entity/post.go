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
	Content      map[string]string `json:"content" bun:"content"`
	ShortContent map[string]string `json:"short_content" bun:"short_content"`
	Slug         string            `json:"slug" bun:"slug"`
	Status       bool              `json:"status" bun:"status"`
	UserID       int               `json:"user_id" bun:"user_id"`
	CreatedBy    int               `json:"-" bun:"created_by"`
}

type CreatePostResponse struct {
	ID           int               `json:"id" bun:"id"`
	Title        map[string]string `json:"title" bun:"title"`
	Content      map[string]string `json:"content" bun:"content"`
	ShortContent map[string]string `json:"short_content" bun:"short_content"`
	Slug         string            `json:"slug" bun:"slug"`
	Status       bool              `json:"status" bun:"status"`
	UserID       int               `json:"user_id" bun:"user_id"`
	Files        []string          `json:"files" bun:"files"`
}

type UpdatePostRequest struct {
	ID           int               `json:"id" bun:"id"`
	Title        map[string]string `json:"title" bun:"title"`
	Content      map[string]string `json:"content" bun:"content"`
	ShortContent map[string]string `json:"short_content" bun:"short_content"`
	Slug         string            `json:"slug" bun:"slug"`
	Status       bool              `json:"status" bun:"status"`
	UserID       int               `json:"user_id" bun:"user_id"`
	Files        []string          `json:"files" bun:"files"`
	UpdatedBy    int               `json:"-" bun:"updated_by"`
}

type UpdatePostResponse struct {
	ID           int               `json:"id" bun:"id"`
	Title        map[string]string `json:"title" bun:"title"`
	Content      map[string]string `json:"content" bun:"content"`
	ShortContent map[string]string `json:"short_content" bun:"short_content"`
	Slug         string            `json:"slug" bun:"slug"`
	Status       bool              `json:"status" bun:"status"`
	UserID       int               `json:"user_id" bun:"user_id"`
	Files        []string          `json:"files" bun:"files"`
}

type UpdatePostColumnsRequest struct {
	ID     int               `json:"id" bun:"id"`
	Fields map[string]string `json:"fields" bun:"fields"`
}

type DeletePostResponse struct {
	Message string `json:"message" bun:"message"`
}

type GetPostResponse struct {
	ID           int               `json:"id" bun:"id"`
	Title        map[string]string `json:"title" bun:"title"`
	Content      map[string]string `json:"content" bun:"content"`
	ShortContent map[string]string `json:"short_content" bun:"short_content"`
	Slug         string            `json:"slug" bun:"slug"`
	Status       bool              `json:"status" bun:"status"`
	UserID       int               `json:"user_id" bun:"user_id"`
	Files        []string          `json:"files" bun:"files"`
}

type ListPostResponse struct {
	Posts []*GetPostResponse `json:"menus" bun:"menu"`
	Total int64              `json:"total" bun:"total"`
}
