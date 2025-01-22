package services

import (
	"mime/multipart"
)

type StorageServiceInterface interface {
	UploadFile(filename string, file multipart.File) (string, error)
}
