package videos

import (
	"fmt"
	"mime/multipart"
	"tech-challenge-hackaton/internal/core/dtos"
	"tech-challenge-hackaton/internal/core/entities"
	"tech-challenge-hackaton/internal/core/repositories"
	"tech-challenge-hackaton/internal/core/services"

	"github.com/google/uuid"
)

type UploadVideoUseCase struct {
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
}

func NewUploadVideoUseCase(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
) *UploadVideoUseCase {
	return &UploadVideoUseCase{
		storageService:  storageService,
		videoRepository: videoRepository,
	}
}

func (uv *UploadVideoUseCase) Execute(
	filename string, file multipart.File, mimeType string) (*dtos.VideoUploadDTO, error) {

	video := entities.CreateVideo(filename, entities.MIMEType(mimeType))

	newFilename, err := uv.storageService.UploadFile(
		fmt.Sprint(video.GetID(), "_", video.GetFilename()), file)
	if err != nil {
		return nil, err
	}

	err = uv.videoRepository.Insert(
		entities.RestoreVideo(video.GetID(), video.GetStatus(), newFilename, video.GetMimeType()),
	)
	if err != nil {
		return nil, err
	}

	return &dtos.VideoUploadDTO{ID: uuid.NewString(), Filename: newFilename}, nil
}
