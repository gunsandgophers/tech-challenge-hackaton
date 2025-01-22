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

func NewAwsS3Service(client *clients.S3Client, awsBucketName string) *AwsS3Service {
	return &AwsS3Service{
		client: client,
		awsBucketName: awsBucketName,
	}
}

func (s *AwsS3Service) UploadVideo(filename string, file multipart.File) (string, error) {
	key := fmt.Sprint("video/", filename)
	return s.client.UploadFile(
		filename,
		file,
		key,
		s.awsBucketName,
	)
}
