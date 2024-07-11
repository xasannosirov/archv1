package entity

type Posts struct {
	ID      int               `json:"id" bun:"id"`
	Title   map[string]string `json:"title" bun:"title"`
	Content map[string]string `json:"content" bun:"content"`
	UserID  int               `json:"user_id" bun:"user_id"`
}

type CreatePostRequest struct {
	Title   map[string]string `json:"title" bun:"title"`
	Content map[string]string `json:"content" bun:"content"`
	UserID  int               `json:"user_id" bun:"user_id"`
}

type CreatePostResponse struct {
	ID      int               `json:"id" bun:"id"`
	Title   map[string]string `json:"title" bun:"title"`
	Content map[string]string `json:"content" bun:"content"`
	UserID  int               `json:"user_id" bun:"user_id"`
}

type UpdatePostRequest struct {
	ID      int               `json:"id" bun:"id"`
	Title   map[string]string `json:"title" bun:"title"`
	Content map[string]string `json:"content" bun:"content"`
	UserID  int               `json:"user_id" bun:"user_id"`
}

type UpdatePostResponse struct {
	ID      int               `json:"id" bun:"id"`
	Title   map[string]string `json:"title" bun:"title"`
	Content map[string]string `json:"content" bun:"content"`
	UserID  int               `json:"user_id" bun:"user_id"`
}

type UpdatePostColumnsRequest struct {
	ID     int               `json:"id" bun:"id"`
	Fields map[string]string `json:"fields" bun:"fields"`
}

type GetPostResponse struct {
	ID      int               `json:"id" bun:"id"`
	Title   map[string]string `json:"title" bun:"title"`
	Content map[string]string `json:"content" bun:"content"`
	UserID  int               `json:"user_id" bun:"user_id"`
}

type ListPostResponse struct {
	Posts []*GetPostResponse `json:"menus" bun:"menu"`
	Total int64              `json:"total" bun:"total"`
}
