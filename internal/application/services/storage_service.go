package services

import (
	"mime/multipart"
)

type StorageServiceInterface interface {
	UploadVideo(filename string, file multipart.File) (string, error)
	DownloadVideo(videoID, filename string) (string, error)
	UploadZipFrames(filename string, file multipart.File) (string, error)
	DownloadZipFrames(videoID string) ([]byte, error)
	GetExternalVideoDir() string
	GetExternalFramesDir() string
	GetLocalVideoDir(videoID string) string
}
