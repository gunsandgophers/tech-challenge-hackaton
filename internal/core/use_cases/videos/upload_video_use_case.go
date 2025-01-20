package videos

import (
	"fmt"
	"mime/multipart"
	"tech-challenge-hackaton/internal/core/dtos"
	"tech-challenge-hackaton/internal/core/entities"
	"tech-challenge-hackaton/internal/core/repositories"
	"tech-challenge-hackaton/internal/core/services"
)

type UploadVideoUseCase struct {
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
	queueService    services.QueueServiceInterface
}

func NewUploadVideoUseCase(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
	queueService services.QueueServiceInterface,
) *UploadVideoUseCase {
	return &UploadVideoUseCase{
		storageService:  storageService,
		videoRepository: videoRepository,
		queueService:    queueService,
	}
}

func (uv *UploadVideoUseCase) Execute(
	filename string, file multipart.File, mimeType string,
) (*dtos.VideoUploadDTO, error) {

	video, err := entities.CreateVideo(filename, entities.MIMEType(mimeType))
	if err != nil {
		return nil, err
	}

	newFilename, err := uv.storageService.UploadFile(
		fmt.Sprint(video.GetID(), "_", video.GetFilename()), file)
	if err != nil {
		return nil, err
	}

	video, _ = entities.
		RestoreVideo(video.GetID(), video.GetStatus(), newFilename, video.GetMimeType())

	err = uv.videoRepository.Insert(video)
	if err != nil {
		return nil, err
	}

	err = uv.queueService.SendVideoForProcessing(video)
	if err != nil {
		return nil, err
	}

	return &dtos.VideoUploadDTO{ID: video.GetID(), Filename: newFilename}, nil
}
