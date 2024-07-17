package entity

type FileUploadRequest struct {
	Category string `json:"category" xml:"category" yaml:"category" toml:"category" form:"category" query:"category"`
	ObjectID int    `json:"object_id" xml:"object_id" yaml:"object_id" toml:"object_id" form:"object_id" query:"object_id"`
}

type FileUploadResponse struct {
	FileURL string `json:"file_url"`
}
