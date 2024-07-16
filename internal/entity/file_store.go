package entity

type Folders struct {
	ID        *int   `json:"id" bun:"id"`
	Name      string `json:"name" bun:"name"`
	ParentID  *int   `json:"parent_id" bun:"parent_id"`
	CreatedBy *int   `json:"created_by" bun:"created_by"`
	UpdatedBy *int   `json:"updated_by" bun:"updated_by"`
}

type Files struct {
	ID        *int   `json:"id" bun:"id"`
	Type      string `json:"type" bun:"type"`
	Link      string `json:"link" bun:"link"`
	FolderID  *int   `json:"folder_id" bun:"folder_id"`
	CreatedBy *int   `json:"created_by" bun:"created_by"`
	UpdatedBy *int   `json:"updated_by" bun:"updated_by"`
}

type CreateFolderRequest struct {
	Name      string `json:"name"`
	ParentID  *int   `json:"parent_id"`
	CreatedBy int    `json:"-"`
}

type CreateFolderResponse struct {
	ID       *int   `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}

type UpdateFolderRequest struct {
	ID        *int   `json:"id"`
	Name      string `json:"name"`
	ParentID  *int   `json:"parent_id"`
	UpdatedBy int    `json:"-"`
}

type UpdateFolderColumnsRequest struct {
	Fields   map[string]string `json:"fields"`
	FolderID int               `json:"folder_id"`
}

type UpdateFolderResponse struct {
	ID       *int   `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}

type DeleteFolderResponse struct {
	Message string `json:"message"`
}

type GetFolderResponse struct {
	ID       *int   `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id"`
}

type ListFolderResponse struct {
	Folders []*GetFolderResponse `json:"folders"`
	Total   int64                `json:"total"`
}

type CreateFileRequest struct {
	Type      string `json:"type"`
	Link      string `json:"link"`
	FolderID  *int   `json:"folder_id"`
	CreatedBy int    `json:"-"`
}

type CreateFileResponse struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Link     string `json:"link"`
	FolderID *int   `json:"folder_id"`
}

type UpdateFileRequest struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Link      string `json:"link"`
	FolderID  *int   `json:"folder_id"`
	UpdatedBy int    `json:"-"`
}

type UpdateFileColumnsRequest struct {
	Fields map[string]string `json:"fields"`
	FileID int               `json:"file_id"`
}

type UpdateFileResponse struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Link     string `json:"link"`
	FolderID *int   `json:"folder_id"`
}

type DeleteFileResponse struct {
	Message string `json:"message"`
}

type GetFileResponse struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Link     string `json:"link"`
	FolderID *int   `json:"folder_id"`
}

type ListFileResponse struct {
	Files []*GetFileResponse `json:"files"`
	Total int64              `json:"total"`
}
