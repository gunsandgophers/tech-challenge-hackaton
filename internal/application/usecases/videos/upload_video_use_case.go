package videos

import (
	"fmt"
	"mime/multipart"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
)

type VideoUploadDTO struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
}

type VideoUploadResponseDTO struct {
	Videos []*VideoUploadDTO `json:"videos"`
}

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
) (*VideoUploadDTO, error) {

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

	return &VideoUploadDTO{ID: video.GetID(), Filename: newFilename}, nil
}
