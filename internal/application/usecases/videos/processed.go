package videos

import (
	"tech-challenge-hackaton/internal/application/repositories"
)

type UpdateProcessedVideoUseCase struct {
	repository repositories.VideoRepositoryInterface
}

type UpdateProcessedVideoInput struct {
	VideoID string
	Filename string
}

func NewUpdateProcessedVideoUseCase(
	repository repositories.VideoRepositoryInterface,
) *UpdateProcessedVideoUseCase {
	return &UpdateProcessedVideoUseCase{repository: repository}
}

func (u *UpdateProcessedVideoUseCase) Execute(input UpdateProcessedVideoInput) error {
	video, err := u.repository.Get(input.VideoID)
	if err != nil {
		return err
	}
	video.Finished()
	return u.repository.Update(video)
}
