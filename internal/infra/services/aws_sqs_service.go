package services

import (
	"context"
	"encoding/json"
	"tech-challenge-hackaton/internal/core/entities"
	"tech-challenge-hackaton/internal/infra/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type AwsSQSService struct {
	client *sqs.Client
}

func NewAwsSQSService(client *sqs.Client) *AwsSQSService {
	return &AwsSQSService{client: client}
}

func (s *AwsSQSService) sendMessage(message interface{}, queueURL string) error {

	body, err := json.Marshal(message)
	if err != nil {
		return nil
	}

	_, err = s.client.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageBody: aws.String(string(body)),
		QueueUrl:    aws.String(queueURL),
	})

	return err
}

func (s *AwsSQSService) SendVideoForProcessing(video *entities.Video) error {
	return s.sendMessage(video, config.QUEUE_PROCESS_VIDEO)
}
