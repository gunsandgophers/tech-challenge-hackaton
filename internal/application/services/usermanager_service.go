package services

import "github.com/golang-jwt/jwt/v5"


type UserAccessToken struct {
	TokenString string
	TokenJWT *jwt.Token
}

type UserManagerServiceInterface interface {
	Login(username string, password string) (string, error)
	ValidateAccessTokenByAuthHeader(authHeader string) (*UserAccessToken, error)
	GetUser(token *UserAccessToken) (string, error)
}
