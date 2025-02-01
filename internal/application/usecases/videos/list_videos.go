package videos

import (
	"tech-challenge-hackaton/internal/application/repositories"
)

type VideoOutput struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Status   string `json:"status"`
	Filename string `json:"filename"`
	MIMEType string `json:"mime_type"`
}

type ListVideosOutput struct {
	Videos []VideoOutput `json:"videos"`
}

type ListVideosUseCase struct {
	videoRepository repositories.VideoRepositoryInterface
}

func NewListVideosUseCase(
	videoRepository repositories.VideoRepositoryInterface,
) *ListVideosUseCase {
	return &ListVideosUseCase{
		videoRepository: videoRepository,
	}
}

func (lv *ListVideosUseCase) Execute(userID string) (*ListVideosOutput, error) {
	videos, err := lv.videoRepository.ListByUserID(userID)
	if err != nil {
		return nil, err
	}
	output := &ListVideosOutput{}
	for _, video := range videos {
		output.Videos = append(output.Videos, VideoOutput{
			ID: video.GetID(),
			UserID: video.GetUserID(),
			Status: video.GetStatus().String(),
			Filename: video.GetFilename(),
			MIMEType: video.GetMimeType().String(),
		})
	}

	return output, err
}
