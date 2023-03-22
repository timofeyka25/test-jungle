package dto

type UploadedFileDTO struct {
	FileURL string `json:"file_url"`
}

type ImageDTO struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	ImagePath string `json:"image_path"`
	ImageURL  string `json:"image_url"`
}
