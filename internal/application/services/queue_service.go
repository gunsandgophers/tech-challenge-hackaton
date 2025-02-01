package services

type VideoMessage struct {
	MessageID string `json:"message_id,omitempty"`
	VideoID        string `json:"id,omitempty"`
	Filename  string `json:"filename,omitempty"`
}

type VideoUploadedMessage struct {
	MessageID string `json:"message_id,omitempty"`
	VideoID        string `json:"id,omitempty"`
	Filename  string `json:"filename,omitempty"`
}

type VideoProcessedMessage struct {
	MessageID string `json:"message_id,omitempty"`
	VideoID        string `json:"id,omitempty"`
	Filename  string `json:"filename,omitempty"`
}

type QueueServiceInterface interface {
	SendVideoUploadedMessage(msg VideoUploadedMessage) error
	ReceiveVideoUploadedMessage() ([]VideoUploadedMessage, error)
	AckVideoUploadedMessage(messageID string) error
	SendVideoProcessedMessage(msg VideoProcessedMessage) error
	ReceiveVideoProcessedMessage() ([]VideoProcessedMessage, error)
	AckVideoProcessedMessage(messageID string) error
}
