package services

import "github.com/golang-jwt/jwt/v5"


type UserAccessToken struct {
	TokenString string
	TokenJWT *jwt.Token
}

type UserDTO struct {
	ID string
	Name string
	Email string
}

type UserManagerServiceInterface interface {
	Login(username string, password string) (string, error)
	ValidateAccessTokenByAuthHeader(authHeader string) (*UserAccessToken, error)
	GetUser(token *UserAccessToken) (*UserDTO, error)
}
