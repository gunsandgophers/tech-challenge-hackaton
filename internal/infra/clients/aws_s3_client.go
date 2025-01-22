package clients

import (
	"context"
	"fmt"
	"mime/multipart"
	"tech-challenge-hackaton/internal/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
}

func NewS3Client(
	awsRegion, awsAcessKeyID, awsSercretAccessKey string,
	awsBaseEndpoint *string,
) *S3Client {
	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(awsRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(awsAcessKeyID, awsSercretAccessKey, ""),
			),
		),
	)
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = awsBaseEndpoint
	})
	return &S3Client{
		client: client,
	}
}

func (s *S3Client) UploadFile(
	filename string,
	file multipart.File,
	key string,
	awsBucketName string,
) (string, error) {
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(awsBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", nil
	}
	return fmt.Sprint(awsBucketName, "/", key), nil
}
