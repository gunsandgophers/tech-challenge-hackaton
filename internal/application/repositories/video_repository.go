package repositories

import "tech-challenge-hackaton/internal/application/entities"

type VideoRepositoryInterface interface {
	Insert(video *entities.Video) error
	Update(video *entities.Video) error
	Get(videoID string) (*entities.Video, error)
	ListByUserID(userID string) ([]*entities.Video, error)
}
