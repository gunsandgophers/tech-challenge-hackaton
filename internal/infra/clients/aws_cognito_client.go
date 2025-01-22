package clients

import (
	"fmt"
	"tech-challenge-hackaton/internal/utils"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient struct {
	client *cognito.CognitoIdentityProvider
	awsRegion string
	appClientID string
	userPoolID string
}

func NewCognitoClient(
	awsRegion string,
	appClientID string,
	userPoolID string,
) *CognitoClient {
	config := &aws.Config{Region: aws.String(awsRegion)}
	sess := utils.Must(session.NewSession(config))
	return &CognitoClient{
		client: cognito.New(sess),
		awsRegion: awsRegion,
		appClientID: appClientID,
		userPoolID: userPoolID,
	}
}

func (c *CognitoClient) getJWKS() (*keyfunc.JWKS, error) {
	jwksURL := fmt.Sprintf(
		"https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json",
		c.awsRegion,
		c.userPoolID,
	)
	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{})
	if err != nil {
		return nil, err
	}
	return jwks, nil
}

func (c *CognitoClient) Login(username string, password string) (string, error) {
	authInput := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		}),
		ClientId: aws.String(c.appClientID),
	}
	result, err := c.client.InitiateAuth(authInput)
	if err != nil {
		return "", err
	}
	return *result.AuthenticationResult.AccessToken, nil
}
