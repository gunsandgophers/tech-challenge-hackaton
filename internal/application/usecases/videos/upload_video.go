package videos

import (
	"mime/multipart"
	"tech-challenge-hackaton/internal/application/entities"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
	"tech-challenge-hackaton/internal/application/vo"
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
	filename string,
	file multipart.File,
	mimeType string,
	userID string,
) (*VideoUploadDTO, error) {
	video, err := entities.CreateVideo(userID, filename, vo.MIMEType(mimeType))
	if err != nil {
		return nil, err
	}

	newFilename, err := uv.storageService.UploadVideo(video.GetFullFilename(), file)
	if err != nil {
		return nil, err
	}

	if err := uv.videoRepository.Insert(video); err != nil {
		return nil, err
	}

	msg := services.VideoUploadedMessage{
		VideoID: video.GetID(),
		Filename: video.GetFullFilename(),
	}
	if err := uv.queueService.SendVideoUploadedMessage(msg); err != nil {
		return nil, err
	}

	return &VideoUploadDTO{ID: video.GetID(), Filename: newFilename}, nil
}
