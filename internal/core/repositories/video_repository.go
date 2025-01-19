package repositories

import "tech-challenge-hackaton/internal/core/entities"

type VideoRepositoryInterface interface {
	Insert(video *entities.Video) error
}
