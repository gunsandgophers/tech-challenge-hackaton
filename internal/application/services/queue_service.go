package services

import "tech-challenge-hackaton/internal/application/entities"

type VideoMessageDTO struct {
	MessageID string `json:"message_id,omitempty"`
	ID        string `json:"id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Status    string `json:"status,omitempty"`
	Filename  string `json:"filename,omitempty"`
	MIMEType  string `json:"mime_type,omitempty"`
}

type VideoUploadedMessage struct {
	ID    string
	Video *entities.Video
}

type VideoProcessedMessage struct {
	ID    string
	Video *entities.Video
}

type QueueServiceInterface interface {
	SendVideoUploadedMessage(video *entities.Video) error
	ReceiveVideoUploadedMessage() ([]VideoUploadedMessage, error)
	AckVideoUploadedMessage(messageID string) error
	SendVideoProcessedMessage(video *entities.Video) error
	ReceiveVideoProcessedMessage() ([]VideoProcessedMessage, error)
	AckVideoProcessedMessage(messageID string) error
}
