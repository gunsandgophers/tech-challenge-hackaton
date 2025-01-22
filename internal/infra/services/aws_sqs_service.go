package services

import (
	"encoding/json"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/infra/clients"
)

type AwsSQSService struct {
	client *clients.SQSClient
	queueProcessVideo string
}

func NewAwsSQSService(client *clients.SQSClient, queueProcessVideo string) *AwsSQSService {
	return &AwsSQSService{
		client: client,
		queueProcessVideo: queueProcessVideo,
	}
}

func (s *AwsSQSService) sendMessage(message interface{}, queueURL string) error {
	body, err := json.Marshal(message)
	if err != nil {
		return nil
	}
	return s.client.SendMessage(string(body), queueURL)
}

func (s *AwsSQSService) SendVideoForProcessing(video *entities.Video) error {
	return s.sendMessage(video, s.queueProcessVideo)
}
