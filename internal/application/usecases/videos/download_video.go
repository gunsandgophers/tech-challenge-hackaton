package videos

import (
	"fmt"
	"tech-challenge-hackaton/internal/application/repositories"
	"tech-challenge-hackaton/internal/application/services"
)

type DownloadVideoFramesUseCaseOutput struct {
	Filename string
	MIMEType string
	Content []byte
}

type DownloadVideoFramesUseCase struct {
	storageService  services.StorageServiceInterface
	videoRepository repositories.VideoRepositoryInterface
}

func NewDownloadVideoFramesUseCase(
	storageService services.StorageServiceInterface,
	videoRepository repositories.VideoRepositoryInterface,
) *DownloadVideoFramesUseCase {
	return &DownloadVideoFramesUseCase{
		storageService:  storageService,
		videoRepository: videoRepository,
	}
}

func (uv *DownloadVideoFramesUseCase) Execute(videoID string, userID string) (*DownloadVideoFramesUseCaseOutput, error) {
	video, err := uv.videoRepository.Get(videoID)
	if err != nil {
		return nil, err
	}

	if err := video.IsAvaiableToDownload(userID); err != nil {
		return nil, err
	}

	zipframes, err := uv.storageService.DownloadZipFrames(video.GetID())
	if err != nil {
		return nil, err
	}

	output := &DownloadVideoFramesUseCaseOutput{
		Filename: fmt.Sprintf("%s.zip", video.GetID()),
		MIMEType: "application/zip",
		Content: zipframes,
	}
	return output, nil
}
