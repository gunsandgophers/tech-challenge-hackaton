package clients

import (
	"context"
	"log"
	"tech-challenge-hackaton/internal/infra/config"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSClient() *sqs.Client {
	cfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithRegion(config.AWS_REGION),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				config.AWS_ACCESS_KEY_ID, config.AWS_SECRET_ACCESS_KEY, ""),
		),
	)
	if err != nil {
		log.Println("Unable to create AWS SQS client", err)
		panic(1)
	}

	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = config.AWS_BASE_ENDPOINT
	})

	return client
}
