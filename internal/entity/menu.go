package entity

import "github.com/lib/pq"

type Menus struct {
	ID        *int           `json:"id" bun:"id"`
	Title     string         `json:"title" bun:"title"`
	Content   string         `json:"content" bun:"content"`
	IsStatic  bool           `json:"is_static" bun:"is_static"`
	Sort      int            `json:"sort" bun:"sort"`
	ParentID  *int           `json:"parent_id" bun:"parent_id"`
	Status    bool           `json:"status" bun:"status"`
	Slug      string         `json:"slug" bun:"slug"`
	Path      string         `json:"path" bun:"path"`
	Files     pq.StringArray `json:"files" bun:"files"`
	CreatedBy *int           `json:"created_by" bun:"created_by"`
	UpdatedBy *int           `json:"updated_by" bun:"updated_by"`
}

type CreateMenuRequest struct {
	Title     map[string]string `json:"title" xml:"title" yaml:"title" toml:"title" form:"title" query:"title"`
	Content   map[string]string `json:"content" xml:"content" yaml:"content" toml:"content" form:"content" query:"content"`
	IsStatic  bool              `json:"is_static" xml:"is_static" yaml:"is_static" toml:"is_static" form:"is_static" query:"is_static"`
	Sort      int               `json:"sort" xml:"sort" yaml:"sort" toml:"sort" form:"sort" query:"sort"`
	ParentID  *int              `json:"parent_id" xml:"parent_id" yaml:"parent_id" toml:"parent_id" form:"parent_id" query:"parent_id"`
	Status    bool              `json:"status" xml:"status" yaml:"status" toml:"status" form:"status" query:"status"`
	Slug      string            `json:"slug" xml:"slug" yaml:"slug" toml:"slug" form:"slug" query:"slug"`
	Path      string            `json:"path" xml:"path" yaml:"path" toml:"path" form:"path" query:"path"`
	CreatedBy int               `json:"-" bun:"created_by"`
}

type CreateMenuResponse struct {
	ID       int               `json:"id"`
	Title    map[string]string `json:"title"`
	Content  map[string]string `json:"content"`
	IsStatic bool              `json:"is_static"`
	Sort     int               `json:"sort"`
	ParentID *int              `json:"parent_id"`
	Status   bool              `json:"status"`
	Slug     string            `json:"slug"`
	Path     string            `json:"path"`
	Files    []string          `json:"files"`
}

type UpdateMenuRequest struct {
	ID        int               `json:"id" xml:"id" yaml:"id" toml:"id" form:"id" query:"id"`
	Title     map[string]string `json:"title" xml:"title" yaml:"title" toml:"title" form:"title" query:"title"`
	Content   map[string]string `json:"content" xml:"content" yaml:"content" toml:"content" form:"content" query:"content"`
	IsStatic  bool              `json:"is_static" xml:"is_static" yaml:"is_static" toml:"is_static" form:"is_static" query:"is_static"`
	Sort      int               `json:"sort" xml:"sort" yaml:"sort" toml:"sort" form:"sort" query:"sort"`
	ParentID  *int              `json:"parent_id" xml:"parent_id" yaml:"parent_id" toml:"parent_id" form:"parent_id" query:"parent_id"`
	Status    bool              `json:"status" xml:"status" yaml:"status" toml:"status" form:"status" query:"status"`
	Slug      string            `json:"slug" xml:"slug" yaml:"slug" toml:"slug" form:"slug" query:"slug"`
	Path      string            `json:"path" xml:"path" yaml:"path" toml:"path" form:"path" query:"path"`
	UpdatedBy int               `json:"-" bun:"updated_by"`
}

type UpdateMenuResponse struct {
	ID       int               `json:"id"`
	Title    map[string]string `json:"title"`
	Content  map[string]string `json:"content"`
	IsStatic bool              `json:"is_static"`
	Sort     int               `json:"sort"`
	ParentID *int              `json:"parent_id"`
	Status   bool              `json:"status"`
	Slug     string            `json:"slug"`
	Path     string            `json:"path"`
	Files    []string          `json:"files"`
}

type UpdateMenuColumnsRequest struct {
	ID     int               `json:"id" xml:"id" yaml:"id" toml:"id" form:"id" query:"id"`
	Fields map[string]string `json:"fields" xml:"fields" yaml:"fields" toml:"fields" form:"fields" query:"fields"`
}

type DeleteMenuResponse struct {
	Message string `json:"message"`
}

type GetMenuResponse struct {
	ID       int               `json:"id"`
	Title    map[string]string `json:"title"`
	Content  map[string]string `json:"content"`
	IsStatic bool              `json:"is_static"`
	Sort     int               `json:"sort"`
	ParentID *int              `json:"parent_id"`
	Status   bool              `json:"status"`
	Slug     string            `json:"slug"`
	Path     string            `json:"path"`
	Files    []string          `json:"files"`
}

type ListMenuResponse struct {
	Menus []*GetMenuResponse `json:"menus"`
	Total int64              `json:"total"`
}

type ParentMenuWithChildren struct {
	ParentMenu GetMenuResponse `json:"parent_menu"`
	Children   interface{}     `json:"children"`
}

type SiteMenuListResponse struct {
	SiteMenus []ParentMenuWithChildren `json:"site_menus"`
	Total     int64                    `json:"total"`
}
