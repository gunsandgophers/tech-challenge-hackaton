package consumers

import (
	"log"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/usecases/videos"
)

type VideoProcessedConsumer struct {
	queueService services.QueueServiceInterface
	repository repositories.VideoRepositoryInterface
}

func NewVideoProcessedConsumer(
	queueService services.QueueServiceInterface,
	repository repositories.VideoRepositoryInterface,
) *VideoProcessedConsumer {
	return &VideoProcessedConsumer{queueService: queueService, repository: repository}
}

func (c *VideoProcessedConsumer) Run() {
	messagesToProcess := make(chan services.VideoProcessedMessage)
	go func(){
		uc := videos.NewUpdateProcessedVideoUseCase(c.repository)
		for msg := range messagesToProcess {
			input := videos.UpdateProcessedVideoInput{
				VideoID: msg.VideoID,
				Filename: msg.Filename,
			}
			if err := uc.Execute(input); err != nil {
				log.Println(err.Error())
			}
			c.queueService.AckVideoProcessedMessage(msg.MessageID)
		}
	}()

	for {
		messages, err := c.queueService.ReceiveVideoProcessedMessage()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		for _, msg := range messages {
			messagesToProcess <- msg
		}
	}
}

