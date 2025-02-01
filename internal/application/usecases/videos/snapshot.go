package videos

import "tech-challenge-hackaton/internal/application/services"

type SnapshotUseCase struct {
	queueService services.QueueServiceInterface
	storageService services.StorageServiceInterface
	snapshotService services.SnapshotServiceInterface
}

type SnapshotInput struct {
	VideoID string
	Filename string
}

func NewSnapshotUseCase(
	queueService services.QueueServiceInterface,
	storageService services.StorageServiceInterface,
	snapshotService services.SnapshotServiceInterface,
) *SnapshotUseCase {
	return &SnapshotUseCase{
		queueService: queueService,
		storageService: storageService,
		snapshotService: snapshotService,
	}
}

func (s *SnapshotUseCase) Execute(input SnapshotInput) error {
	// PROCESS ....
	// END PROCESS ...
	msg := services.VideoProcessedMessage{
		VideoID: input.VideoID,
		Filename: input.Filename,
	}
	if err := s.queueService.SendVideoProcessedMessage(msg); err != nil {
		return err
	}
	return nil
}
