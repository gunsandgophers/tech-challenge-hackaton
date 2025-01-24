package services

import (
	"fmt"
	"strings"
	"tech-challenge-hackaton/internal/application/services"
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

func (um *AWSCognitoService) Login(username string, password string) (string, error) {
	return um.client.Login(username, password)
}

func (um *AWSCognitoService) ValidateAccessTokenByAuthHeader(authHeader string) (*services.UserAccessToken, error) {
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return nil, fmt.Errorf("invalid authorization header token")
	}
	tokenString := splitToken[1]
	if tokenString == "" {
		return nil, fmt.Errorf("invalid authorization header token")
	}
	tokenJWT, err := um.client.ValidateAccessToken(tokenString)
	if err != nil {
		return nil, err
	}
	token := &services.UserAccessToken{
		TokenString: tokenString,
		TokenJWT:    tokenJWT,
	}
	return token, nil
}

func (um *AWSCognitoService) GetUser(token *services.UserAccessToken) (*services.UserDTO, error) {
	user, err := um.client.GetUser(token.TokenString)
	if err != nil {
		return nil, err
	}
	return &services.UserDTO{
		ID: user.Username,
		Name: user.Name,
		Email: user.Email,
	}, nil
}
