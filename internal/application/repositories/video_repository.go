package repositories

import "tech-challenge-hackaton/internal/application/entities"

type VideoRepositoryInterface interface {
	Insert(video *entities.Video) error
	ListByUserID(userID string) ([]*entities.Video, error)
}
