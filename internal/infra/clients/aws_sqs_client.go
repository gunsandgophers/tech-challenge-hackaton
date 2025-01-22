package clients

import (
	"context"
	"tech-challenge-hackaton/internal/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSClient struct {
	client *sqs.Client
}

func NewSQSClient(
	awsRegion, awsAcessKeyID, awsSercretAccessKey string,
	awsBaseEndpoint *string,
) *SQSClient {

	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(awsRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(awsAcessKeyID, awsSercretAccessKey, ""),
			),
		),
	)
	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = awsBaseEndpoint
	})
	return &SQSClient{
		client: client,
	}
}

func (s *SQSClient) SendMessage(body string, queueURL string) error {
	_, err := s.client.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageBody: aws.String(body),
		QueueUrl:    aws.String(queueURL),
	})
	return err
}
