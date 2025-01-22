package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AwsS3Service struct {
	client *s3.Client
	awsBucketName string
}

func NewAwsS3Service(client *s3.Client, awsBucketName string) *AwsS3Service {
	return &AwsS3Service{client: client, awsBucketName: awsBucketName}
}

func (s *AwsS3Service) UploadFile(filename string, file multipart.File) (string, error) {
	uploader := manager.NewUploader(s.client)

	key := fmt.Sprint("video/", filename)

	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.awsBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", nil
	}

	return fmt.Sprint(s.awsBucketName, "/", key), nil
}
