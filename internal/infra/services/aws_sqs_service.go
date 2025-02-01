package services

import (
	"encoding/json"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/infra/clients"
)

type AwsSQSService struct {
	client            *clients.SQSClient
	queueProcessVideo string
	queueResultVideo string
}

const maxMessagesReceive int = 1

func NewAwsSQSService(client *clients.SQSClient, queueProcessVideo, queueResultVideo string) *AwsSQSService {
	return &AwsSQSService{
		client:            client,
		queueProcessVideo: queueProcessVideo,
		queueResultVideo: queueResultVideo,
	}
}

func (s *AwsSQSService) sendMessage(message interface{}, queueURL string) error {
	body, err := json.Marshal(message)
	if err != nil {
		return nil
	}
	return s.client.SendMessage(string(body), queueURL)
}

func (s *AwsSQSService) receiveVideoMessage(queueURL string) ([]services.VideoMessage, error) {
	messages, err := s.client.ReceiveMessages(maxMessagesReceive, queueURL)
	if err != nil {
		return nil, err
	}

	var videoMessages []services.VideoMessage
	for _, message := range messages {
		var m services.VideoMessage
		if err := json.Unmarshal([]byte(*message.Body), &m); err != nil {
			return nil, err
		}
		m.MessageID = *message.ReceiptHandle
		videoMessages = append(videoMessages, m)
	}
	return videoMessages, nil
}

func (s *AwsSQSService) SendVideoUploadedMessage(msg services.VideoUploadedMessage) error {
	return s.sendMessage(msg, s.queueProcessVideo)
}

func (s *AwsSQSService) ReceiveVideoUploadedMessage() ([]services.VideoUploadedMessage, error) {
	videoMessages, err := s.receiveVideoMessage(s.queueProcessVideo)
	if err != nil {
		return nil, err
	}

	var messages []services.VideoUploadedMessage
	for _, vm := range videoMessages {
		messages = append(messages, services.VideoUploadedMessage{
			MessageID: vm.MessageID,
			VideoID: vm.VideoID,
			Filename: vm.Filename,
		})
	}
	return messages, nil
}

func (s *AwsSQSService) AckVideoUploadedMessage(messageID string) error {
	return s.client.DeleteMessage(messageID, s.queueProcessVideo)
}

func (s *AwsSQSService) SendVideoProcessedMessage(msg services.VideoProcessedMessage) error {
	return s.sendMessage(msg, s.queueResultVideo)
}

func (s *AwsSQSService) ReceiveVideoProcessedMessage() ([]services.VideoProcessedMessage, error) {
	videoMessages, err := s.receiveVideoMessage(s.queueResultVideo)
	if err != nil {
		return nil, err
	}

	var messages []services.VideoProcessedMessage
	for _, vm := range videoMessages {
		messages = append(messages, services.VideoProcessedMessage{
			MessageID: vm.MessageID,
			VideoID: vm.VideoID,
			Filename: vm.Filename,
		})
	}
	return messages, nil
}

func (s *AwsSQSService) AckVideoProcessedMessage(messageID string) error {
	return s.client.DeleteMessage(messageID, s.queueResultVideo)
}
