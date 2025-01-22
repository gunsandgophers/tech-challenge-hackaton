package clients

import (
	"context"
	"tech-challenge-hackaton/internal/utils"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSClient(AWSRegion, AWSAcessKeyID, AWSSercretAccessKey string, AWSBaseEndpoint *string) *sqs.Client {
	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(AWSRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(AWSAcessKeyID, AWSSercretAccessKey, ""),
			),
		),
	)
	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = AWSBaseEndpoint
	})
	return client
}
