package videos

import (
	"log"
	"os"
	"tech-challenge-hackaton/internal/application/services"
)

type SnapshotUseCase struct {
	queueService    services.QueueServiceInterface
	storageService  services.StorageServiceInterface
	snapshotService services.SnapshotServiceInterface
}

type SnapshotInput struct {
	VideoID  string
	Filename string
}

func NewSnapshotUseCase(
	queueService services.QueueServiceInterface,
	storageService services.StorageServiceInterface,
	snapshotService services.SnapshotServiceInterface,
) *SnapshotUseCase {
	return &SnapshotUseCase{
		queueService:    queueService,
		storageService:  storageService,
		snapshotService: snapshotService,
	}
}

const snapshotInterval int = 20

func (s *SnapshotUseCase) Execute(input SnapshotInput) error {
	_, err := s.storageService.DownloadVideo(input.VideoID, input.Filename)
	if err != nil {
		return err
	}

	zipFilenameComplete, zipFilename, err := s.snapshotService.Snapshot(
		input.VideoID,
		s.storageService.GetLocalVideoDir(input.VideoID),
		input.Filename,
		snapshotInterval,
	)
	if err != nil {
		log.Println("Error snapshot service")
		return err
	}

	file, err := os.Open(zipFilenameComplete)
	if err != nil {
		return err
	}
	defer file.Close()
	defer os.RemoveAll(s.storageService.GetLocalVideoDir(input.VideoID))

	_, err = s.storageService.UploadZipFrames(zipFilename, file)
	if err != nil {
		return err
	}

	msg := services.VideoProcessedMessage{
		VideoID:  input.VideoID,
		Filename: input.Filename,
	}
	if err := s.queueService.SendVideoProcessedMessage(msg); err != nil {
		return err
	}
	return nil
}
