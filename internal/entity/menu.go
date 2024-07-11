package entity

type Menus struct {
	ID       *int   `json:"id" bun:"id"`
	Title    string `json:"title" bun:"title"`
	Content  string `json:"content" bun:"content"`
	IsStatic bool   `json:"is_static" bun:"is_static"`
	Sort     int    `json:"sort" bun:"sort"`
	ParentID *int   `json:"parent_id" bun:"parent_id"`
	Status   bool   `json:"status" bun:"status"`
	Slug     string `json:"slug" bun:"slug"`
	Path     string `json:"path" bun:"path"`
}

type CreateMenuRequest struct {
	Title    map[string]string `json:"title" bun:"title"`
	Content  map[string]string `json:"content" bun:"content"`
	IsStatic bool              `json:"is_static" bun:"is_static"`
	Sort     int               `json:"sort" bun:"sort"`
	ParentID *int              `json:"parent_id" bun:"parent_id"`
	Status   bool              `json:"status" bun:"status"`
	Slug     string            `json:"slug" bun:"slug"`
	Path     string            `json:"path" bun:"path"`
}

type CreateMenuResponse struct {
	ID       int               `json:"id" bun:"id"`
	Title    map[string]string `json:"title" bun:"title"`
	Content  map[string]string `json:"content" bun:"content"`
	IsStatic bool              `json:"is_static" bun:"is_static"`
	Sort     int               `json:"sort" bun:"sort"`
	ParentID *int              `json:"parent_id" bun:"parent_id"`
	Status   bool              `json:"status" bun:"status"`
	Slug     string            `json:"slug" bun:"slug"`
	Path     string            `json:"path" bun:"path"`
}

type UpdateMenuRequest struct {
	ID       int               `json:"id" bun:"id"`
	Title    map[string]string `json:"title" bun:"title"`
	Content  map[string]string `json:"content" bun:"content"`
	IsStatic bool              `json:"is_static" bun:"is_static"`
	Sort     int               `json:"sort" bun:"sort"`
	ParentID *int              `json:"parent_id" bun:"parent_id"`
	Status   bool              `json:"status" bun:"status"`
	Slug     string            `json:"slug" bun:"slug"`
	Path     string            `json:"path" bun:"path"`
}

type UpdateMenuResponse struct {
	ID       int               `json:"id" bun:"id"`
	Title    map[string]string `json:"title" bun:"title"`
	Content  map[string]string `json:"content" bun:"content"`
	IsStatic bool              `json:"is_static" bun:"is_static"`
	Sort     int               `json:"sort" bun:"sort"`
	ParentID *int              `json:"parent_id" bun:"parent_id"`
	Status   bool              `json:"status" bun:"status"`
	Slug     string            `json:"slug" bun:"slug"`
	Path     string            `json:"path" bun:"path"`
}

type UpdateMenuColumnsRequest struct {
	ID     int               `json:"id" bun:"id"`
	Fields map[string]string `json:"fields" bun:"fields"`
}

type DeleteMenuResponse struct {
	Message string `json:"message" bun:"message"`
}

type GetMenuResponse struct {
	ID       int               `json:"id" bun:"id"`
	Title    map[string]string `json:"title" bun:"title"`
	Content  map[string]string `json:"content" bun:"content"`
	IsStatic bool              `json:"is_static" bun:"is_static"`
	Sort     int               `json:"sort" bun:"sort"`
	ParentID *int              `json:"parent_id" bun:"parent_id"`
	Status   bool              `json:"status" bun:"status"`
	Slug     string            `json:"slug" bun:"slug"`
	Path     string            `json:"path" bun:"path"`
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
