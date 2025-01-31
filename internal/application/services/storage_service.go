package services

import (
	"mime/multipart"
)

type StorageServiceInterface interface {
	UploadVideo(videoID, filename string, file multipart.File) (string, error)
}
