package clients

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"tech-challenge-hackaton/internal/utils"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	cognito "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/golang-jwt/jwt/v5"
)

type CognitoUser struct {
	Username string
	Email string
	Name string
}

type CognitoClient struct {
	client *cognito.Client
	awsRegion string
	appClientID string
	appClientSecret string
	userPoolID string
	jwks *keyfunc.JWKS
}

func NewCognitoClient(
	awsRegion, awsAccessKeyID, awsSercretAccessKey string,
	appClientID, appClientSecret string,
	userPoolID string,
) *CognitoClient {
	cfg := utils.Must(
		awsconfig.LoadDefaultConfig(
			context.Background(),
			awsconfig.WithRegion(awsRegion),
			awsconfig.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSercretAccessKey, ""),
			),
		),
	)
	cognitoClient := &CognitoClient{
		client: cognito.NewFromConfig(cfg),
		awsRegion: awsRegion,
		appClientID: appClientID,
		appClientSecret: appClientSecret,
		userPoolID: userPoolID,
	}
	cognitoClient.updateJWKS()
	return cognitoClient
}

func (c *CognitoClient) updateJWKS() {
	jwksURL := fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json",
		c.awsRegion,
		c.userPoolID,
	)
	c.jwks = utils.Must(keyfunc.Get(jwksURL, keyfunc.Options{}))
}

func (c *CognitoClient) calculateSecretHash(username string) (string, error) {
	message := username + c.appClientID
	h := hmac.New(sha256.New, []byte(c.appClientSecret))
	_, err := h.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("erro ao gerar HMAC: %v", err)
	}
	secretHash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return secretHash, nil
}

func (c *CognitoClient) Login(username string, password string) (string, error) {
	secretHash, err := c.calculateSecretHash(username)
	if err != nil {
		return "", err
	}
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeUserPasswordAuth,
		ClientId: aws.String(c.appClientID),
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
			"SECRET_HASH": secretHash,
		},
	}
	result, err := c.client.InitiateAuth(context.Background(), authInput)
	if err != nil {
		return "", err
	}
	return *result.AuthenticationResult.AccessToken, nil
}

func (c *CognitoClient) GetUser(tokenString string) (*CognitoUser, error) {
	input := &cognito.GetUserInput{
		AccessToken: aws.String(tokenString),
	}
	output, err := c.client.GetUser(context.Background(), input)
	if err != nil {
		return nil, err
	}

	user := &CognitoUser{Username: *output.Username}
	for _, v := range output.UserAttributes {
		switch *v.Name {
		case "name":
			user.Name = *v.Value
		case "email":
			user.Email = *v.Value
		}
	}

	return user, nil
}

func (c *CognitoClient) ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := c.parseValidToken(tokenString)
	if err != nil {
		return token, err
	}
	return token, c.validateAccessTokenClaims(token)
}

func (c *CognitoClient) parseValidToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		tokenString,
		c.jwks.Keyfunc,
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithIssuer(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s", c.awsRegion, c.userPoolID)),
	)
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, fmt.Errorf("invalid token signature")
	}
	return token, nil
}

func (c *CognitoClient) validateAccessTokenClaims(token *jwt.Token) error {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid token claims")
	}
	if err := c.validateAccessTokenClaimsExpiration(claims); err != nil {
		return err
	}
	if err := c.validateAccessTokenClaimsTokenUseAccess(claims); err != nil {
		return err
	}
	if err := c.validateAccessTokenClaimsClientID(claims); err != nil {
		return err
	}
	return nil
}

func (c *CognitoClient) validateAccessTokenClaimsExpiration(claims jwt.MapClaims) error {
	expClaim, err := claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("invalid token claims")
	}
	if expClaim.Unix() < time.Now().Unix() {
		return fmt.Errorf("token expired")
	}
	return nil
}

func (c *CognitoClient) validateAccessTokenClaimsTokenUseAccess(claims jwt.MapClaims) error {
	tokenUseClaim, ok := claims["token_use"].(string)
	if !ok {
		return fmt.Errorf("error on get token use claims")
	}
	if tokenUseClaim != "access" {
		return fmt.Errorf("invalid token use claim access")
	}
	return nil
}

func (c *CognitoClient) validateAccessTokenClaimsClientID(claims jwt.MapClaims) error {
	clientIDClaim, ok := claims["client_id"].(string)
	if !ok {
		return fmt.Errorf("error on get client id")
	}
	if clientIDClaim != c.appClientID {
		return fmt.Errorf("invalid client id")
	}
	return nil
}
