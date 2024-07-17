package entity

import "github.com/lib/pq"

type Posts struct {
	ID           *int           `json:"id" bun:"id"`
	Title        string         `json:"title" bun:"title"`
	Content      string         `json:"content" bun:"content"`
	ShortContent string         `json:"short_content" bun:"short_content"`
	Slug         string         `json:"slug" bun:"slug"`
	Status       bool           `json:"status" bun:"status"`
	UserID       int            `json:"user_id" bun:"user_id"`
	Files        pq.StringArray `json:"files" bun:"files"`
	CreatedBy    *int           `json:"created_by" bun:"created_by"`
	UpdatedBy    *int           `json:"updated_by" bun:"updated_by"`
}

type CreatePostRequest struct {
	Title        map[string]string `json:"title" xml:"title" yaml:"title" toml:"title" query:"title" form:"title"`
	Content      map[string]string `json:"content" xml:"content" yaml:"content" toml:"content" query:"content" form:"content"`
	ShortContent map[string]string `json:"short_content" xml:"short_content" yaml:"short_content" toml:"short_content" form:"short_content"`
	Slug         string            `json:"slug" xml:"slug" yaml:"slug" toml:"slug" form:"slug"`
	Status       bool              `json:"status" xml:"status" yaml:"status" toml:"status" form:"status"`
	UserID       int               `json:"user_id" bun:"user_id" xml:"user_id" yaml:"user_id" toml:"user_id"`
	CreatedBy    int               `json:"-" bun:"created_by"`
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
	ID           int               `json:"id" xml:"id" yaml:"id" toml:"id" query:"id" form:"id"`
	Title        map[string]string `json:"title" xml:"title" yaml:"title" toml:"title" query:"title" form:"title"`
	Content      map[string]string `json:"content" xml:"content" yaml:"content" toml:"content" query:"content" form:"content"`
	ShortContent map[string]string `json:"short_content" xml:"short_content" yaml:"short_content" toml:"short_content" query:"short_content" form:"short_content"`
	Slug         string            `json:"slug" xml:"slug" yaml:"slug" toml:"slug" query:"slug" form:"slug"`
	Status       bool              `json:"status" xml:"status" yaml:"status" toml:"status" query:"status" form:"status"`
	UserID       int               `json:"user_id" xml:"user_id" yaml:"user_id" toml:"user_id" query:"user_id" form:"user_id"`
	Files        []string          `json:"files" xml:"files" yaml:"files" toml:"files" query:"files" form:"files"`
	UpdatedBy    int               `json:"-" bun:"updated_by"`
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
	ID     int               `json:"id" xml:"id" yaml:"id" toml:"id" query:"id" form:"id"`
	Fields map[string]string `json:"fields" xml:"fields" yaml:"fields" toml:"fields" query:"fields" form:"fields"`
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
