package consumers

import (
	"log"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/usecases/videos"
)

type VideoUploadedConsumer struct {
	workers         int
	queueService    services.QueueServiceInterface
	storageService  services.StorageServiceInterface
	snapshotService services.SnapshotServiceInterface
}

func NewVideoUploadedConsumer(
	workers int,
	queueService services.QueueServiceInterface,
	storageService services.StorageServiceInterface,
	snapshotService services.SnapshotServiceInterface,
) *VideoUploadedConsumer {
	return &VideoUploadedConsumer{
		workers:         workers,
		queueService:    queueService,
		storageService:  storageService,
		snapshotService: snapshotService,
	}
}

func (c *VideoUploadedConsumer) Run() {
	messagesToProcess := make(chan services.VideoUploadedMessage)
	for i := 0; i < c.workers; i++ {
		go func() {
			snapshotUC := videos.NewSnapshotUseCase(c.queueService, c.storageService, c.snapshotService)
			for msg := range messagesToProcess {

				log.Println("Message -> ", msg)

				input := videos.SnapshotInput{
					VideoID:  msg.VideoID,
					Filename: msg.Filename,
				}
				if err := snapshotUC.Execute(input); err != nil {
					log.Println("Error on UseCase")
					log.Println(err.Error())
				}
				c.queueService.AckVideoUploadedMessage(msg.MessageID)

				log.Println("Finish snapshot process")
			}
		}()
	}

	for {
		messages, err := c.queueService.ReceiveVideoUploadedMessage()
		if err != nil {
			log.Println("Error receive video message")
			log.Println(err.Error())
			continue
		}

		for _, msg := range messages {
			messagesToProcess <- msg
		}
	}
}
