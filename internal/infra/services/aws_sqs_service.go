package services

import (
	"encoding/json"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/vo"
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

func (s *AwsSQSService) receiveVideoMessage(queueURL string) ([]services.VideoMessageDTO, error) {
	messages, err := s.client.ReceiveMessages(maxMessagesReceive, queueURL)
	if err != nil {
		return nil, err
	}

	var videoMessages []services.VideoMessageDTO
	for _, message := range messages {
		var m services.VideoMessageDTO
		if err := json.Unmarshal([]byte(*message.Body), &m); err != nil {
			return nil, err
		}
		m.MessageID = *message.ReceiptHandle
		videoMessages = append(videoMessages, m)
	}
	return videoMessages, nil
}

func (s *AwsSQSService) SendVideoUploadedMessage(video *entities.Video) error {
	message := services.VideoMessageDTO{
		ID:       video.GetID(),
		UserID:   video.GetUserID(),
		Status:   video.GetStatus().String(),
		Filename: video.GetFilename(),
		MIMEType: video.GetMimeType().String(),
	}
	return s.sendMessage(message, s.queueProcessVideo)
}

func (s *AwsSQSService) ReceiveVideoUploadedMessage() ([]services.VideoUploadedMessage, error) {
	videoMessages, err := s.receiveVideoMessage(s.queueProcessVideo)
	if err != nil {
		return nil, err
	}

	var messages []services.VideoUploadedMessage
	for _, vm := range videoMessages {
		video, err := entities.RestoreVideo(
			vm.ID,
			vm.UserID,
			vo.VideoStatus(vm.Status),
			vm.Filename,
			vo.MIMEType(vm.MIMEType),
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, services.VideoUploadedMessage{ID: vm.MessageID, Video: video})
	}
	return messages, nil
}

func (s *AwsSQSService) AckVideoUploadedMessage(messageID string) error {
	return s.client.DeleteMessage(messageID, s.queueProcessVideo)
}

func (s *AwsSQSService) SendVideoProcessedMessage(video *entities.Video) error {
	message := services.VideoMessageDTO{
		ID:       video.GetID(),
		UserID:   video.GetUserID(),
		Status:   video.GetStatus().String(),
		Filename: video.GetFilename(),
		MIMEType: video.GetMimeType().String(),
	}
	return s.sendMessage(message, s.queueResultVideo)
}

func (s *AwsSQSService) ReceiveVideoProcessedMessage() ([]services.VideoProcessedMessage, error) {
	videoMessages, err := s.receiveVideoMessage(s.queueResultVideo)
	if err != nil {
		return nil, err
	}

	var messages []services.VideoProcessedMessage
	for _, vm := range videoMessages {
		video, err := entities.RestoreVideo(
			vm.ID,
			vm.UserID,
			vo.VideoStatus(vm.Status),
			vm.Filename,
			vo.MIMEType(vm.MIMEType),
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, services.VideoProcessedMessage{ID: vm.MessageID, Video: video})
	}
	return messages, nil
}

func (s *AwsSQSService) AckVideoProcessedMessage(messageID string) error {
	return s.client.DeleteMessage(messageID, s.queueResultVideo)
}
