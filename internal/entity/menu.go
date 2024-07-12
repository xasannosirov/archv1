package entity

type Menus struct {
	ID        *int     `json:"id" bun:"id"`
	Title     string   `json:"title" bun:"title"`
	Content   string   `json:"content" bun:"content"`
	IsStatic  bool     `json:"is_static" bun:"is_static"`
	Sort      int      `json:"sort" bun:"sort"`
	ParentID  *int     `json:"parent_id" bun:"parent_id"`
	Status    bool     `json:"status" bun:"status"`
	Slug      string   `json:"slug" bun:"slug"`
	Path      string   `json:"path" bun:"path"`
	Files     []string `json:"files" bun:"files"`
	CreatedBy *int     `json:"created_by" bun:"created_by"`
	UpdatedBy *int     `json:"updated_by" bun:"updated_by"`
}

type CreateMenuRequest struct {
	Title     map[string]string `json:"title"`
	Content   map[string]string `json:"content"`
	IsStatic  bool              `json:"is_static"`
	Sort      int               `json:"sort"`
	ParentID  *int              `json:"parent_id"`
	Status    bool              `json:"status"`
	Slug      string            `json:"slug"`
	Path      string            `json:"path"`
	CreatedBy int               `json:"created_by"`
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
	ID        int               `json:"id"`
	Title     map[string]string `json:"title"`
	Content   map[string]string `json:"content"`
	IsStatic  bool              `json:"is_static"`
	Sort      int               `json:"sort"`
	ParentID  *int              `json:"parent_id"`
	Status    bool              `json:"status"`
	Slug      string            `json:"slug"`
	Path      string            `json:"path"`
	UpdatedBy int               `json:"updated_by"`
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
	ID     int               `json:"id"`
	Fields map[string]string `json:"fields"`
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
