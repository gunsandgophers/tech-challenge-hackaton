package services

import "tech-challenge-hackaton/internal/core/entities"

type QueueServiceInterface interface {
	SendVideoForProcessing(video *entities.Video) error
}
