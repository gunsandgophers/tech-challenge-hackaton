package services

import (
	"mime/multipart"
)

type StorageServiceInterface interface {
	UploadVideo(filename string, file multipart.File) (string, error)
}
