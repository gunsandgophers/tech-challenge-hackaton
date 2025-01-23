package main

import (
	"tech-challenge-hackaton/configs"
	"tech-challenge-hackaton/internal/infra/clients"
	"tech-challenge-hackaton/internal/infra/database"
	httpserver "tech-challenge-hackaton/internal/infra/http"
	"tech-challenge-hackaton/internal/infra/repositories"
	"tech-challenge-hackaton/internal/infra/services"
	"tech-challenge-hackaton/internal/utils"
	"tech-challenge-hackaton/internal/web/app"
)

func main() {
	config := utils.Must(configs.LoadConfig("../.", ".env.dev"))
	// HTTP Server and Database Connection
	httpServer := httpserver.NewGinHTTPServerAdapter()
	connection := database.NewPGXConnectionAdapter(
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
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
		config.AWSCognitoUserPoolID,
	)
	// SERVICES
	storageService := services.NewAwsS3Service(storageClient, config.AWSS3BucketName)
	queueService := services.NewAwsSQSService(queueClient, config.QueueProcessVideo)
	userManagerService := services.NewAWSCognitoService(cognitoClient)
	// REPOSITORIES
	videoRepository := repositories.NewVideoRepositoryDB(connection)
	// APP
	app := app.NewAPIApp(
		httpServer,
		storageService,
		videoRepository,
		queueService,
		userManagerService,
	)
	app.Run()
	defer connection.Close()
}
