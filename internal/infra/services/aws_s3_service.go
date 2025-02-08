package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"tech-challenge-hackaton/internal/infra/clients"
)

type AwsS3Service struct {
	client        *clients.S3Client
	awsBucketName string
}

const videoDir string = "videos"
const frameDir string = "frames"
const localDir string = "tmp"

func NewAwsS3Service(client *clients.S3Client, awsBucketName string) *AwsS3Service {
	return &AwsS3Service{
		client:        client,
		awsBucketName: awsBucketName,
	}
}

func (s *AwsS3Service) GetExternalVideoDir() string {
	return videoDir
}

func (s *AwsS3Service) GetExternalFramesDir() string {
	return frameDir
}

func (s *AwsS3Service) GetLocalVideoDir(videoID string) string {
	return filepath.Join(localDir, videoID)
}

func (s *AwsS3Service) UploadVideo(filename string, file multipart.File) (string, error) {
	filenameCompleteExt := fmt.Sprintf("%s/%s", s.GetExternalVideoDir(), filename)
	return s.client.UploadFile(
		filenameCompleteExt,
		file,
		s.awsBucketName,
	)
}

func (s *AwsS3Service) DownloadVideo(videoID, filename string) (string, error) {
	filenameCompleteExt := fmt.Sprintf("%s/%s", s.GetExternalVideoDir(), filename)
	return s.client.DownloadFile(
		s.GetLocalVideoDir(videoID),
		filename,
		filenameCompleteExt,
		s.awsBucketName,
	)
}

func (s *AwsS3Service) DownloadZipFrames(videoID string) ([]byte, error) {
	filenameCompleteExt := fmt.Sprintf("%s/%s.zip", s.GetExternalFramesDir(), videoID)
	return s.client.GetFile(filenameCompleteExt, s.awsBucketName)
}

func (s *AwsS3Service) UploadZipFrames(filename string, file multipart.File) (string, error) {
	filenameCompleteExt := fmt.Sprintf("%s/%s", s.GetExternalFramesDir(), filename)
	return s.client.UploadFile(
		filenameCompleteExt,
		file,
		s.awsBucketName,
	)
}
