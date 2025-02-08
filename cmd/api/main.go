package main

import (
	"tech-challenge-hackaton/configs"
	"tech-challenge-hackaton/internal/consumers"
	"tech-challenge-hackaton/internal/infra/clients"
	"tech-challenge-hackaton/internal/infra/database"
	httpserver "tech-challenge-hackaton/internal/infra/http"
	"tech-challenge-hackaton/internal/infra/repositories"
	"tech-challenge-hackaton/internal/infra/services"
	"tech-challenge-hackaton/internal/utils"
	"tech-challenge-hackaton/internal/web/app"

	_ "tech-challenge-hackaton/docs"
)


// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3003
// @BasePath  /api/v1
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token like this format: Bearer {token}

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	config := utils.Must(configs.LoadConfig("../.", ".env"))
	// HTTP Server and Database Connection
	httpServer := httpserver.NewGinHTTPServerAdapter()
	connection := database.NewPGXConnectionAdapter(
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	defer connection.Close()
	// CLIENTS
	storageClient := clients.NewS3Client(
		config.AWSS3Region,
		config.AWSS3AccessKeyID,
		config.AWSS3SecretAccessKey,
		config.AWSS3BaseEndpoint,
	)
	queueClient := clients.NewSQSClient(
		config.AWSSQSRegion,
		config.AWSSQSAccessKeyID,
		config.AWSSQSSecretAccessKey,
		config.AWSSQSBaseEndpoint,
	)
	cognitoClient := clients.NewCognitoClient(
		config.AWSCognitoRegion,
		config.AWSCognitoAccessKeyID,
		config.AWSCognitoSecretAccessKey,
		config.AWSCognitoAppClientID,
		config.AWSCognitoAppClientSecret,
		config.AWSCognitoUserPoolID,
	)
	// SERVICES
	storageService := services.NewAwsS3Service(storageClient, config.AWSS3BucketName)
	queueService := services.NewAwsSQSService(queueClient, config.QueueProcessVideo, config.QueueResultVideo)
	userManagerService := services.NewAWSCognitoService(cognitoClient)
	// REPOSITORIES
	videoRepository := repositories.NewVideoRepositoryDB(connection)
	// CONSUMER
	updateVideoProcessedConsumer := consumers.NewVideoProcessedConsumer(queueService, videoRepository)
	// APP
	app := app.NewAPIApp(
		httpServer,
		storageService,
		videoRepository,
		queueService,
		userManagerService,
	)

	go updateVideoProcessedConsumer.Run()
	app.Run()
}
