package services

import "tech-challenge-hackaton/internal/infra/clients"

type FFMPEGService struct {
	client *clients.FFMPEGClient
}

func NewFFMPEGService(client *clients.FFMPEGClient) *FFMPEGService {
	return &FFMPEGService{
		client: client,
	}
}

func (f *FFMPEGService) Snapshot(videoID, filename string, interval int) error {
	return nil
}

