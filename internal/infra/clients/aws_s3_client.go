package clients

import (
	"context"
	"tech-challenge-hackaton/internal/utils"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(AWSRegion, AWSAcessKeyID, AWSSercretAccessKey string, AWSBaseEndpoint *string) *s3.Client {
	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(AWSRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(AWSAcessKeyID, AWSSercretAccessKey, ""),
			),
		),
	)
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = AWSBaseEndpoint
	})
	return client
}
