package entity

type FileUploadRequest struct {
	Category string `json:"category"`
	ObjectID int    `json:"object_id"`
}

type FileUploadResponse struct {
	FileURL string `json:"file_url"`
}
