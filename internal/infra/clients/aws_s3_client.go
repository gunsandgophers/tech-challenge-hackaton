package clients

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
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
	awsRegion, awsAccessKeyID, awsSercretAccessKey string,
	awsBaseEndpoint *string,
) *S3Client {
	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(awsRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSercretAccessKey, ""),
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
	key string,
	file multipart.File,
	awsBucketName string,
) (string, error) {
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(awsBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Println("Error on upload")
		log.Println(err.Error())
		return "", nil
	}
	return fmt.Sprint(awsBucketName, "/", key), nil
}

func (s *S3Client) DownloadFile(targetDir string, filename string, key string, awsBucketName string) (string, error) {
	filenameCompleteLocal := filepath.Join(targetDir, filename)
	if err := os.MkdirAll(filepath.Dir(filenameCompleteLocal), 0775); err != nil {
		return filenameCompleteLocal, err
	}

	fd, err := os.Create(filenameCompleteLocal)
	if err != nil {
		return filenameCompleteLocal, err
	}
	defer fd.Close()

	downloader := manager.NewDownloader(s.client)
	_, err = downloader.Download(
		context.Background(),
		fd,
		&s3.GetObjectInput{Bucket: aws.String(awsBucketName), Key: aws.String(key)},
	)
	if err != nil {
		log.Println("Error on download")
		log.Println(err.Error())
		return "", err
	}
	return filenameCompleteLocal, nil
}
