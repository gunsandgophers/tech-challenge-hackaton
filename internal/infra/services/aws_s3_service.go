package services

import (
	"fmt"
	"mime/multipart"
	"tech-challenge-hackaton/internal/infra/clients"
)

type AwsS3Service struct {
	client *clients.S3Client
	awsBucketName string
}

const videoDir string = "videos"

func NewAwsS3Service(client *clients.S3Client, awsBucketName string) *AwsS3Service {
	return &AwsS3Service{
		client: client,
		awsBucketName: awsBucketName,
	}
}

func (s *AwsS3Service) UploadVideo(videoID, filename string, file multipart.File) (string, error) {
	filenameComplete := fmt.Sprintf("%s/%s-%s", videoDir, videoID, filename)
	return s.client.UploadFile(
		filenameComplete,
		file,
		s.awsBucketName,
	)
}
