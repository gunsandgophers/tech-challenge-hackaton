package clients

import (
	"context"
	"tech-challenge-hackaton/internal/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSClient struct {
	client *sqs.Client
}

func NewSQSClient(
	awsRegion, awsAccessKeyID, awsSercretAccessKey string,
	awsBaseEndpoint *string,
) *SQSClient {

	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(awsRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSercretAccessKey, ""),
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

func (s *SQSClient) ReceiveMessages(maxMessages int, queueURL string) ([]types.Message, error) {
	resp, err := s.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: int32(maxMessages),
		VisibilityTimeout:   600,
		WaitTimeSeconds:     10,
	})
	if err != nil {
		return nil, err
	}
	return resp.Messages, nil
}

func (s *SQSClient) DeleteMessage(receiptHandle string, queueURL string) error {
	_, err := s.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	})
	if err != nil {
		return err
	}
	return nil
}
