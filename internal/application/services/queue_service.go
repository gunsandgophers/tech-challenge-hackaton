package services

import "tech-challenge-hackaton/internal/application/entities"

type QueueServiceInterface interface {
	SendVideoForProcessing(video *entities.Video) error
}
