package services

import (
	"context"
	"encoding/json"
	"tech-challenge-hackaton/internal/application/entities"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type AwsSQSService struct {
	client *sqs.Client
	queueProcessVideo string
}

func NewAwsSQSService(client *sqs.Client, queueProcessVideo string) *AwsSQSService {
	return &AwsSQSService{client: client, queueProcessVideo: queueProcessVideo}
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
	return s.sendMessage(video, s.queueProcessVideo)
}
