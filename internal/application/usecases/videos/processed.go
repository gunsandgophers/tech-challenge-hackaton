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
	return nil
}
