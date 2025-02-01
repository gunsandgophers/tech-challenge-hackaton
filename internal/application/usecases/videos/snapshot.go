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

const snapshotInterval int = 20

func (s *SnapshotUseCase) Execute(input SnapshotInput) error {
	// PROCESS ....
	// Download video
	_, err := s.snapshotService.Snapshot(input.VideoID, input.Filename, snapshotInterval)
	if err != nil {
		return err
	}
	// Upload zip file
	// Delete local zip file
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
