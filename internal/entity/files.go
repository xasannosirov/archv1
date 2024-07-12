package entity

import "mime/multipart"

type FileUploadRequest struct {
	Category string         `json:"category"`
	File     multipart.File `json:"file"`
	ObjectID int            `json:"object_id"`
}

type FileUploadResponse struct {
	FileURL string `json:"file_url"`
}
