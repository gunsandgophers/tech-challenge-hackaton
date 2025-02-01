package videos

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"tech-challenge-hackaton/internal/application/services"
)

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
	_, err := s.storageService.DownloadVideo(input.VideoID, input.Filename)
	if err != nil {
		return err
	}

	zipFilename, err := s.snapshotService.Snapshot(
		input.VideoID,
		s.storageService.GetLocalVideoDir(input.VideoID),
		input.Filename,
		snapshotInterval,
	)
	if err != nil {
		return err
	}

	file, err := os.Open(zipFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	s.storageService.UploadZipFrames(input.Filename, file)
	os.RemoveAll(s.storageService.GetLocalVideoDir(input.VideoID))

	msg := services.VideoProcessedMessage{
		VideoID: input.VideoID,
		Filename: input.Filename,
	}
	if err := s.queueService.SendVideoProcessedMessage(msg); err != nil {
		return err
	}
	return nil
}
