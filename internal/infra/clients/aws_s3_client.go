package clients

import (
	"context"
	"log"
	"tech-challenge-hackaton/internal/infra/config"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client() *s3.Client {
	cfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithRegion(config.AWS_REGION),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				config.AWS_ACCESS_KEY_ID, config.AWS_SECRET_ACCESS_KEY, ""),
		),
	)
	if err != nil {
		log.Println("Unable to create AWS S3 client", err)
		panic(1)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = config.AWS_BASE_ENDPOINT
	})

	return client
}
