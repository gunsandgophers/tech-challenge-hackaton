package dtos

type VideoUploadDTO struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
}

type VideoUploadResponseDTO struct {
	Videos []*VideoUploadDTO `json:"videos"`
}
