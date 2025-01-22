package services

import (
	"tech-challenge-hackaton/internal/infra/clients"
)

type AWSCognitoService struct {
	client *clients.CognitoClient
}

func NewAWSCognitoService(client *clients.CognitoClient) *AWSCognitoService {
	return &AWSCognitoService{
		client: client,
	}
}

// TODO: Receber DTO UserLogin e retornar AccessToken tipo
func (um *AWSCognitoService) Login(username string, password string) (string, error) {
	return um.client.Login(username, password)
}
