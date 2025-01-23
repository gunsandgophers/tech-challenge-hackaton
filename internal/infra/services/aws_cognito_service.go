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

// TODO: Receber DTO UserLoginDTO e retornar AccessToken tipo
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

// TODO: Deve retornar um UserDTO
func (um *AWSCognitoService) GetUser(token *services.UserAccessToken) (string, error) {
	// TODO: Deve retornar todos os dados uteis do client
	return um.client.GetUser(token.TokenString)
}
